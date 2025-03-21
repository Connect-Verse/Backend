package request

import "database/sql"

type CreateUserRequest struct {
	Name     sql.NullString `validate:"required,min=1,max=200" json:"name"`
	Email    string `validate:"required,min=1,max=50" json:"email"`
    Password string `validate:"required,min=1,max=20" json:"password"`
}   
   
type UpdateTaskRequest struct {
	Id    int    `validate:"required"`
	Name  sql.NullString `validate:"required,max=200,min=1" json:"name"`
	Email string `validate:"required,min=1,max=200" json:"description"`
    Password string
}