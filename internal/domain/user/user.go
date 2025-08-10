package user

import "github.com/Haya372/go-samples/internal/domain/vo"

type User struct {
	Id       vo.UserId
	Password vo.Password

	Name string
}
