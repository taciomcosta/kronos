package data

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/domain"
)

func NewRepository() domain.Repository {
	return &sqliteRepository{}
}

type sqliteRepository struct{}

func (r *sqliteRepository) CreateJob(job *domain.Job) error {
	stmt, err := db.Prepare("INSERT INTO job VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.Exec(job.Name, job.Command, job.Tick)
	return err
}

func (r *sqliteRepository) FindJobs() []domain.Job {
	stmt, err := db.Prepare("SELECT * FROM job")
	if err != nil {
		return []domain.Job{}
	}
	defer stmt.Close()
	return r.scanAllJobs(stmt)
}

func (r *sqliteRepository) scanAllJobs(stmt *sqlite3.Stmt) []domain.Job {
	jobs := make([]domain.Job, 0)
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		job := r.readOneJob(stmt)
		jobs = append(jobs, job)
	}
	return jobs
}

func (r *sqliteRepository) readOneJob(stmt *sqlite3.Stmt) domain.Job {
	request := domain.CreateJobRequest{}
	stmt.Scan(&request.Name, &request.Command, &request.Tick)
	job, _ := domain.NewJob(request)
	return job
}
