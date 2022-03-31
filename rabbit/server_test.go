package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"testing"
)

const rabbitURL = "amqp://guest:guest@localhost:5672/"

func TestServer(t *testing.T) {
	expect := assert.New(t)
	conn, err := amqp.Dial(rabbitURL)
	expect.Nil(err)
	defer conn.Close()

	ch, err := conn.Channel()
	expect.Nil(err)
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"wwjtest",
		false,
		false,
		false,
		false,
		nil,
	)
	expect.Nil(err)

	msgCh, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	expect.Nil(err)

	for msg := range msgCh {
		fmt.Println(msg)
		msg.Ack(false)
	}
}
