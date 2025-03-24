package services

import (
	"errors"
	"fmt"
	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/data/response"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/user"
	"github.com/go-playground/validator/v10"
)


type UserSerImpl struct {
	UserRepo   user.UserRepository
	Validate   *validator.Validate
 }

type NullString struct{
	String string
	Valid bool
}

func NewUserServiceImp(user user.UserRepository, validate *validator.Validate) (service *UserSerImpl,err error) {
   if validate==nil{
	fmt.Println("validate cannot be nil")
	return nil,errors.New("validator instance cannot be nil")
}
	
	return &UserSerImpl{
	UserRepo: user,
	Validate: validate,
   }, err
}


func (u *UserSerImpl)  FindAll() (users []response.UserResponse,err error){
	result,err := u.UserRepo.FindAll()


	if err !=nil {
		return nil,err
	}

	for _ , value := range result{
	
       user:= response.UserResponse{
		Name: *value.Name,
		Email: value.Email,
		CreatedAt: value.CreatedAt,
		UpdatedAt: value.UpdatedAt,
	   }

	   users= append(users, user)
	}

	return users,nil
}


func (u *UserSerImpl) Create(user request.CreateUserRequest) error{
	err:= u.Validate.Struct(user)
    
	if err!=nil {
		return err
	}
	
    
    ur:= models.User{
		Name: &user.Name,
		Password: user.Password,
		Email: user.Email,
	}
    
	err=u.UserRepo.Create(ur)
	if err!=nil {
		return err
	}
	
	return nil
}

func (u *UserSerImpl)  FindbyEmail(email string) ( response.UserResponse , error){
	result,err := u.UserRepo.FindbyEmail(email)


	if err !=nil {
		return response.UserResponse{},err
	}

	

	return response.UserResponse{
		Id: result.Id,
		Name: *result.Name,
		Email: result.Email,
	},nil
}
