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
	DBUser       string `required:"true"`
	DBPassword   string `required:"true"`
	DBAddress    string `required:"true"`
	DBPort       string `required:"true"`
	DBName       string `required:"true"`
	ServerPort   string `required:"true"`
	KeyFilePath  string `required:"true"`
	MigrationDir string `required:"true"`
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
