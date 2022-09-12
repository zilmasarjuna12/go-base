package response

import "go-base/model"

type (
	GetUser struct {
		BaseResponse
		User []model.User `json:"user"`
	}

	PostUser struct {
		BaseResponse
		User model.User `json:"user"`
	}
)
