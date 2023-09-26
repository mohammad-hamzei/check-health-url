package usecases

import (
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"testing"
)

// یک ماک‌آپ از ریپازیتوری برای تست ایجاد می‌کنیم.
type MockRepository struct {
	Err error
}

func (m *MockRepository) EnableCheckHealth(id uint64, enable bool) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) DeleteAPI(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) Get() ([]*entities.Tashilcar, error) {
	return []*entities.Tashilcar{}, nil
}

func TestExec(t *testing.T) {
	mockRepo := &MockRepository{}

	usecase := NewGetAPIs(mockRepo)

	result, err := usecase.Exec()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected an empty array, but got %v", result)
	}
}
