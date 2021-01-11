package domain

func NewMockRepository() Repository {
	return &mockRepository{}
}

type mockRepository struct{}

func (mr *mockRepository) CreateJob(job *Job) error {
	return nil
}
