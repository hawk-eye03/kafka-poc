package consumers

type Consumer interface {
	GetTopicName() []string
	ProcessMessage(string)
}
