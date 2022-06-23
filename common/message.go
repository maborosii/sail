package common

/*
decorate message
*/
type IMessage interface {
	GetName() (string, error)
	GetResource() (string, error)
	GetStatus() (MessageStatus, error)

	InitStatus()
	ConvertStatus(MessageStatus) error
}
