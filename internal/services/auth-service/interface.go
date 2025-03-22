package authservice

import (
	"github.com/connect-verse/internal/data/request"
)

type AuthService interface{
	Login(user request.CreateUserRequest) (LoginResponse,error)
	SignUp(user request.CreateUserRequest,token string) error
	Verify(email string,tokenId string) error
}