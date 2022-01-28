package module

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var (
	confName = "env"
	confType = "yaml"

	env = NewEnv()

	AppDebug  = env.GetBool("app.debug")
	AppSuffix = env.GetString("app.suffix")
)

type EnvMysqlSt struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Charset  string
	Debug    bool
	AppDebug bool
}

func NewEnv() *viper.Viper {
	path, _ := os.Getwd()
	env := viper.New()
	env.AddConfigPath(path)
	env.SetConfigName(confName)
	env.SetConfigType(confType)
	env.WatchConfig()
	if err := env.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			env.Set("app.port", 1323)
			env.Set("app.suffix", "go")
			env.Set("app.debug", false)
			env.Set("database.mysql.debug", false)
			env.Set("database.mysql.1.host", "127.0.0.1")
			env.Set("database.mysql.1.port", 3306)
			env.Set("database.mysql.1.database", "database")
			env.Set("database.mysql.1.username", "root")
			env.Set("database.mysql.1.password", "root")
			env.Set("database.mysql.1.charset", "utf8mb4")
			_ = env.SafeWriteConfig()
		}
	}
	return env
}

func GetMysql(i int) *EnvMysqlSt {
	return &EnvMysqlSt{
		Host:     env.GetString(fmt.Sprintf("database.mysql.%d.host", i)),
		Port:     env.GetString(fmt.Sprintf("database.mysql.%d.port", i)),
		Database: env.GetString(fmt.Sprintf("database.mysql.%d.database", i)),
		Username: env.GetString(fmt.Sprintf("database.mysql.%d.username", i)),
		Password: env.GetString(fmt.Sprintf("database.mysql.%d.password", i)),
		Charset:  env.GetString(fmt.Sprintf("database.mysql.%d.charset", i)),
		Debug:    env.GetBool("database.mysql.debug"),
		AppDebug: AppDebug,
	}
}
