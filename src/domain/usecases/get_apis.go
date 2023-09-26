package usecases

import (
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"git.tashilcar.com/core/tashilcar/domain/repositories"
)

type GetAPIs interface {
	Exec() ([]*entities.Tashilcar, error)
}

func NewGetAPIs(r repositories.TashilcarRepository) GetAPIs {
	return &getAPIs{
		repo: r,
	}
}

type getAPIs struct {
	repo repositories.TashilcarRepository
}

func (c getAPIs) Exec() ([]*entities.Tashilcar, error) {
	urls, err := c.repo.Get()
	if err != nil {
		return nil, err
	}
	return urls, nil
}
