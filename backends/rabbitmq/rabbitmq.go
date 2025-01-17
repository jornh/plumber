package rabbitmq

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/types"
	"github.com/batchcorp/plumber/validate"
	"github.com/batchcorp/rabbit"
)

const (
	BackendName = "rabbitmq"
)

var (
	ErrEmptyRoutingKey   = errors.New("routing key cannot be empty")
	ErrEmptyExchangeName = errors.New("exchange name cannot be empty")
	ErrEmptyQueueName    = errors.New("queue name cannot be empty")
	ErrEmptyBindingKey   = errors.New("binding key cannot be empty")
)

// RabbitMQ holds all attributes required for performing a read/write operations
// in RabbitMQ. This struct should be instantiated via the rabbitmq.Read(..) or
// rabbitmq.Write(..) functions.
type RabbitMQ struct {
	connOpts *opts.ConnectionOptions
	connArgs *args.RabbitConn
	client   rabbit.IRabbit
	log      *logrus.Entry
}

func New(opts *opts.ConnectionOptions) (*RabbitMQ, error) {
	if err := validateBaseConnOpts(opts); err != nil {
		return nil, errors.Wrap(err, "invalid connection options")
	}

	connArgs := opts.GetRabbit()

	r := &RabbitMQ{
		connOpts: opts,
		connArgs: connArgs,
		log:      logrus.WithField("backend", BackendName),
	}

	return r, nil
}

func (r *RabbitMQ) newRabbitForRead(readArgs *args.RabbitReadArgs) (rabbit.IRabbit, error) {
	rmq, err := rabbit.New(&rabbit.Options{
		URL:            r.connArgs.Address,
		QueueName:      readArgs.QueueName,
		ExchangeName:   readArgs.ExchangeName,
		RoutingKey:     readArgs.BindingKey,
		QueueExclusive: readArgs.QueueExclusive,
		QueueDurable:   readArgs.QueueDurable,
		QueueDeclare:   readArgs.QueueDeclare,
		AutoAck:        readArgs.AutoAck,
		ConsumerTag:    readArgs.ConsumerTag,
		UseTLS:         r.connArgs.UseTls,
		SkipVerifyTLS:  r.connArgs.TlsSkipVerify,
		Mode:           rabbit.Consumer,
	})

	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize rabbitmq consumer")
	}

	return rmq, nil
}

func (r *RabbitMQ) newRabbitForWrite(writeArgs *args.RabbitWriteArgs) (rabbit.IRabbit, error) {
	rmq, err := rabbit.New(&rabbit.Options{
		URL:                r.connArgs.Address,
		ExchangeName:       writeArgs.ExchangeName,
		RoutingKey:         writeArgs.RoutingKey,
		ExchangeDeclare:    writeArgs.ExchangeDeclare,
		ExchangeDurable:    writeArgs.ExchangeDurable,
		ExchangeAutoDelete: writeArgs.ExchangeAutoDelete,
		ExchangeType:       writeArgs.ExchangeType,
		AppID:              writeArgs.AppId,
		UseTLS:             r.connArgs.UseTls,
		SkipVerifyTLS:      r.connArgs.TlsSkipVerify,
		Mode:               rabbit.Producer,
	})

	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize rabbitmq consumer")
	}

	return rmq, nil
}

func (r *RabbitMQ) Name() string {
	return BackendName
}

func (r *RabbitMQ) Close(_ context.Context) error {
	// Since these are instantiated inside the Write/Read/Relay methods,
	// there is chance that r.client might be nil
	if r.client != nil {
		return r.client.Close()
	}
	return nil
}

func (r *RabbitMQ) Test(_ context.Context) error {
	return types.NotImplementedErr
}

func validateBaseConnOpts(connOpts *opts.ConnectionOptions) error {
	if connOpts == nil {
		return validate.ErrMissingConnOpts
	}

	if connOpts.Conn == nil {
		return validate.ErrMissingConnCfg
	}

	args := connOpts.GetRabbit()
	if args == nil {
		return validate.ErrMissingConnArgs
	}

	if args.Address == "" {
		return validate.ErrMissingAddress
	}

	return nil
}
