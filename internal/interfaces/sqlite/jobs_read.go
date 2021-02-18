package sqlite

import (
	"errors"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindJobs finds all jobs.
func (wr *WriterReader) FindJobs() []entities.Job {
	response := wr.FindJobsResponse()
	return wr.mapDTOToEntitities(response.Jobs)
}

func (wr *WriterReader) mapDTOToEntitities(dtos []uc.JobDTO) []entities.Job {
	var jobs []entities.Job
	for _, dto := range dtos {
		job, _ := entities.NewJob(dto.Name, dto.Command, dto.Tick)
		jobs = append(jobs, job)
	}
	return jobs
}

// FindJobsResponse returns all jobs in FindJobsResponse format
func (wr *WriterReader) FindJobsResponse() uc.FindJobsResponse {
	stmt, err := db.Prepare("SELECT * FROM job")
	if err != nil {
		return uc.FindJobsResponse{}
	}
	response := wr.readAllJobsResponse(stmt)
	_ = stmt.Close()
	return response
}

func (wr *WriterReader) readAllJobsResponse(stmt *sqlite3.Stmt) uc.FindJobsResponse {
	var response uc.FindJobsResponse
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		job := wr.readJobDTO(stmt)
		response.Jobs = append(response.Jobs, job)
	}
	response.Count = len(response.Jobs)
	return response
}

func (wr *WriterReader) readJobDTO(stmt *sqlite3.Stmt) uc.JobDTO {
	job := uc.JobDTO{}
	_ = stmt.Scan(&job.Name, &job.Command, &job.Tick)
	return job
}

// FindOneJob finds all jobs.
func (wr *WriterReader) FindOneJob(name string) (entities.Job, error) {
	var dto entities.Job
	stmt, _ := db.Prepare("SELECT * FROM job WHERE name=?")
	_ = stmt.Exec(name)
	hasRow, _ := stmt.Step()
	if !hasRow {
		return dto, errors.New("resource not found")
	}
	_ = stmt.Scan(&dto.Name, &dto.Command, &dto.Tick)
	return entities.NewJob(dto.Name, dto.Command, dto.Tick)
}
