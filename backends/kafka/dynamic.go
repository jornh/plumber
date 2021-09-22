package kafka

import (
	"context"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber/dynamic"
	"github.com/pkg/errors"
	skafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

// Dynamic starts up a new GRPC client connected to the dProxy service and receives a stream of outbound replay messages
// which are then written to the message bus.
func (k *Kafka) Dynamic(ctx context.Context, opts *opts.DynamicOptions) error {
	llog := logrus.WithField("pkg", "kafka/dynamic")

	if err := validateDynamicOptions(opts); err != nil {
		return errors.Wrap(err, "unable to validate dynamic options")
	}

	// Start up writer
	writer, err := NewWriter(k.dialer, k.connArgs)
	if err != nil {
		return errors.Wrap(err, "unable to create new writer")
	}

	defer writer.Close()

	// Start up dynamic connection
	grpc, err := dynamic.New(opts, BackendName)
	if err != nil {
		return errors.Wrap(err, "could not establish connection to Batch")
	}

	go grpc.Start()

	// Continually loop looking for messages on the channel.
MAIN:
	for {
		select {
		case outbound := <-grpc.OutboundMessageCh:
			for _, topic := range opts.Kafka.Args.Topics {
				if err := writer.WriteMessages(ctx, skafka.Message{
					Topic: topic,
					Key:   []byte(opts.Kafka.Args.Key),
					Value: outbound.Blob,
				}); err != nil {
					llog.Errorf("Unable to replay message: %s", err)
					break MAIN
				}
			}

		case <-ctx.Done():
			k.log.Warning("context cancelled")
			break MAIN
		}
	}

	k.log.Debug("dynamic exiting")

	return nil
}

func validateDynamicOptions(opts *opts.DynamicOptions) error {
	if opts == nil {
		return errors.New("dynamic options cannot be nil")
	}

	if opts.Kafka == nil {
		return errors.New("kafka options cannot be nil")
	}

	if opts.Kafka.Args == nil {
		return errors.New("kafka args cannot be nil")
	}

	if len(opts.Kafka.Args.Topics) == 0 {
		return errors.New("at least one topic must be provided as an arg")
	}

	return nil
}
