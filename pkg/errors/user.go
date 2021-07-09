package errors

var (
	LoginFail = NewError(401, "LoginFail")
	CreateUserFail  = NewError(400, "CreateUserFail")
	UpdateUserFail  = NewError(400, "UpdateUserFail")
	DeleteUserFail  = NewError(400, "DeleteUserFail")
)
