package config

import (
	"os"

	"src/engine/encoding"
	"src/engine/encoding/yaml"
)

var Conf = new(Config)

func Setup(filename string) error {
	// 读取文件
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// 反序列化配置到结构体
	if err = encoding.GetCodec(yaml.Name).Unmarshal(file, Conf); err != nil {
		return err
	}

	return nil
}

type Config struct {
	APP    APP    `yaml:"app"`
	Server Server `yaml:"server"`
	Data   Data   `yaml:"data"`
}

type APP struct {
	Env string `yaml:"env"`
}

type Server struct {
	Http HTTP `yaml:"http"`
}

type HTTP struct {
	Port int `yaml:"port"`
}

type Data struct {
	Cache    map[string]Cache    `yaml:"cache"`
	Database map[string]Database `yaml:"database"`
}

type Database struct {
	DSN         string `yaml:"dsn"`
	MaxIdleConn int    `yaml:"max_idle_conn"`
	MaxOpenConn int    `yaml:"max_open_conn"`
}

type Cache struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
