package errcode

var (
	Success                 = NewError(1000, "ok")
	ServerError             = NewError(1001, "internal error")
	NotFound                = NewError(1002, "not found")
	RequestTypeNotSupport   = NewError(1003, "not support message type")
	RequestSourceNotSupport = NewError(1004, "not support message source")
	RequestTypeNotAdapt     = NewError(1005, "not adaptable message type")
	BadRequest              = NewError(1006, "bad request")
	RequestExpired          = NewError(1007, "message is expired")
)
