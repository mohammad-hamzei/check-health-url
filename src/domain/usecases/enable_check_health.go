package usecases

import (
	"git.tashilcar.com/core/tashilcar/domain/repositories"
)

type EnableCheckHealth interface {
	Exec(id uint64, enable bool) error
}

func NewEnableCheckHealth(r repositories.TashilcarRepository) EnableCheckHealth {
	return &enableCheckHealth{
		repo: r,
	}
}

type enableCheckHealth struct {
	repo repositories.TashilcarRepository
}

func (c enableCheckHealth) Exec(id uint64, enable bool) error {
	err := c.repo.EnableCheckHealth(id, enable)
	if err != nil {
		return err
	}
	return nil
}
