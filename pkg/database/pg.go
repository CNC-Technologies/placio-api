package database

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"placio-pkg/logger"
	"placio-pkg/models"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func (d *Database) Migrate(models ...interface{}) error {
	return d.AutoMigrate(models...)
}

// Connect to database
func Connect(dsn string) (*Database, error) {
	ctx := context.Background()
	logger.Info(ctx, "Connecting to database")
	db, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	logger.Info(ctx, "======================================")
	logger.Info(ctx, "== Connected to database successfully ==")
	logger.Info(ctx, "======================================")

	return db, nil
}

// LoadModels load models
func (d *Database) LoadModels() error {
	var modelList []interface{}
	modelList = append(modelList, &models.User{}, &models.Event{}, &models.Ticket{}, &models.Booking{}, &models.Payment{}, models.Business{}, models.Conversation{}, models.Group{}, models.Message{})
	return d.Migrate(modelList...)
}

//
//type Database struct {
//	*gorm.DB
//	pool *pgxpool.Pool
//}
//
//func NewDatabase(dsn string, poolSize int) (*Database, error) {
//	pool, err := pgxpool.Connect(context.Background(), dsn)
//	if err != nil {
//		return nil, err
//	}
//
//	db, err := gorm.Open(postgres.New(postgres.Config{
//		Conn: pool.Acquire,
//	}), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//
//	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
//	db.SetMaxIdleConns(poolSize)
//
//	// SetMaxOpenConns sets the maximum number of open connections to the database.
//	db.SetMaxOpenConns(poolSize)
//
//	return &Database{db, pool}, nil
//}
//
//func (d *Database) Migrate(models ...interface{}) error {
//	return d.AutoMigrate(models...)
//}
//
//// Connect to database
//func Connect(dsn string, poolSize int) (*Database, error) {
//	ctx := context.Background()
//	logger.Info(ctx, "Connecting to database")
//	db, err := NewDatabase(dsn, poolSize)
//	if err != nil {
//		return nil, err
//	}
//
//	logger.Info(ctx, "======================================")
//	logger.Info(ctx, "== Connected to database successfully ==")
//	logger.Info(ctx, "======================================")
//
//	return db, nil
//}
//
//// LoadModels load models
//func (d *Database) LoadModels() error {
//	var modelList []interface{}
//	modelList = append(modelList, &models.User{}, &models.Event{}, &models.Ticket{}, &models.Booking{}, &models.Payment{}, models.Business{}, models.Conversation{}, models.Group{}, models.Message{})
//	return d.Migrate(modelList...)
//}
//
//func (d *Database) Close() error {
//	d.pool.Close()
//	return nil
//}
