package usecases

import (
	"git.tashilcar.com/core/tashilcar/domain/repositories"
)

type DeleteAPI interface {
	Exec(id uint64) error
}

func NewDeleteAPI(r repositories.TashilcarRepository) DeleteAPI {
	return &deleteAPI{
		repo: r,
	}
}

type deleteAPI struct {
	repo repositories.TashilcarRepository
}

func (d deleteAPI) Exec(id uint64) error {
	err := d.repo.DeleteAPI(id)
	if err != nil {
		return err
	}
	return nil
}
