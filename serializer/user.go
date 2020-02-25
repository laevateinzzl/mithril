package serializer

import "mithril/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Account:   user.Account,
		Username:  user.Username,
		Status:    user.Status,
		Avatar:    user.AvatarURL(),
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
