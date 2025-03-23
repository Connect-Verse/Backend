package avatars

import "github.com/connect-verse/internal/models"

type AvatarsRepository interface{
  CreateAvatar(avatar models.Avatars) (models.Avatars, error)
  DeleteAvatar(avatarId string) (models.Avatars, error)
  FindAvatar(avatarId string) (models.Avatars, error)
  UpdateAvatar(updAvatar models.Avatars) (models.Avatars, error)
  FindAllAvatar() ([]models.Avatars, error)
}