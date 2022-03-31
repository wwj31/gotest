package rabbit

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
	"testing"
)

const rabbitURL = "amqp://guest:guest@localhost:5672/"

func TestMain(t *testing.T) {
	expect := assert.New(t)
	conn, err := amqp.Dial(rabbitURL)
	expect.Nil(err)

	ch, err := conn.Channel()
	expect.Nil(err)
}
