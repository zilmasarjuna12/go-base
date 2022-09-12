package response

import (
	"fmt"
	constants "go-base/constants"
	"net/http"
)

type BaseResponse struct {
	ResponseCode    string           `json:"response_code"`
	ResponseMessage string           `json:"response_message"`
	DeviceID        string           `json:"device_id,omitempty"`
	TraceID         string           `json:"trace_id"`
	Meta            BaseResponseMeta `json:"meta"`
}

type BaseResponseMeta struct {
	DebugParam string `json:"debug_param"`
}

// New response
func NewBaseResponse(code string, response *BaseResponse, err error) {
	switch code {
	case constants.CODE_INTERNAL_SERVER:
		response.ResponseCode = constants.CODE_INTERNAL_SERVER
		response.ResponseMessage = constants.CODE_INTERNAL_SERVER_MSG
	case constants.CODE_BAD_REQUEST:
		response.ResponseCode = constants.CODE_BAD_REQUEST
		response.ResponseMessage = constants.CODE_BAD_REQUEST_MSG
	case constants.CODE_SUCCESS:
		response.ResponseCode = constants.CODE_SUCCESS
		response.ResponseMessage = constants.CODE_SUCCESS_MSG
	case constants.CODE_PENDING:
		response.ResponseCode = constants.CODE_PENDING
		response.ResponseMessage = constants.CODE_PENDING_MSG
	case constants.CODE_UNAUTHORIZED:
		response.ResponseCode = constants.CODE_UNAUTHORIZED
		response.ResponseMessage = constants.CODE_UNAUTHORIZED_MSG
	case constants.CODE_UNAUTHORIZED_ACCESS:
		response.ResponseCode = constants.CODE_UNAUTHORIZED_ACCESS
		response.ResponseMessage = constants.CODE_UNAUTHORIZED_ACCESS_MSG
	default:
		response.ResponseCode = constants.CODE_INTERNAL_SERVER
		response.ResponseMessage = constants.CODE_INTERNAL_SERVER_MSG
	}

	// mapping to debug param
	if err != nil {
		response.Meta = BaseResponseMeta{
			DebugParam: fmt.Sprintf("%v", err),
		}
	}
}

// New response
func NewBaseResponseStatusCode(code int, response *BaseResponse, err error) {
	switch code {
	case http.StatusInternalServerError:
		response.ResponseCode = constants.CODE_INTERNAL_SERVER
		response.ResponseMessage = constants.CODE_INTERNAL_SERVER_MSG
	case http.StatusBadRequest:
		response.ResponseCode = constants.CODE_BAD_REQUEST
		response.ResponseMessage = constants.CODE_BAD_REQUEST_MSG
	case http.StatusUnauthorized:
		response.ResponseCode = constants.CODE_UNAUTHORIZED
		response.ResponseMessage = constants.CODE_UNAUTHORIZED_MSG
	case http.StatusOK:
		response.ResponseCode = constants.CODE_SUCCESS
		response.ResponseMessage = constants.CODE_SUCCESS_MSG
	default:
		response.ResponseCode = constants.CODE_INTERNAL_SERVER
		response.ResponseMessage = constants.CODE_INTERNAL_SERVER_MSG
	}

	// mapping to debug param
	if err != nil {
		response.Meta = BaseResponseMeta{
			DebugParam: fmt.Sprintf("%v", err),
		}
	}
}
