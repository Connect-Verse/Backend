package request

import (
	"database/sql"
	"time"
)

type CreateUserRequest struct {
	Name     sql.NullString `validate:"required,min=1,max=200" json:"name"`
	Email    string `validate:"required,min=1,max=50" json:"email"`
    Password string `validate:"required,min=1,max=20" json:"password"`
}   
   
type UpdateTaskRequest struct {
	Name  	 sql.NullString `validate:"required,max=200,min=1" json:"name"`
	Email	 string `validate:"required,min=1,max=200" json:"description"`
    Password string `validate:"required,min=1,max=20" json:"password"`
}

type MapRequest struct {
 Image          string `validate:"required,max=200,min=1" json:"image"`
 Tiles	        string `validate:"required,max=200,min=1" json:"tiles"`
 Info		    string `validate:"required,max=200,min=1" json:"info"`
}

type AvatarRequest struct{
	Name    string `validate:"required,max=200,min=1" json:"Name"`
    Image   string `validate:"required,max=200,min=1" json:"image"`
	ExistedFrom time.Time `validate:"required,max=200,min=1" json:"existedFrom"`
}