package data

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/domain"
	"github.com/taciomcosta/kronos/internal/usecases"
)

// NewSqliteRepository returns a Sqlite repository implementation
func NewSqliteRepository() domain.Repository {
	return &sqliteRepository{}
}

type sqliteRepository struct{}

// CreateJob creates a job.
func (r *sqliteRepository) CreateJob(job *domain.Job) error {
	stmt, err := db.Prepare("INSERT INTO job VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.Exec(job.Name, job.Command, job.Tick)
	return err
}

// CountJobs counts the total of jobs.
func (r *sqliteRepository) CountJobs() int {
	var count int
	stmt, _ := db.Prepare("SELECT COUNT(*) FROM job")
	defer stmt.Close()
	stmt.Step()
	stmt.Scan(&count)
	return count
}

// FindJobs finds all jobs.
func (r *sqliteRepository) FindJobs() []domain.Job {
	stmt, err := db.Prepare("SELECT * FROM job")
	if err != nil {
		return []domain.Job{}
	}
	defer stmt.Close()
	return r.readAllJobs(stmt)
}

func (r *sqliteRepository) readAllJobs(stmt *sqlite3.Stmt) []domain.Job {
	jobs := make([]domain.Job, 0)
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		job := r.readOneJob(stmt)
		jobs = append(jobs, job)
	}
	return jobs
}

func (r *sqliteRepository) readOneJob(stmt *sqlite3.Stmt) domain.Job {
	request := usecases.CreateJobRequest{}
	stmt.Scan(&request.Name, &request.Command, &request.Tick)
	job, _ := domain.NewJob(request.Name, request.Command, request.Tick)
	return job
}
