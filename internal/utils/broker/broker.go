package broker

import "go-fiber-template/pkg/logging"

var log = logging.BrokerLogger

type Publisher interface {
	Publish(topic string, data map[string]interface{}) error
	Close()
}