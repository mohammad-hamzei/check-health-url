package usecases

import (
	"git.tashilcar.com/core/tashilcar/domain/entities"
	"testing"
)

func (m *MockRepository) Insert(tashilcar *entities.Tashilcar) error {
	// در اینجا می‌توانید عملیات موک‌آپ ریپازیتوری را انجام دهید.
	// در این مثال، هیچ خطایی را برنمی‌گردانیم.
	return nil
}

func TestExecCreateAPI(t *testing.T) {
	mockRepo := &MockRepository{}

	usecase := NewCreateAPI(mockRepo)

	tashilcar := entities.Tashilcar{
		ID:                      0,
		HealthCheck:             false,
		HealthCheckTimeInterval: 10,
		RequestURL:              "TestRequestURL",
		RequestHTTPMethod:       "RequestHTTPMethodTest",
		RequestHeaders:          "RequestHeadersTest",
		RequestBody:             "RequestBodyTest",
		ResponseStatus:          200,
		CreatedAt:               0,
		UpdatedAt:               0,
	}

	err := usecase.Exec(tashilcar)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

}
