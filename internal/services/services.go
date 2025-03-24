package services

import (
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
)

type UserServices interface{
	 FindAll() (users []response.UserResponse,err error)
	 Create(user request.CreateUserRequest) error
	 FindbyEmail(email string) ( response.UserResponse , error)
    // Update(user request.UpdateTaskRequest) (error)
	// Delete(userId string) (user response.UserResponse , err error)
	
}
