package usecases

import (
	"testing"
)

func (m *MockRepository) UpdateResponseStatus(id uint64, status uint16) error {
	return nil
}

func TestExecUpdateResponseStatus(t *testing.T) {
	mockRepo := &MockRepository{}

	usecase := NewUpdateResponseStatus(mockRepo)

	id := uint64(1)
	status := uint16(16)
	err := usecase.Exec(id, status)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
