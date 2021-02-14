package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
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

// FindJobsResponse returns all jobs in FindJobsResponse format
func (r *WriterReader) FindJobsResponse() uc.FindJobsResponse {
	stmt, err := db.Prepare("SELECT * FROM job")
	if err != nil {
		return uc.FindJobsResponse{}
	}
	response := r.readAllJobsResponse(stmt)
	_ = stmt.Close()
	return response
}

func (r *WriterReader) readAllJobsResponse(stmt *sqlite3.Stmt) uc.FindJobsResponse {
	var response uc.FindJobsResponse
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		job := uc.JobDTO{}
		_ = stmt.Scan(&job.Name, &job.Command, &job.Tick)
		response.Jobs = append(response.Jobs, job)
	}
	response.Count = len(response.Jobs)
	return response
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
	request := uc.CreateJobRequest{}
	_ = stmt.Scan(&request.Name, &request.Command, &request.Tick)
	// TODO: add usecases.NewJob() so sqlite doesn't have to know about usecase.Host
	job, _ := entities.NewJob(
		request.Name,
		request.Command,
		request.Tick,
		uc.GetHost().GetDettachedStream())
	return job
}
