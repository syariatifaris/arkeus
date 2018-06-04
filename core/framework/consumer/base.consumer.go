package consumer

import (
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/syariatifaris/arkeus/core/config"
	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/log/tokolog"
	"github.com/syariatifaris/arkeus/core/mq"
	"github.com/syariatifaris/arkeus/core/panics"
)

const (
	DefaultDelay        = 0
	DefaultBackoffDelay = -1
	Delay5Seconds       = time.Second * 5
)

//Message Queue Handler Interface
type MQConsumer interface {
	//The Topic Name
	TopicName() string
	//Get the configuration
	Configuration() *config.MQTopicConfig
	//Register All Handlers
	RegisterHandlers(mq.MessageQueue)
}

//Base Message Queue Handler
type BaseMQConsumer struct {
	//All MQ handler should have a topic
	TopicConfigurations map[string]*config.MQTopicConfig
}

//GetConfiguration Get MQ Topic Configuration
func (b *BaseMQConsumer) GetConfiguration(topicName string) *config.MQTopicConfig {
	if t, ok := b.TopicConfigurations[topicName]; ok {
		return t
	}

	return nil
}

//FinishError marks a message as finish due an error
func (b *BaseMQConsumer) FinishError(message *nsq.Message, err error) {
	message.Finish()
	tokolog.WARN.Printf("[BaseMQConsumer][FinishError] NSQ operation aborted at attempt %d, Err: %s. Please check the "+
		"error log", message.Attempts, err.Error())
}

//RetryWithoutBackoff retries without backfoff
func (b *BaseMQConsumer) RetryWithoutBackoff(message *nsq.Message, config *config.MQTopicConfig) {
	if message.Attempts < config.MaxAttempts {
		message.RequeueWithoutBackoff(time.Duration(DefaultDelay))
		tokolog.WARN.Printf("[BaseMQConsumer][RetryWithoutBackoff] NSQ operation retry at attempts %d\n", message.Attempts)
	} else {
		b.FinishError(message, errors.New("NSQ attempts exceeding the limit"))
	}
}

//Retry retries with default backoff period
func (b *BaseMQConsumer) Retry(message *nsq.Message, delay time.Duration, config *config.MQTopicConfig) {
	if message.Attempts < config.MaxAttempts {
		message.Requeue(delay)
		tokolog.WARN.Printf("[BaseMQConsumer][RetryWithDelay] NSQ operation retry at attempts %d/%d Delay: %d\n",
			message.Attempts, config.MaxAttempts, delay)
	} else {
		b.FinishError(message, errors.New("NSQ attempts exceeding the limit"))
	}
}

//Success finishes the message
func (b *BaseMQConsumer) FinishSuccess(message *nsq.Message) {
	message.Finish()
	tokolog.INFO.Printf("[BaseMQConsumer][FinishSuccess] NSQ operation finished at attempt %d", message.Attempts)
}

//CapturePanics wrap the handler message function to capture and restore panics
func (b *BaseMQConsumer) CapturePanic(h mq.HandlerFunc) func(message *nsq.Message) error {
	return func(message *nsq.Message) error {
		defer panics.Restore()
		return h(message)
	}
}
