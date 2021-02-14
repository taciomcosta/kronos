package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/entities"
	"github.com/taciomcosta/kronos/internal/usecases"
)

// NewWriter returns a Sqlite writer implementation
func NewWriter(name string) usecases.Writer {
	newDB(name)
	return &sqliteWriter{}
}

type sqliteWriter struct{}

// CreateJob creates a job.
func (r *sqliteWriter) CreateJob(job *entities.Job) error {
	stmt, err := db.Prepare("INSERT INTO job VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	err = stmt.Exec(job.Name, job.Command, job.Tick)
	_ = stmt.Close()
	return err
}

func (r *sqliteWriter) CountJobs() int {
	var count int
	stmt, _ := db.Prepare("SELECT COUNT(*) FROM job")
	_, _ = stmt.Step()
	_ = stmt.Scan(&count)
	_ = stmt.Close()
	return count
}

// FindJobs finds all jobs.
func (r *sqliteWriter) FindJobs() []entities.Job {
	stmt, err := db.Prepare("SELECT * FROM job")
	if err != nil {
		return []entities.Job{}
	}
	jobs := r.readAllJobs(stmt)
	_ = stmt.Close()
	return jobs
}

func (r *sqliteWriter) readAllJobs(stmt *sqlite3.Stmt) []entities.Job {
	jobs := make([]entities.Job, 0)
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		job := r.readOneJob(stmt)
		jobs = append(jobs, job)
	}
	return jobs
}

func (r *sqliteWriter) readOneJob(stmt *sqlite3.Stmt) entities.Job {
	request := usecases.CreateJobRequest{}
	_ = stmt.Scan(&request.Name, &request.Command, &request.Tick)
	// TODO: add usecases.NewJob() so sqlite doesn't have to know about usecase.Host
	job, _ := entities.NewJob(
		request.Name,
		request.Command,
		request.Tick,
		usecases.GetHost().GetDettachedStream())
	return job
}
