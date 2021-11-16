package infrastructure

import (
	"database/sql"

	"github.com/ITK13201/holodule-bot/interfaces/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/ITK13201/holodule-bot/config"
)

type SqlHandler struct {
	Db *sqlx.DB
}

func NewSqlHandler() database.SqlHandler {
	cfg := *config.Cfg

	db, err := sqlx.Connect("mysql", cfg.DatabaseUrl)
	if err != nil {
		panic(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Db = db

	return sqlHandler
}

func (handler *SqlHandler) NamedExec(query string, arg interface{}) (sql.Result, error) {
	res := SqlResult{}
	result, err := handler.Db.NamedExec(query, arg)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Exec(query string, args ...interface{}) (sql.Result, error) {
	res := SqlResult{}
	result, err := handler.Db.Exec(query, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Get(dest interface{}, query string, args ...interface{}) error {
	err := handler.Db.Get(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (handler *SqlHandler) Select(dest interface{}, query string, args ...interface{}) error {
	err := handler.Db.Select(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}
