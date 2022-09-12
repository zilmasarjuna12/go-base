package request

type (
	PostUser struct {
		Name     string `json:"name"`
		Fullname string `json:"fullname"`
	}
)
