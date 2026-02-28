package note

type NotificationSender interface {
	Send(message string)string //должно быть еррор но я еще не научился
	GetStatus() string //отправлено или нет
}

type NotificationModule struct {
	notificationSender NotificationSender
	//хранит историю отправленных уведомлений
	//способ отправки
}

func NewNotificationModule(notificationSender NotificationSender) *NotificationModule {
	return &NotificationModule{

	}
}

func(m NotificationModule) SendNotification(recipient, message string)int{//вернуть ID
}

func(m NotificationModule) GetHistory()map[int]NotificationInfo{//все отправленные уведомления
}

func(m NotificationModule) GetByRecipient(recipient string)[]NotificationInfo{ //фильтр по получателю

}

func(m NotificationModule) ChangeSender(sender NotificationSender){ //изменение способа отправки

}
