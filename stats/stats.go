package stats

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/nakabonne/tstorage"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
)

type IStats interface {
	AddCounter(cfg *opts.Counter) *Counter
	GetAllCounters() map[string]*Counter
	GetCounter(counterType opts.Counter_Type, resourceType opts.Counter_Resource, resourceID string) (*Counter, error)
	RemoveCounter(counter *Counter) error
}

type Counter struct {
	cfg     *opts.Counter
	mtx     *sync.RWMutex
	storage tstorage.Storage
	doneCh  chan struct{}
	log     *logrus.Entry
}

type Stats struct {
	*Config
	countersMtx *sync.RWMutex
	counters    map[string]*Counter
	log         *logrus.Entry
}

type Config struct {
	// FlushInterval is how often the goroutine will flush a counter's value to time-series db
	FlushInterval time.Duration

	// ServiceShutdownCtx is used to signal all counters to flush and then cleanly close the database
	ServiceShutdownCtx context.Context

	Storage tstorage.Storage
}

var (
	ErrInvalidFlushInterval = errors.New("FlushInterval cannot be 0")
	ErrCounterNotFound      = errors.New("counter not found")
	ErrMissingShutdownCtx   = errors.New("ServiceShutdownCtx cannot be nil")
	ErrMissingStorage       = errors.New("ErrMissingStorage cannot be empty")
)

// New returns a configured struct for the statistics service
func New(cfg *Config) (*Stats, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, errors.Wrap(err, "unable to validate config")
	}

	s := &Stats{
		counters:    make(map[string]*Counter),
		countersMtx: &sync.RWMutex{},
		Config:      cfg,
		log:         logrus.WithField("pkg", "stats"),
	}

	go s.runWatchForShutdown()

	return s, nil
}

// validateConfig validates that the correct options are supplied when initiating the Stats service
func validateConfig(cfg *Config) error {
	if cfg.FlushInterval == 0 {
		return ErrInvalidFlushInterval
	}

	if cfg.ServiceShutdownCtx == nil {
		return ErrMissingShutdownCtx
	}

	if cfg.Storage == nil {
		return ErrMissingStorage
	}

	return nil
}

// runWatchForShutdown is a goroutine that is ran when the Stats services is created. It listens for
// a context cancellation and will stop/flush/delete all counters
func (s *Stats) runWatchForShutdown() error {
	<-s.ServiceShutdownCtx.Done()

	s.countersMtx.RLock()
	counters := s.counters // make copy of map since RemoveCounter will need to acquire mutex lock each time
	s.countersMtx.RUnlock()

	for _, c := range counters {
		s.RemoveCounter(c)
	}

	if err := s.Storage.Close(); err != nil {
		err = errors.Wrap(err, "unable to safely shutdown stats service database")
		s.log.Error(err)
		return err
	}

	return nil
}

// genCounterID generates a string which is used to identify/retreive a counter from the counters map
func genCounterID(counterType opts.Counter_Type, resourceType opts.Counter_Resource, resourceID string) string {
	return fmt.Sprintf("%s-%s-%s", counterType, resourceType, resourceID)
}

func genCounterName(counterType opts.Counter_Type, resourceType opts.Counter_Resource, resourceID string) string {
	var t string

	switch counterType {
	case opts.Counter_TYPE_SCHEMA_VIOLATION:
		t = "Schema Violations"
	case opts.Counter_TYPE_MESSAGE_RECEIVED:
		t = "Messages Received"
	case opts.Counter_TYPE_MESSAGE_REPLAYED:
		t = "Messages Replayed"
	}

	rt := strings.ToTitle(resourceType.String())

	return fmt.Sprintf("%s %s for %s", t, rt, resourceID)
}

// AddCounter creates a new counter and launches a runFlusher() goroutine for it
func (s *Stats) AddCounter(cfg *opts.Counter) *Counter {
	// Check for existing counter
	existing, _ := s.GetCounter(cfg.Type, cfg.Resource, cfg.ResourceId)
	if existing != nil {
		return existing
	}

	cfg.Name = genCounterName(cfg.Type, cfg.Resource, cfg.ResourceId)

	logger := logrus.WithFields(logrus.Fields{
		"pkg":           "stats",
		"resource_type": cfg.Resource.String(),
		"counter_type":  cfg.Type.String(),
		"id":            cfg.ResourceId,
		"name":          cfg.Name,
	})

	c := &Counter{
		cfg:     cfg,
		mtx:     &sync.RWMutex{},
		doneCh:  make(chan struct{}, 1),
		storage: s.Storage,
		log:     logger,
	}

	id := genCounterID(cfg.Type, cfg.Resource, cfg.ResourceId)

	s.countersMtx.Lock()
	s.counters[id] = c
	s.countersMtx.Unlock()

	go c.runFlusher(s.FlushInterval)

	s.log.Debugf("created counter '%s'", id)

	return c
}

