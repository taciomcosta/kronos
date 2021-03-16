package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var pageSize = 10

// FindExecutionsResponse returns all executions in FindExecutionsResponse format
func (wr *WriterReader) FindExecutionsResponse(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	if request.JobName == "" {
		return wr.findExecutionsResponse(
			findiAllExecutions,
			pageSize, request.Page*pageSize,
		)
	}
	return wr.findExecutionsResponse(
		findJobExecutions,
		request.JobName, pageSize, request.Page*pageSize,
	)
}

func (wr *WriterReader) findExecutionsResponse(sql string, args ...interface{}) uc.FindExecutionsResponse {
	stmt, err := db.Prepare(sql, args...)
	if err != nil {
		return uc.FindExecutionsResponse{}
	}
	response := wr.readExecutionsResponse(stmt)
	_ = stmt.Close()
	return response
}

func (wr *WriterReader) readExecutionsResponse(stmt *sqlite3.Stmt) uc.FindExecutionsResponse {
	var response uc.FindExecutionsResponse
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		execution := wr.readExecutionDTO(stmt)
		response.Executions = append(response.Executions, execution)
	}
	return response
}

func (wr *WriterReader) readExecutionDTO(stmt *sqlite3.Stmt) uc.ExecutionDTO {
	execution := uc.ExecutionDTO{}
	_ = stmt.Scan(
		&execution.JobName,
		&execution.Date,
		&execution.Status,
		&execution.MemUsage,
		&execution.CPUTime,
	)
	return execution
}
