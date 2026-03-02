package main

import (
	note "notification/Note"
	"notification/Note/sender"

	"github.com/k0kubun/pp"
)

func main(){
	userPrefersSMS := false//внешняя реализация

	//пробный обьект
	trial := note.NewNotificationModule(func() note.NotificationSender {
		if userPrefersSMS {
			return sender.SMSSender{}//любой сендер любая функция
		}
		return sender.EmailSender{}
	})

	trial.SendNotification("Dauren", "Hi! How are you?")
	trial.SendNotification("Darkhan", "Hi! asdfasfs")

	pp.Println(trial.GetHistory())

	trial.Count()
}