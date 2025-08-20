package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"os"
	"strconv"
	"strings"
)

type DB struct {
	db *bun.DB
}
type Filter map[string]any

type Database interface {
	Insert(ctx context.Context, model any) (sql.Result, error)
	Delete(ctx context.Context, tableName string, filter Filter) (sql.Result, error)
	Update(ctx context.Context, tableName string, Set Filter, Condition Filter) (sql.Result, error)
	SelectOne(ctx context.Context, model any, columnName string, params any) error
	SelectAll(ctx context.Context, tableName string, model any) error
	Raw(ctx context.Context, model any, query string, args ...interface{}) error
	Close() error
}

func (d *DB) whereCondition(filter Filter, ConditionType string) string {
	var whereClauses []string
	for key, value := range filter {
		var formateValue string
		switch value.(type) {
		case string:
			formateValue = fmt.Sprintf("%v", value)
		case int:
			formateValue = fmt.Sprintf("%v", value)
		case int64:
			formateValue = fmt.Sprintf("%v", value)
		case float64:
			formateValue = fmt.Sprintf("%.2f", value)
		default:
			log.Fatal("whereCondition: unknown filter type")
		}
		whereClauses = append(whereClauses, fmt.Sprintf("%s = %s", key, formateValue))
	}

	var result string
	if len(whereClauses) > 0 {
		if ConditionType == "SET" {
			result = strings.Join(whereClauses, " , ")
		} else if ConditionType == "AND" {
			result = strings.Join(whereClauses, " AND ")
		}
	}
	return result
}

func (d *DB) Insert(ctx context.Context, model any) (sql.Result, error) {
	result, err := d.db.NewInsert().
		Model(model).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DB) Delete(ctx context.Context, tableName string, filter Filter) (sql.Result, error) {
	result, err := d.db.NewDelete().
		Table(tableName).
		Where(d.whereCondition(filter, "AND")).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DB) Update(ctx context.Context, tableName string, Set Filter, Condition Filter) (sql.Result, error) {
	result, err := d.db.NewUpdate().
		Table(tableName).
		Where(d.whereCondition(Set, "SET")).
		Where(d.whereCondition(Condition, "AND")).
		Exec(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *DB) SelectOne(ctx context.Context, model any, columnName string, params any) error {
	err := d.db.NewSelect().
		Model(model).
		Where(fmt.Sprintf("%s = ?", columnName), params).
		Scan(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) SelectAll(ctx context.Context, tableName string, model any) error {
	err := d.db.NewSelect().Table(tableName).Scan(ctx, model)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Raw(ctx context.Context, model any, query string, args ...interface{}) error {
	err := d.db.NewRaw(query, args...).
		Scan(ctx, model)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func New() DB {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	_port, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPass, dbHost, _port, dbName)
	_database := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	_db := bun.NewDB(_database, pgdialect.New())
	return DB{db: _db}

}
