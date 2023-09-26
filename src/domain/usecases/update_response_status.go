package usecases

import "git.tashilcar.com/core/tashilcar/domain/repositories"

type UpdateResponseStatus interface {
	Exec(id uint64, status uint16) error
}

func NewUpdateResponseStatus(r repositories.TashilcarRepository) UpdateResponseStatus {
	return &updateResponseStatus{
		repo: r,
	}
}

type updateResponseStatus struct {
	repo repositories.TashilcarRepository
}

func (c updateResponseStatus) Exec(id uint64, status uint16) error {
	err := c.repo.UpdateResponseStatus(id, status)
	if err != nil {
		return err
	}
	return nil
}
