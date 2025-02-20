package relation

import (
	"database/sql"
	"go-fiber-template/pkg/logging"
	"github.com/jmoiron/sqlx"
)

var log = logging.RelationLogger

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) withTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Error().Err(err).Msg("Ошибка создания транзакции")
		return err
	}
	defer func() {
		if err != nil {
			log.Error().Err(tx.Rollback()).Msg("Ошибка отката транзакции")
		}
	}()
	if err = fn(tx); err != nil {
		log.Error().Err(err).Msg("Ошибка выполнения транзакции")
		return err
	}
	if err = tx.Commit(); err != nil {
		log.Error().Err(err).Msg("Ошибка сохранения транзакции")
		return err
	}
	return nil
}