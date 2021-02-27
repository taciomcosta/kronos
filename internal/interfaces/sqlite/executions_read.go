package sqlite

import uc "github.com/taciomcosta/kronos/internal/usecases"

// FindExecutionsResponse returns all executions in FindExecutionsResponse format
func (wr *WriterReader) FindExecutionsResponse() uc.FindExecutionsResponse {
	return uc.FindExecutionsResponse{}
	//stmt, err := db.Prepare("SELECT * FROM job")
	//if err != nil {
	//return uc.FindJobsResponse{}
	//}
	//response := wr.readAllJobsResponse(stmt)
	//_ = stmt.Close()
	//return response
}
