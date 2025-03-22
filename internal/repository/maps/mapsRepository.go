package maps

import "github.com/connect-verse/internal/models"

type MapsRepository interface{
  CreateMap(mapp models.Maps) (models.Maps, error)
  DeleteMap(mappId string) (models.Maps, error)
  FindMap(mappId string) (models.Maps, error)
  UpdateMap(updMap models.Maps) (models.Maps, error)
  FindAllMaps() ([]models.Maps,error)
}