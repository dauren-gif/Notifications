package sender

import (
	"math/rand"

	"github.com/k0kubun/pp"
)

type EmailSender struct{}

func (e EmailSender) Send(message string) int {
	pp.Printf("📧 Отправка Email: %s\n", message)
	pp.Println("Email sent")
	return rand.Int()
}
