package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	PONG                = &Errno{Code: 0, Message: "pong"}
	SayHello            = &Errno{Code: 0, Message: "Hello,World!"}
	Err404              = &Errno{Code: 404, Message: "Err404"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrConf = &Errno{Code: 10100, Message: "Config Error."}
	//SDK errors
	ErrNewSdk     = &Errno{Code: 10101, Message: "NewSdk error."}
	ErrExecuteSdk = &Errno{Code: 10102, Message: "ExecuteSdk error."}

	// 验证errors
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	// Token验证errors
	ErrToken = &Errno{Code: 20002, Message: "Error occurred while signing the token."}
	// 找不到errors
	ErrNotFound = &Errno{Code: 20003, Message: "The item was NotFound."}

	// 数据库查询errors
	ErrDatabase = &Errno{Code: 20100, Message: "Database error."}
	//user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	ErrValidationMobile = &Errno{Code: 20105, Message: "Validation failed(Mobile)."}
	ErrValidationEmail  = &Errno{Code: 20106, Message: "Validation failed(Email)."}
	ErrCreateUUID       = &Errno{Code: 20107, Message: "Error occurred while creating the user UUID."}

	// 授权管理 errors
	ErrAuthTokenCreate = &Errno{Code: 30001, Message: "AuthTokenCreate Error."}
	ErrAuthTokenGet    = &Errno{Code: 30001, Message: "AuthTokenGet Error."}
	// Cookie管理 errors
	ErrCookieGet = &Errno{Code: 30001, Message: "CookieGet Error."}

	// 文件上传管理
	ErrTooBigFile = &Errno{Code: 50001, Message: "The default memory allocation is 10M"}
	// BadRequest
	ErrBadRequest = &Errno{Code: 50002, Message: "BadRequest"}
)
