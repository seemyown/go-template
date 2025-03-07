package broker

import "go-fiber-template/pkg/logging"

var log = logging.New(logging.Config{
	FileName: "broker",
	Name:     "broker",
})

type Publisher interface {
	Publish(topic string, data map[string]interface{}) error
	Close()
}
