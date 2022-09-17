package repository

import (
	"fmt"
	"testing"

	"github.com/perfectzoo/cat-todolist/user/internal/service"
)

func TestUser_Create(t *testing.T) {
	InitDB()
	f := new(User)
	req := new(service.UserRequest)
	req.UserName = "FanOne"
	req.NickName = "CocaineCong"
	req.Password = "12345678"
	req.PasswordConfirm = "12345678"
	err := f.Create(req)
	fmt.Println(err)
}
