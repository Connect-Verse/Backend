package maps

import "github.com/connect-verse/internal/models"

type Maps interface{
  createMap(maps models.Maps) error
  deleteMap(mapId string) error
  findMap(mapId string) (mapp models.Maps,err error)
  updateMap(mapp models.Maps, err error)
  findAllMaps() (maps []models.Maps ,err error)
}