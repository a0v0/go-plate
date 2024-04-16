package database

import (
	"context"
	"fmt"

	"go_plate/internal/ent"
	"go_plate/internal/ent/migrate"
	"go_plate/pkg/config"

	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
)

type Database struct {
	Ent *ent.Client
	Log *zap.Logger
	Cfg *config.Config
}

type Seeder interface {
	Seed(*ent.Client) error
	Count() (int, error)
}

func NewDatabase(cfg *config.Config, log *zap.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

func (db *Database) ConnectDatabase() {
	conn, err := sql.Open("pgx", db.Cfg.DB.Postgres.DSN)
	if err != nil {

		db.Log.Error("An unknown error occurred when to connect the database!", zap.Error(err))
	} else {
		db.Log.Info("Connected the database succesfully!")
	}

	drv := entsql.OpenDB(dialect.Postgres, conn)
	db.Ent = ent.NewClient(ent.Driver(drv))
}

func (db *Database) ShutdownDatabase() {
	if err := db.Ent.Close(); err != nil {
		db.Log.Error("An unknown error occurred when to shutdown the database!", zap.Error(err))
	}
}

func (db *Database) MigrateModels() {
	if err := db.Ent.Schema.Create(context.Background(), schema.WithAtlas(true), migrate.WithGlobalUniqueID(true)); err != nil {
		db.Log.Error("Failed creating schema resources!", zap.Error(err))
	} else {
		db.Log.Info("Models were migrated successfully!")
	}
}

func (db *Database) SeedModels(seeder ...Seeder) {
	for _, v := range seeder {

		count, err := v.Count()
		if err != nil {
			db.Log.Panic("", zap.Error(err))
		}

		if count == 0 {
			err = v.Seed(db.Ent)
			if err != nil {
				db.Log.Panic("", zap.Error(err))
			}
			db.Log.Debug("Table has seeded successfully.")

		} else {
			db.Log.Warn("Table has seeded already. Skipping!")
		}
	}
	db.Log.Info("Seeding was completed!")
}

// Rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
