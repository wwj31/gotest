package rabbit

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient(t *testing.T) {
	expect := assert.New(t)
	conn, err := amqp.Dial(rabbitURL)
	expect.Nil(err)
	defer conn.Close()

	ch, err := conn.Channel()
	expect.Nil(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"wwjtest",
		false,
		false,
		false,
		false,
		nil,
	)
	expect.Nil(err)
	for {
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: "abcd",
			ReplyTo:       "fjwofjeow",
			Body:          []byte("hello world"),
		})
		expect.Nil(err)
	}
}
