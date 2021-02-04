package domain

// NewMockRepository returns a mock implementation of repository.
func NewMockRepository() Repository {
	return &mockRepository{}
}

type mockRepository struct{}

// CreateJob creates a job.
func (mr *mockRepository) CreateJob(job *Job) error {
	return nil
}

// CountJobs counts the total of jobs.
func (mr *mockRepository) CountJobs() int {
	return 1
}

// FindJobs finds all jobs.
func (mr *mockRepository) FindJobs() []Job {
	return []Job{
		{
			Name:    "list",
			Command: "ls",
			Tick:    "* * * * *",
		},
	}

}
