package kafka

import (
	"context"

	"github.com/batchcorp/plumber/validate"

	"github.com/pkg/errors"
	skafka "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber/dynamic"
)

// Dynamic starts up a new GRPC client connected to the dProxy service and receives a stream of outbound replay messages
// which are then written to the message bus.
func (k *Kafka) Dynamic(ctx context.Context, opts *opts.DynamicOptions, dynamicSvc dynamic.IDynamic) error {
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

	go dynamicSvc.Start("Kafka")

	outboundCh := dynamicSvc.Read()

	// Continually loop looking for messages on the channel.
MAIN:
	for {
		select {
		case outbound := <-outboundCh:
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

func validateDynamicOptions(dynamicOpts *opts.DynamicOptions) error {
	if dynamicOpts == nil {
		return validate.ErrEmptyDynamicOpts
	}

	if dynamicOpts.Kafka == nil {
		return validate.ErrEmptyBackendGroup
	}

	if dynamicOpts.Kafka.Args == nil {
		return validate.ErrEmptyBackendArgs
	}

	if len(dynamicOpts.Kafka.Args.Topics) == 0 {
		return ErrMissingTopic
	}

	return nil
}
