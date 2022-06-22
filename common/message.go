package common

/*
decorate message
*/
type IMessage interface {
	GetName() string
	GetResource() string

	Tag() string
	ConvertStatus() error
}
