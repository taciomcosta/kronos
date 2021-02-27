package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindExecutionsResponse returns all executions in FindExecutionsResponse format
func (wr *WriterReader) FindExecutionsResponse() uc.FindExecutionsResponse {
	stmt, err := db.Prepare("SELECT * FROM execution")
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
		&execution.CPUUsage,
		&execution.NetIn,
		&execution.NetOut,
	)
	return execution
}
