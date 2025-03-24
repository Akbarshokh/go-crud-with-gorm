package status

var (
	NoError                  = 0
	ErrorCodeValidation      = -10
	ErrorAuthorization       = -11
	ErrorInternalServerError = -12
	ErrorBadRequest          = -13
	ErrorNotFound            = -14
	ErrorCreateFailed        = -15
	ErrorUpdateFailed        = -16
	ErrorDeleteFailed        = -17
	ErrorBindFailed          = -18
	ErrorUnauthorizedAccess  = -19
)

var (
	Success = "Success"
	Failure = "Failure"
)
