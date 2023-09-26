package usecases

import (
	"testing"
)

func (m *MockRepository) UpdateResponseStatus(id uint64, status string) error {
	return nil
}

func TestExecUpdateResponseStatus(t *testing.T) {
	mockRepo := &MockRepository{}

	usecase := NewUpdateResponseStatus(mockRepo)

	id := uint64(1)
	status := "completed"
	err := usecase.Exec(id, status)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

/*func TestExecUpdateResponseStatusWithError(t *testing.T) {
	mockRepo := &MockRepository{}

	mockRepo.Err = errors.New("some error")

	usecase := NewUpdateResponseStatus(mockRepo)

	id := uint64(1)
	status := "completed"
	err := usecase.Exec(id, status)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}*/
