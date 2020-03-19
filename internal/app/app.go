package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go-admin-x/internal/controllers/common"
	"go-admin-x/internal/middleware"
	models "go-admin-x/internal/model"
	routers "go-admin-x/internal/router"
	"go-admin-x/internal/util/conf"
	"golang.org/x/sync/errgroup"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	g errgroup.Group
)

// 运行
func Run() {
	initDB()
	common.InitCsbinEnforcer()
	initHTTPServer()
}

func initDB() {
	models.InitDB()
	models.Migration()
}

func frontend() http.Handler {
	e := gin.New()
	// 注册自定义函数
	e.SetFuncMap(template.FuncMap{
		"map": func(json string) gin.H {
			var out gin.H
			_ = jsoniter.UnmarshalFromString(json, &out)
			return out
		},
	})

	//e.Static("/static","web/admin/static")
	// 加载模板 注意模板可以定义命名为模块嵌套使用
	e.LoadHTMLGlob("web/admin/tmpl/**/*")
	// 统一访问
	e.GET("/*filepath", func(context *gin.Context) {
		if strings.HasPrefix(context.Request.URL.Path, "/static") || context.Request.URL.Path == "/favicon.ico" {
			context.File("web/admin/" + context.Request.URL.Path)
			return
		}
		context.HTML(http.StatusOK, context.Request.URL.Path, nil)
	})
	return e
}

func backend() http.Handler {
	e := gin.New()
	e.Use(gin.Logger())
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:12001"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "X-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:12002"
		},
		MaxAge: 12 * time.Hour,
	}))
	e.NoRoute(middleware.NoRouteHandler())
	// 崩溃恢复
	e.Use(middleware.RecoveryMiddleware())
	routers.RegisterRouter(e)
	return e
}

func initHTTPServer() {
	gin.SetMode(gin.DebugMode) //调试模式

	server01 := &http.Server{
		Addr:         ":" + conf.GetString("PORT"),
		Handler:      frontend(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":12002",
		Handler:      backend(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