// GetCounter retrieves a single counter from the counters map
func (s *Stats) GetCounter(counterType opts.Counter_Type, resourceType opts.Counter_Resource, resourceID string) (*Counter, error) {
	s.countersMtx.RLock()
	defer s.countersMtx.RUnlock()

	id := genCounterID(counterType, resourceType, resourceID)

	c, ok := s.counters[id]
	if !ok {
		return nil, ErrCounterNotFound
	}

	return c, nil
}

// GetAllCounters returns the entire counters map
func (s *Stats) GetAllCounters() map[string]*Counter {
	s.countersMtx.RLock()
	defer s.countersMtx.RUnlock()
	return s.counters
}

// RemoveCounter will flush a counter's value to persistent storage, stop the
// associated runFlusher() go routine, and delete the counter from the counters map
func (s *Stats) RemoveCounter(c *Counter) error {
	// Signal flusher goroutine to flush any data we have to time-series db and terminate
	c.doneCh <- struct{}{}

	id := genCounterID(c.cfg.Type, c.cfg.Resource, c.cfg.ResourceId)

	s.countersMtx.Lock()
	delete(s.counters, id)
	s.countersMtx.Unlock()

	s.log.Debugf("removed counter '%s'", id)

	return nil
}

// Incr increases the value of a counter by the value of i
func (c *Counter) Incr(i float64) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.cfg.Value += i
}

// Value returns the current value of a counter
func (c *Counter) Value() float64 {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.cfg.Value
}

// runFlusher is a goroutine that is launched when a new counter is created. It is used
// to flush the count to persistent storage periodically and during application shutdown
func (c *Counter) runFlusher(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

MAIN:
	for {
		select {
		case <-ticker.C:
			if err := c.flush(); err != nil {
				c.log.Errorf("unable to flush counter '%s': %s", c.cfg.ResourceId, err)
				break
			}

			c.log.Debug("flushed counter")
		case <-c.doneCh:
			if err := c.flush(); err != nil {
				c.log.Errorf("unable to flush counter '%s': %s", c.cfg.ResourceId, err)
				break
			}

			c.log.Debug("flushed counter")
			break MAIN
		}
	}
}

// getStorageLabels returns list of labels that this counter is stored under in the time-series db
func (c *Counter) getStorageLabels() []tstorage.Label {
	return []tstorage.Label{
		{Name: "ResourceType", Value: c.cfg.Resource.String()}, // Connection, read, schema, replay, etc
		{Name: "ResourceID", Value: c.cfg.ResourceId},
	}
}

// flush flushes data to the time-series db
func (c *Counter) flush() error {
	// Get current value, reset to zero, and then release lock
	c.mtx.Lock()
	val := c.cfg.Value
	c.cfg.Value = 0
	c.mtx.Unlock()

	rows := []tstorage.Row{
		{
			Metric:    c.cfg.Type.String(),
			Labels:    c.getStorageLabels(),
			DataPoint: tstorage.DataPoint{Timestamp: time.Now().UTC().Unix(), Value: val},
		},
	}

	if err := c.storage.InsertRows(rows); err != nil {
		return errors.Wrap(err, "unable to flush count to time-series db")
	}

	return nil
}

// GetTSHistory gets time-series history for the counter
func (c *Counter) GetTSHistory(from, to int64) ([]*tstorage.DataPoint, error) {
	points, err := c.storage.Select(c.cfg.Type.String(), c.getStorageLabels(), from, to)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get count history")
	}

	return points, nil
}

// GetTotal aggregates the in-memory and historical total
func (c *Counter) GetTotal() (float64, error) {
	points, err := c.GetTSHistory(0, time.Now().UTC().Unix())
	if err != nil {
		// Nothing in the time-series, just return current memory value
		if errors.Is(err, tstorage.ErrNoDataPoints) {
			return c.Value(), nil
		}
		return 0, err
	}

	total := c.Value()
	for _, point := range points {
		total += point.Value
	}

	return total, nil
}

// GetConfig returns counter config
func (c *Counter) GetConfig() *opts.Counter {
	return c.cfg
}
