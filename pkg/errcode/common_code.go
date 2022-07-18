package errcode

var (
	Success                 = NewError(1000, "ok")
	ServerError             = NewError(1001, "internal error")
	NotFound                = NewError(1002, "not found")
	RequestTypeNotSupport   = NewError(1003, "not support message type")
	RequestSourceNotSupport = NewError(1003, "not support message source")
	RequestTypeNotAdapt     = NewError(1004, "not adaptable message type")
	BadRequest              = NewError(1005, "bad request")
)
