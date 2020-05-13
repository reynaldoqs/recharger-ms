package msgbroker

// MessageBroker defines our interface for connecting, producing and consuming messages
type MessageBroker interface {
	PublishOnQueue(body []byte, queueName string) error
	Subscribe(exchangeName string, handlerFunc func(data []byte)) error
	Close()
}

/*
// Defines our interface for connecting, producing and consuming messages.
type IMessagingClient interface {
	ConnectToBroker(connectionString string)
	Publish(msg []byte, exchangeName string, exchangeType string) error
	PublishOnQueue(msg []byte, queueName string) error
	Subscribe(exchangeName string, exchangeType string, consumerName string, handlerFunc func(amqp.Delivery)) error
	SubscribeToQueue(queueName string, consumerName string, handlerFunc func(amqp.Delivery)) error
	Close()
}
*/
