package broker

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
)

type NatsBroker struct {
	natsConn *nats.Conn
}

func NewNatsBroker(natsConn *nats.Conn) *NatsBroker {
	return &NatsBroker{natsConn: natsConn}
}

func (n *NatsBroker) Publish(topic string, data map[string]interface{}) error {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Error(err, "Ошибка дампа")
		return err
	}
	err = n.natsConn.Publish(topic, msg)
	if err != nil {
		log.Error(err,
			"Ошибка отправки сообщения %s в %s", data, topic,
		)
		return err
	}
	log.Debug("Сообщение %s успешно отправлено в %s", data, topic)
	return nil
}

func (n *NatsBroker) Close() {
	n.natsConn.Close()
}
