package sqlite

import (
	"github.com/bvinc/go-sqlite-lite/sqlite3"
	"github.com/taciomcosta/kronos/internal/entities"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

// FindNotifiersResponse returns all notifiers in FindNotifiersResponse format
func (wr *WriterReader) FindNotifiersResponse() uc.FindNotifiersResponse {
	stmt, err := db.Prepare(findAllNotifiersSQL)
	if err != nil {
		return uc.FindNotifiersResponse{}
	}
	response := wr.readAllNotifiersResponse(stmt)
	_ = stmt.Close()
	return response
}

func (wr *WriterReader) readAllNotifiersResponse(stmt *sqlite3.Stmt) uc.FindNotifiersResponse {
	var response uc.FindNotifiersResponse
	for hasRow, _ := stmt.Step(); hasRow; hasRow, _ = stmt.Step() {
		notifier := wr.readNotifierDTO(stmt)
		response.Notifiers = append(response.Notifiers, notifier)
	}
	response.Count = len(response.Notifiers)
	return response
}

func (wr *WriterReader) readNotifierDTO(stmt *sqlite3.Stmt) uc.NotifierDTO {
	notifier := uc.NotifierDTO{}
	var nType int
	_ = stmt.Scan(&notifier.Name, &nType)
	notifier.Type = notifierTypeToString(entities.NotifierType(nType))
	return notifier
}

func notifierTypeToString(nType entities.NotifierType) string {
	switch nType {
	case entities.SlackNotifierType:
		return "slack"
	default:
		return "unknown"
	}
}

// FindOneNotifier finds all notifiers
func (wr *WriterReader) FindOneNotifier(name string) (entities.Notifier, error) {
	// TODO: implement query
	return entities.Notifier{}, nil
}
