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
	err = stmt.Exec(nil, job.Name, job.Command, job.Tick)
	return err
}

func (r *sqliteRepository) FindJobs() []domain.Job {
	stmt, err := db.Prepare("SELECT * FROM job")
	if err != nil {
		return []domain.Job{}
	}
	defer stmt.Close()
	jobs := make([]domain.Job, 0)
	r.scanAllJobs(jobs, stmt)
	return jobs
}

func (r *sqliteRepository) scanAllJobs(jobs []domain.Job, stmt *sqlite3.Stmt) {
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		job := r.scanOneJob(stmt)
		jobs = append(jobs, job)
	}
}

func (r *sqliteRepository) scanOneJob(stmt *sqlite3.Stmt) domain.Job {
	request := domain.CreateJobRequest{}
	stmt.Scan(&request.Name, &request.Command, &request.Tick)
	job, _ := domain.NewJob(request)
	return job
}
