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
	notifierRow, err := wr.findBaseNotifier(name)
	if err != nil {
		return entities.Notifier{}, err
	}
	slackRow, err := wr.findSlackMetadata(name)
	if err != nil {
		return entities.Notifier{}, err
	}
	return entities.NewNotifier(
		notifierRow.name,
		entities.NotifierType(notifierRow.ntype),
		slackRowToMetadata(slackRow))
}

func (wr *WriterReader) findBaseNotifier(name string) (notifierRow, error) {
	var row notifierRow
	stmt, _ := db.Prepare(findOneNotifierSQL)
	_ = stmt.Exec(name)
	hasRow, _ := stmt.Step()
	if !hasRow {
		return row, errResourceNotFound
	}
	err := stmt.Scan(&row.name, &row.ntype)
	return row, err
}

func (wr *WriterReader) findSlackMetadata(name string) (slackRow, error) {
	var row slackRow
	stmt, _ := db.Prepare(findOneSlackSQL)
	_ = stmt.Exec(name)
	hasRow, _ := stmt.Step()
	if !hasRow {
		return row, errResourceNotFound
	}
	err := stmt.Scan(&row.authToken, &row.channelIds)
	return row, err
}

func slackRowToMetadata(row slackRow) map[string]string {
	metadata := make(map[string]string)
	metadata["auth_token"] = row.authToken
	metadata["channel_ids"] = row.channelIds
	return metadata
}

// DescribeNotifierResponse finds notifier in DescribeNotifierResponse format
func (wr *WriterReader) DescribeNotifierResponse(name string) (uc.DescribeNotifierResponse, error) {
	notifierRow, err := wr.findBaseNotifier(name)
	if err != nil {
		return uc.DescribeNotifierResponse{}, err
	}
	slackRow, err := wr.findSlackMetadata(name)
	if err != nil {
		return uc.DescribeNotifierResponse{}, err
	}
	return uc.DescribeNotifierResponse{
		Name:     notifierRow.name,
		Type:     notifierTypeToString(entities.NotifierType(notifierRow.ntype)),
		Metadata: slackRowToMetadata(slackRow),
	}, nil
}
