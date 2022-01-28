package module

import (
	"app/app/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type (
	MysqlSt struct{}
)

func (*MysqlSt) Mysql_1(envMysql *EnvMysqlSt) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		envMysql.Username,
		envMysql.Password,
		envMysql.Host,
		envMysql.Port,
		envMysql.Database,
		envMysql.Charset)

	logLevel := logger.Silent
	if envMysql.Debug {
		logLevel = logger.Info
	}
	nerLogger := logger.New(
		Log,
		logger.Config{
			SlowThreshold: 0 * time.Microsecond,
			LogLevel:      logLevel,
			Colorful:      false,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 跳过默认开启事务模式
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: nerLogger,
	})
	if err != nil {
		Log.Error(fmt.Errorf("failed to initialize database, got error %w\n", err))
		return
	}
	models.BindDb(db)
}
