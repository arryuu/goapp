package module

import (
	"app/app/api"
	"fmt"
	"github.com/arryuu/golib/golog"
	"github.com/arryuu/golib/goutil"
	"github.com/labstack/echo/v4"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var (
	SyncOnce sync.Once
)

func Run() {
	SyncOnce.Do(func() {
		Log = golog.New(AppDebug)
		api.BindParams(AppDebug, Log, AppSuffix)

		goutil.EachStructMethod(&MysqlSt{}, func(valueOf reflect.Value, i int) {
			funcName := valueOf.Type().Method(i).Name
			pos := strings.Index(valueOf.Type().Method(i).Name, "_")
			if strings.ToLower(funcName[:pos]) == "mysql" {
				inti, _ := strconv.Atoi(funcName[pos+1:])
				valueOf.Method(i).Interface().(func(st *EnvMysqlSt))(GetMysql(inti))
			}
		})

		serverStart()
	})
}

func serverStart() {
	// 路由
	e := echo.New()
	g := e.Group("")
	api.BindApi(g)

	// windows
	//e.Logger.Fatal(e.Start(":1323"))
	e.Server.Addr = fmt.Sprintf(":%s", env.GetString("app.port"))
	err := e.Server.ListenAndServe()
	if err != nil {
		e.Logger.Fatal(err)
		return
	}

	// linux
	/*err := gracehttp.Serve(e.Server)
	if err != nil {
		return
	}*/
}
