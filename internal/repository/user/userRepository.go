package user

import (
	"github.com/connect-verse/internal/models"
)

type UserRepository interface{

	FindAll() ([]models.User,error)
	Create(user models.User) (models.User,error)
	FindbyId(userId string) (models.User , error)
	FindbyEmail(userEmail string) (models.User , error)
    Update(user models.User) (users models.User,err error)
	Delete(userId string) (err error)
}






