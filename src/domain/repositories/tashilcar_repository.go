package repositories

import (
	"git.tashilcar.com/core/tashilcar/domain/entities"
)

type TashilcarRepository interface {
	Insert(tashilcar *entities.Tashilcar) error
	Get() ([]*entities.Tashilcar, error)
	EnableCheckHealth(id uint64, enable bool) error
	DeleteAPI(id uint64) error
	UpdateResponseStatus(id uint64, status uint16) error
}
