package sqlite

import (
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
		job, _ := entities.NewJob(dto.Name, dto.Command, dto.Tick, dto.Status)
		jobs = append(jobs, job)
	}
	return jobs
}

// FindJobsResponse returns all jobs in FindJobsResponse format
func (wr *WriterReader) FindJobsResponse() uc.FindJobsResponse {
	stmt, err := db.Prepare(findAllJobsSQL)
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
		row := wr.readJobRow(stmt)
		dto := mapJobRowToDto(row)
		response.Jobs = append(response.Jobs, dto)
	}
	response.Count = len(response.Jobs)
	return response
}

func mapJobRowToDto(row jobRow) uc.JobDTO {
	return uc.JobDTO{
		Name:    row.name,
		Command: row.command,
		Tick:    row.tick,
		Status:  row.status,
	}
}

func (wr *WriterReader) readJobRow(stmt *sqlite3.Stmt) jobRow {
	var row jobRow
	_ = stmt.Scan(&row.name, &row.command, &row.tick, &row.status)
	return row
}

// FindOneJob finds all jobs.
func (wr *WriterReader) FindOneJob(name string) (entities.Job, error) {
	var row jobRow
	stmt, _ := db.Prepare(findOneJobSQL)
	_ = stmt.Exec(name)
	hasRow, _ := stmt.Step()
	if !hasRow {
		return entities.Job{}, errResourceNotFound
	}
	_ = stmt.Scan(&row.name, &row.command, &row.tick, &row.status)
	return entities.NewJob(row.name, row.command, row.tick, row.status)
}

// DescribeJobResponse finds job in DescribeJobResponse format
func (wr *WriterReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	var r uc.DescribeJobResponse
	stmt, _ := db.Prepare(describeJobSQL)
	_ = stmt.Exec(name, name)
	hasRow, _ := stmt.Step()
	if !hasRow {
		return r, errResourceNotFound
	}
	_ = stmt.Scan(
		&r.Name,
		&r.Command,
		&r.Tick,
		&r.LastExecution,
		&r.Status,
		&r.ExecutionsSucceeded,
		&r.ExecutionsFailed,
		&r.AverageCPU,
		&r.AverageMem,
	)
	return r, nil
}
