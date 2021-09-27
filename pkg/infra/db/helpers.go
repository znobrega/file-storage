package db

import (
	"errors"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Helper struct {
	dbConnection *gorm.DB
}

func (h *Helper) InitDatabase() error {
	log.Println("[DATABASE] Connecting to database")
	dsn := viper.GetString("database.connectionUri")

	var err = errors.New("connecting db")
	var gormDb *gorm.DB
	for err != nil {
		gormDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			CreateBatchSize: viper.GetInt("database.insertBatchSize"),
		})
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
		}
	}

	if viper.GetBool("database.showSql") {
		gormDb = gormDb.Debug()
	}

	if viper.GetBool("database.migrations.enabled") {
		err = h.runMigrations(gormDb)
		if err != nil {
			return err
		}
	}

	h.dbConnection = gormDb
	log.Println("[DATABASE] Connected")
	return nil
}

func (h *Helper) runMigrations(gormDb *gorm.DB) error {
	migrationsDirectory := migrate.FileMigrationSource{
		Dir: viper.GetString("database.migrations.directory"),
	}
	sqlDb, err := gormDb.DB()
	if err != nil {
		return err
	}
	amountOfMigrations, err := migrate.Exec(sqlDb, viper.GetString("database.dialect"), migrationsDirectory, migrate.Up)
	if err != nil {
		return err
	}
	log.Println("executed migrations: ", amountOfMigrations)
	return nil
}

func (h *Helper) GetDatabase() *gorm.DB {
	return h.dbConnection
}
