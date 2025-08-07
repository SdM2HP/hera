package cache

type Option func(opt *options)

type options struct {
	addr     string
	password string
	db       int
}

func WithAddr(addr string) Option {
	return func(opt *options) {
		opt.addr = addr
	}
}

func WithPassword(password string) Option {
	return func(opt *options) {
		opt.password = password
	}
}

func WithDB(db int) Option {
	return func(opt *options) {
		opt.db = db
	}
}
