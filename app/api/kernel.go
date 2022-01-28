package api

import (
	"fmt"
	"github.com/arryuu/golib/golog"
	"github.com/arryuu/golib/goutil"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
	"regexp"
)

var (
	appDebug  bool
	appSuffix string
	log       *golog.LoggerSt
)

func regGroupRouter(g *echo.Group, f interface{}, middleware ...echo.MiddlewareFunc) {
	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}
	goutil.EachStructMethod(f, func(value reflect.Value, i int) {
		for _, method := range methods {
			regStr := fmt.Sprintf("(%s)(.*)", method)
			if r, err := regexp.Compile(regStr); err == nil {
				if r.Match([]byte(value.Type().Method(i).Name)) {
					funcName := value.Type().Method(i).Name
					apiPath := "/" + funcName[len(method):]
					if apiPath != "/" && appSuffix != "" {
						apiPath += "." + appSuffix
					}
					if m := g.Add(funcName[:len(method)], apiPath, value.Method(i).Interface().(func(echo.Context) error), middleware...); appDebug == true {
						log.Debug(fmt.Sprintf("%s - /%s", m.Method, m.Path))
					}
				}
			}
		}
	})
}

func BindParams(ad bool, l *golog.LoggerSt, as string) {
	appDebug = ad
	log = l
	appSuffix = as
}
