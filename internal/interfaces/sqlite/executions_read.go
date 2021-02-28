package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindExecutionsResponse returns all executions in FindExecutionsResponse format
func (wr *WriterReader) FindExecutionsResponse(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	if request.JobName != "" {
		return wr.findOne(request)
	}
	return wr.findAll(request)
}

func (wr *WriterReader) findAll(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	sql := "SELECT * FROM execution ORDER BY date DESC LIMIT ?"
	stmt, err := db.Prepare(sql, request.Last)
	if err != nil {
		return uc.FindExecutionsResponse{}
	}
	response := wr.readExecutionsResponse(stmt)
	_ = stmt.Close()
	return response
}

func (wr *WriterReader) findOne(request uc.FindExecutionsRequest) uc.FindExecutionsResponse {
	sql := "SELECT * FROM execution WHERE job_name = ? ORDER BY date DESC LIMIT ?"
	stmt, err := db.Prepare(sql, request.JobName, request.Last)
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
