package sender

import (
	"math/rand"

	"github.com/k0kubun/pp"
)

type SMSSender struct{}

func (s SMSSender) Send(message string) int {
	pp.Printf("📱 Отправка SMS: %s\n", message)
	pp.Println("SMS sent")
	return rand.Int()
}
