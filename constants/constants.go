package constants

const (
	InternalRcSuccess = 2

	/* code response app */
	CODE_PREFIX = "PB"

	CODE_SUCCESS             = "PB-200"
	CODE_BAD_REQUEST         = "PB-400"
	CODE_INTERNAL_SERVER     = "PB-500"
	CODE_PENDING             = "PB-PAY-01"
	CODE_UNAUTHORIZED        = "401"
	CODE_UNAUTHORIZED_ACCESS = "PB-403"

	CODE_SUCCESS_MSG             = "Success"
	CODE_BAD_REQUEST_MSG         = "Bad Request"
	CODE_INTERNAL_SERVER_MSG     = "Internal Server Error"
	CODE_PENDING_MSG             = "Transaksi anda sedang di proses"
	CODE_UNAUTHORIZED_MSG        = "Unauthorized"
	CODE_UNAUTHORIZED_ACCESS_MSG = "Unauthorized Access"
)
