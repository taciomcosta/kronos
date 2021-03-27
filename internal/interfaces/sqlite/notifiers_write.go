package sqlite

import (
	"github.com/taciomcosta/kronos/internal/entities"
)

// CreateNotifier creates a new job into database
func (wr *WriterReader) CreateNotifier(notifier *entities.Notifier) error {
	return db.WithTx(func() error {
		return wr.createSlackNotifier(notifier)
	})
}

// DeleteNotifier deletes a notifier
func (wr *WriterReader) DeleteNotifier(name string) error {
	// TODO: implement query
	return nil
}

func (wr *WriterReader) createSlackNotifier(notifier *entities.Notifier) error {
	err := wr.runWriteOperation(
		insertNotifierSQL,
		notifier.Name,
		int(notifier.Type),
	)
	if err != nil {
		return err
	}
	err = wr.runWriteOperation(
		insertSlackSQL,
		notifier.Metadata["auth_token"],
		notifier.Metadata["channel_ids"],
		notifier.Name,
	)
	return err
}
