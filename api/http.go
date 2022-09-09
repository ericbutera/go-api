// TODO:
// https://github.com/axiaoxin-com/ratelimiter
// https://github.com/gin-gonic/contrib

package api

import (
	"log"
	"time"

	config "github.com/ericbutera/go-api/config"
	_ "github.com/ericbutera/go-api/docs"

	// "github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
	Config *config.Config
}

func NewApp(config *config.Config) (*App, error) {
	// TODO: check out google wire DI
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	app := &App{
		Logger: logger,
		Config: config,
	}

	return app, nil
}

// @title Storage API
// @BasePath /
func (app *App) Serve() {

	r := gin.Default()
	// TODO: gin.SetMode(gin.DebugMode)

	// r.Use(requestid.New())
	app.InitLogger(r)
	// app.InitDataDog(r)
	InitOpenTel(app, r)

	app.Routes(r)

	app.Logger.Info("Starting server")
	app.Sugar.Infow("starting",
		"app name", app.Config.AppName,
		"version", app.Config.Version,
	)

	r.Run()
}

func (app *App) InitLogger(r *gin.Engine) {
	// TODO: NewProduction()
	logger, err := zap.NewDevelopment()
	defer logger.Sync()

	if err != nil {
		log.Fatalf("log error %v", err)
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	app.Logger = logger
	app.Sugar = logger.Sugar()
}

func (app *App) Routes(r *gin.Engine) {
	r.GET("/swagger/*any", swagger())
	r.GET("/", app.Docs)
	r.GET("/health", app.Health)
	r.GET("/wait", app.Wait)
	// r.POST("/scan/start", app.Start)
	// r.POST("/scan/:id/save", app.Save)
	// r.POST("/scan/:id/finish", app.Finish)

}
