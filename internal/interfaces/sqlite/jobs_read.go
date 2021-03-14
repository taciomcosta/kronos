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
		return dto, errResourceNotFound
	}
	_ = stmt.Scan(&dto.Name, &dto.Command, &dto.Tick)
	return entities.NewJob(dto.Name, dto.Command, dto.Tick)
}

// DescribeJobResponse finds job in DeDescribeJobResponse format
func (wr *WriterReader) DescribeJobResponse(name string) (uc.DescribeJobResponse, error) {
	r := uc.DescribeJobResponse{}
	stmt, _ := db.Prepare(
		`SELECT
			MAX(j.name) AS name,
			MAX(j.command) as command,
			MAX(j.tick) AS tick,
			MAX(e.date) AS last_execution,
			'enabled' AS status,
			COUNT(CASE e.STATUS WHEN 'Succeeded' THEN 1 ELSE null END) AS executions_succeeded,
			COUNT(CASE e.STATUS WHEN 'Failed' THEN 1 ELSE null END) AS executions_failed,
			AVG(e.cpu_time) AS average_cpu,
			AVG(e.mem_usage) AS average_mem
		 FROM execution e
		 INNER JOIN job j ON j.name = e.job_name
		 GROUP BY e.job_name
		 WHERE e.name = ?`,
	)
	_ = stmt.Exec(name)
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
