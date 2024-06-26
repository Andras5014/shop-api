package response

import "time"

type UserResponse struct {
	Id       int32     `json:"id"`
	NickName string    `json:"nick_name"`
	Birthday time.Time `json:"birthday"`
	Mobile   string    `json:"mobile"`
	Gender   string    `json:"gender"`
}
