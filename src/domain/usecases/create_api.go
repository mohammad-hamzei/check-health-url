package usecases

import (
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"git.tashilcar.com/core/tashilcar/domain/repositories"
)

type CreateAPI interface {
	Exec(tashilcar entities.Tashilcar) error
}

func NewCreateAPI(r repositories.TashilcarRepository) CreateAPI {
	return &createAPI{
		repo: r,
	}
}

type createAPI struct {
	repo repositories.TashilcarRepository
}

func (c createAPI) Exec(tashilcar entities.Tashilcar) error {
	err := c.repo.Insert(&tashilcar)
	if err != nil {
		return err
	}
	return nil
}
