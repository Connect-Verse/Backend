package authservice

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/connect-verse/internal/data/request"
	"github.com/connect-verse/internal/models"
	"github.com/connect-verse/internal/repository/user"
	"github.com/connect-verse/internal/repository/verificationToken"
	"github.com/go-playground/validator/v10"
)

type AuthServiceImplementation struct{
   User         user.UserRepository
   Verification verificationtoken.VerifyRepository
   Validate     *validator.Validate
}

func NewAuthServiceImplementation(user user.UserRepository, verification verificationtoken.VerifyRepository, validate *validator.Validate) (service *AuthServiceImplementation, err error){
    if validate==nil{
		return nil,errors.New("validator is not provided in this")
	}
	return &AuthServiceImplementation{
		User: user,
		Verification: verification,
		Validate: validate,
	}, nil
}

type LoginResponse struct{
	Id  	 string
	Name 	 sql.NullString
	Password string
	Email    string
}

type SignUpResponse struct{

}

func (a *AuthServiceImplementation) Login(user request.CreateUserRequest) (LoginResponse,error){
    result,err:= a.User.FindbyEmail(user.Email)

	if err!=nil {
		return LoginResponse{},err
	}
     res:= LoginResponse{
		Id: result.Id,
	//	Name: result.Name,
		Password: result.Password,
		Email: result.Email,
	 }
	return res,nil
}



func (a *AuthServiceImplementation) SignUp(user request.CreateUserRequest,token string) error{
    req:=models.User{
      Email: user.Email,
	  Password: "sogjei",
	  Name: &user.Name,
	}
	err:= a.User.Create(req)
   
	if err!=nil{
		log.Print(err.Error(),"this is error")
		return err
	}

	 err = a.Verification.Create(models.VerificationToken{
	 	EmailIdentifier: user.Email,
	 	Token:token,
	 	ExpiresAt: time.Now().Add(24 * time.Hour),
	})


	if err!=nil{
		log.Print(err.Error())
		return err
	}


	return nil
}



func (a *AuthServiceImplementation) Verify(email string,token string) error {
   result,err:= a.Verification.FindbyEmail(email)

   if err!=nil{
	  return err
   }

   if token==result.Token {
	err= a.Verification.Delete(result.Id)
	if err!=nil {
		return err
	}
	  return nil
   }
   
   return errors.New("token provide is not matching with token stored in the database")  
}