package common

/*
decorate message
*/
type InMessage interface {
	GetName() (string, error)
	GetResource() (string, error)
	GetStatus() (MessageStatus, error)

	InitStatus()
	ConvertStatus(MessageStatus) error
}

type OutMessage interface {
	RealText()
}
type OutMessageTemplate interface {
	GetSentence() []string
}
