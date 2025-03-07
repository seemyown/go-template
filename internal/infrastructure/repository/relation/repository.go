package relation

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"go-fiber-template/pkg/logging"
)

var log = logging.New(logging.Config{
	FileName: "relation",
	Name:     "relation",
})

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) withTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		log.Error(err, "Ошибка создания транзакции")
		return err
	}
	defer func() {
		if err != nil {
			log.Error(tx.Rollback(), "Ошибка отката транзакции")
		}
	}()
	if err = fn(tx); err != nil {
		log.Error(err, "Ошибка выполнения транзакции")
		return err
	}
	if err = tx.Commit(); err != nil {
		log.Error(err, "Ошибка сохранения транзакции")
		return err
	}
	return nil
}
