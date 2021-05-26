package response

const (
	Success = 200
	//權限相關
	NoAuthority    = 100
	ActionNotFound = 101

	AccountNotFound      = 110
	AccountNotActive     = 111
	AccountAlreadyActive = 112
	AccountAlreadyExist  = 113

	LoginInfoWithoutCaptcha = 117
	LoginCaptchaWrong       = 118
	AccountLoginFail        = 119

	RoleGroupAlreadyExist = 121

	RequestIdWrong   = 131
	RequestIdExpired = 132
	RequestIdUsed    = 133

	//資料處理相關
	DataNotFound      = 401
	KeyValueDuplicate = 402
	DataIsLocked      = 403
	CustomerNotFound  = 404

	//其他
	Exception = 999
)
