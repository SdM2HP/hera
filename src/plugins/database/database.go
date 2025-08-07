package database

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(opts ...Option) *gorm.DB {
	o := &options{
		dsn:             "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		maxOpenConn:     100,
		maxIdleConn:     1,
		connMaxIdleTime: 1 * time.Hour,
		connMaxLifetime: 1 * time.Hour,
		logLevel:        logger.Info,
	}

	for _, opt := range opts {
		opt.apply(o)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: o.dsn, // data source name
	}), &gorm.Config{
		Logger: logger.Default.LogMode(o.logLevel),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(o.maxIdleConn)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(o.maxOpenConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(o.connMaxLifetime)

	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	sqlDB.SetConnMaxIdleTime(o.connMaxIdleTime)

	return db
}

func DBName(db *gorm.DB) (string, error) {
	var dbName string
	if err := db.Raw("SELECT DATABASE()").Scan(&dbName).Error; err != nil {
		return "", err
	}
	return dbName, nil
}
