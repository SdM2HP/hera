package database

import (
	"strings"
	"time"

	"gorm.io/gorm/logger"
)

type options struct {
	dsn             string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
	logLevel        logger.LogLevel
}

type Option interface {
	apply(*options)
}

type dsnOption string

func (o dsnOption) apply(opts *options) {
	opts.dsn = string(o)
}

func WithDSN(dsn string) Option {
	return dsnOption(dsn)
}

type maxIdleConnOption int

func (o maxIdleConnOption) apply(opts *options) {
	opts.maxIdleConn = int(o)
}

func WithMaxIdleConn(maxIdleConn int) Option {
	return maxIdleConnOption(maxIdleConn)
}

type maxOpenConnOption int

func (o maxOpenConnOption) apply(opts *options) {
	opts.maxOpenConn = int(o)
}

func WithMaxOpenConn(maxOpenConn int) Option {
	return maxOpenConnOption(maxOpenConn)
}

type connMaxLifetimeOption time.Duration

func (o connMaxLifetimeOption) apply(opts *options) {
	opts.connMaxLifetime = time.Duration(o)
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) Option {
	return connMaxLifetimeOption(connMaxLifetime)
}

type connMaxIdleTimeOption time.Duration

func (o connMaxIdleTimeOption) apply(opts *options) {
	opts.connMaxIdleTime = time.Duration(o)
}

func WithConnMaxIdleTime(connMaxIdleTime time.Duration) Option {
	return connMaxIdleTimeOption(connMaxIdleTime)
}

type logLevelOption logger.LogLevel

func (o logLevelOption) apply(opts *options) {
	opts.logLevel = logger.LogLevel(o)
}

func WithLogLevel(level string) Option {
	var val int
	switch strings.ToLower(level) {
	case "silent":
		val = 1
	case "error":
		val = 2
	case "warn":
		val = 3
	default:
		val = 4
	}
	return logLevelOption(val)
}
