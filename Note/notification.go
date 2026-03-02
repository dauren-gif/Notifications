package note

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type NotificationSender interface {
	Send(message string) int //получаем id
}

type NotificationModule struct {
	// notificationSender NotificationSender
	history map[int]NotificationInfo //key=ID
	index   map[string][]int
	//хранит историю отправленных уведомлений
	//способ отправки
	getSender func() NotificationSender //функция если понадобится на ходу менять sender
}

func NewNotificationModule(getSender func() NotificationSender) *NotificationModule {
	return &NotificationModule{
		getSender: getSender,
		history:   make(map[int]NotificationInfo),
		index:     make(map[string][]int),
	}
}

func (m NotificationModule) SendNotification(recipient, message string) int { //тут не нужен указатель потому что мапа под капотом содержит указатель
	id := m.getSender().Send(message) //новый id

	Info := NotificationInfo{
		Message:   message,
		Recipient: recipient,
		Status:    "sent",
		Sender:    fmt.Sprintf("%T", m.getSender()),
	}

	m.history[id] = Info //записали через структуру в мапу
	m.index[recipient] = append(m.index[recipient], id)
	return id
}

func (m NotificationModule) GetHistory() map[int]NotificationInfo { //все отправленные уведомления
	tempMap := make(map[int]NotificationInfo, len(m.history))

	for k, v := range m.history {
		tempMap[k] = v
	}
	return tempMap
}

func (m NotificationModule) GetByRecipient(recipient string) []NotificationInfo { //фильтр по получателю
	recInfo := []NotificationInfo{}
	list, ok := m.index[recipient]
	if !ok {
		return []NotificationInfo{} //пустая структура
	}
	for _, v := range list {
		recInfo = append(recInfo, m.history[v])
	}
	return recInfo
}

func (m NotificationModule) Count() { //кол-во отправленных
	pp.Println(len(m.history)) 
}
