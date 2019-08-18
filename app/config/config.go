package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"reflect"
	"strings"
)

var _env *environment

const (
	EnvPrefix = "Treasure"
)

func Env() *environment {
	return _env
}

type environment struct {
	DBUser         string `default:"root"`
	DBPassword     string `default:""`
	DBAddress      string `default:"127.0.0.1"`
	DBPort         string `default:"3306"`
	DBName         string `default:"nico"`
	ServerPort     string `default:"8080"`
	TokenSecretKey string `default:"token-secret-key"`
	MigrationDir   string `default:"migration"`
}

func (e environment) GetPrefixedFieldNames() []string {
	ret := make([]string, 0)
	rv := reflect.New(reflect.TypeOf(e)).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		ret = append(ret, strings.ToUpper(EnvPrefix+"_"+f.Name))
	}
	return ret
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".envファイルの読み込みに失敗:", err)
	}

	var env environment
	if err := envconfig.Process(EnvPrefix, &env); err != nil {
		log.Fatal(err)
	}
	_env = &env
}
