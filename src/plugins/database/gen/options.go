package gen

import "gorm.io/gen"

type (
	Option        func(opt *options)
	ModelNameFunc func(tableName string) (modelName string)
)

type options struct {
	outPath       string
	modelNameFunc ModelNameFunc
	modelOpts     []gen.ModelOpt
}

func WithOutPath(outPath string) Option {
	return func(opt *options) {
		opt.outPath = outPath
	}
}

func WithModelNameFunc(modelNameFunc ModelNameFunc) Option {
	return func(opt *options) {
		opt.modelNameFunc = modelNameFunc
	}
}

func WithModelOpt(modelOpts []gen.ModelOpt) Option {
	return func(opt *options) {
		opt.modelOpts = modelOpts
	}
}
