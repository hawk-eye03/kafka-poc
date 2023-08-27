package consumers

type DailyReminderConsumer struct {
	BaseConsumer
}

func (d *DailyReminderConsumer) GetTopicName() string {
	return "watsapp-msg-received2"
}

func (d *DailyReminderConsumer) ProcessMessage(string) {

}
