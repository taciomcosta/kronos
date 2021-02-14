package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/entities"
	"github.com/taciomcosta/kronos/internal/usecases"
)

// NewWriterReader returns a Sqlite writer implementation
func NewWriterReader(name string) *WriterReader {
	newDB(name)
	return &WriterReader{}
}

// WriterReader implements usecase.Writer and usecase.Reader
type WriterReader struct{}

// CreateJob creates a job.
func (r *WriterReader) CreateJob(job *entities.Job) error {
	stmt, err := db.Prepare("INSERT INTO job VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	err = stmt.Exec(job.Name, job.Command, job.Tick)
	_ = stmt.Close()
	return err
}

// CountJobs counts total of jobs
func (r *WriterReader) CountJobs() int {
	var count int
	stmt, _ := db.Prepare("SELECT COUNT(*) FROM job")
	_, _ = stmt.Step()
	_ = stmt.Scan(&count)
	_ = stmt.Close()
	return count
}

// FindJobs finds all jobs.
func (r *WriterReader) FindJobs() []entities.Job {
	stmt, err := db.Prepare("SELECT * FROM job")
	if err != nil {
		return []entities.Job{}
	}
	jobs := r.readAllJobs(stmt)
	_ = stmt.Close()
	return jobs
}

func (r *WriterReader) readAllJobs(stmt *sqlite3.Stmt) []entities.Job {
	jobs := make([]entities.Job, 0)
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		job := r.readOneJob(stmt)
		jobs = append(jobs, job)
	}
	return jobs
}

func (r *WriterReader) readOneJob(stmt *sqlite3.Stmt) entities.Job {
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
