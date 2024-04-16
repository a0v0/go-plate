package database

import (
	"context"
	"database/sql"
	"fmt"

	"go_plate/internal/config"
	"go_plate/internal/ent"
	"go_plate/internal/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	Ent *ent.Client

	Cfg *config.Config
}

type Seeder interface {
	Seed(*ent.Client) error
	Count() (int, error)
}

func NewDatabase(cfg *config.Config) *Database {
	db := &Database{
		Cfg: cfg,
	}

	return db
}

func (db *Database) ConnectDatabase() error {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		db.Cfg.Database.Postgres.User,
		db.Cfg.Database.Postgres.Password,
		db.Cfg.Database.Postgres.Host,
		db.Cfg.Database.Postgres.Port,
		db.Cfg.Database.Postgres.DBName,
		db.Cfg.Database.Postgres.SSLMode)
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	drv := entsql.OpenDB(dialect.Postgres, conn)
	db.Ent = ent.NewClient(ent.Driver(drv))
	return nil
}

func (db *Database) ShutdownDatabase() error {
	return db.Ent.Close()
}

func (db *Database) MigrateModels() error {
	return db.Ent.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
}

func (db *Database) SeedModels(seeder ...Seeder) error {
	for _, v := range seeder {

		count, err := v.Count()
		if err != nil {
			return err
		}

		if count == 0 {
			err = v.Seed(db.Ent)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

// Rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
