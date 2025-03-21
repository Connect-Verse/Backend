package verificationtoken

import (
	"github.com/connect-verse/internal/models"
)

type VerifyRepository interface{

	
	Create(token models.VerificationToken) error
	FindbyId(tokenId string) (models.VerificationToken , error)
	FindbyEmail(email string) (models.VerificationToken , error)
	Update(token models.VerificationToken) ( models.VerificationToken, error)
	Delete(tokenId string) ( error)
}

