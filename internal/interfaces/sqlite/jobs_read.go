package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindJobs finds all jobs.
func (wr *WriterReader) FindJobs() []entities.Job {
	stmt, err := db.Prepare(findAllJobsSQL)
	if err != nil {
		return []entities.Job{}
	}
	rows := wr.readAllJobRows(stmt)
	_ = stmt.Close()
	return wr.mapJobRowsToEntitities(rows)
}

func (wr *WriterReader) readAllJobRows(stmt *sqlite3.Stmt) []jobRow {
	var rows []jobRow
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		row := wr.readJobRow(stmt)
		rows = append(rows, row)
	}
	return rows
}

func (wr *WriterReader) readJobRow(stmt *sqlite3.Stmt) jobRow {
	var row jobRow
	_ = stmt.Scan(&row.name, &row.command, &row.tick, &row.status)
	return row
}

func (wr *WriterReader) mapJobRowsToEntitities(rows []jobRow) []entities.Job {
	var jobs []entities.Job
	for _, row := range rows {
		job, _ := entities.NewJob(row.name, row.command, row.tick, row.status)
		jobs = append(jobs, job)
	}
	return jobs
}

// FindOneJob finds all jobs.
func (wr *WriterReader) FindOneJob(name string) (entities.Job, error) {
	stmt, _ := db.Prepare(findOneJobSQL)
	_ = stmt.Exec(name)
	if hasRow, _ := stmt.Step(); !hasRow {
		return entities.Job{}, errResourceNotFound
	}
	row := wr.readJobRow(stmt)
	return entities.NewJob(row.name, row.command, row.tick, row.status)
}

// DescribeJobResponse finds job in DescribeJobResponse format
func (wr *WriterReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	var r uc.DescribeJobResponse
	stmt, _ := db.Prepare(describeJobSQL)
	_ = stmt.Exec(name, name)
	if hasRow, _ := stmt.Step(); !hasRow {
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
	r.AssignedNotifiers = readAssignedNotifiers(name)
	return r, nil
}

func readAssignedNotifiers(name string) []string {
	assignedNotifiers := []string{}
	stmt, _ := db.Prepare(describeJobAssignedNotifiersSQL)
	_ = stmt.Exec(name)
	var notifier string
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		_ = stmt.Scan(&notifier)
		assignedNotifiers = append(assignedNotifiers, notifier)
	}
	return assignedNotifiers
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
		dto := wr.readJobDTO(stmt)
		response.Jobs = append(response.Jobs, dto)
	}
	response.Count = len(response.Jobs)
	return response
}

func (wr *WriterReader) readJobDTO(stmt *sqlite3.Stmt) uc.JobDTO {
	var dto uc.JobDTO
	_ = stmt.Scan(&dto.Name, &dto.Command, &dto.Tick, &dto.Status)
	return dto
}
