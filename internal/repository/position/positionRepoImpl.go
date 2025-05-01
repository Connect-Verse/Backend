package position

import "github.com/connect-verse/internal/models"

type PositionRepository interface{
	FindPosition(metaId string) (models.PlayerPosition,error)
	SetPosition(position models.PlayerPosition) (models.PlayerPosition, error)
}

