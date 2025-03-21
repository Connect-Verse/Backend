package response

import (
	"time"
	"database/sql"
)

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type UserResponse struct {
	Id  	string    `json:"id"`
	Name 	sql.NullString `json:"name"`
	Email 	string		`json:"email"`
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time	 `json:"updatedAt"`
}

type VerificationToken struct {
	
}