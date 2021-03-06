package orchestrate

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/apex/log"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"gopkg.in/gin-contrib/cors.v1"

	"github.com/ooni/orchestra/common/middleware"

	"github.com/ooni/orchestra/orchestrate/orchestrate/api/v1"
	"github.com/ooni/orchestra/orchestrate/orchestrate/sched"
)

var ctx = log.WithFields(log.Fields{
	"pkg": "orchestrate",
	"cmd": "ooni-orchestrate",
})

func initDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", viper.GetString("database.url"))
	if err != nil {
		ctx.Error("failed to open database")
		return nil, err
	}
	return db, err
}

func loadTemplates(list ...string) multitemplate.Render {
	r := multitemplate.New()
	for _, x := range list {
		templateString, err := Asset("orchestrate/data/templates/" + x)
		if err != nil {
			ctx.WithError(err).Error("failed to load template")
		}

		tmplMessage, err := template.New(x).Parse(string(templateString))
		if err != nil {
			ctx.WithError(err).Error("failed to parse template")
		}
		r.Add(x, tmplMessage)
	}
	return r
}

// SetupRouter will create a gin.Engine
func SetupRouter(dbURL string) *gin.Engine {
	var (
		err error
	)

	dbMiddleware, err := middleware.InitDatabaseMiddleware("postgres", dbURL)
	if err != nil {
		ctx.WithError(err).Error("failed to init database middleware")
		return nil
	}

	authMiddleware, err := middleware.InitAuthMiddleware(dbMiddleware.DB)
	if err != nil {
		ctx.WithError(err).Error("failed to initialise authMiddlewareDevice")
		return nil
	}
	schedMiddleware, err := sched.InitSchedMiddleware(dbMiddleware.DB)
	if err != nil {
		ctx.WithError(err).Error("failed to initialise schedMiddleware")
		return nil
	}

	router := gin.Default()
	router.Use(schedMiddleware.MiddlewareFunc())
	router.Use(dbMiddleware.MiddlewareFunc())
	router.Use(cors.New(middleware.CorsConfig()))
	router.HTMLRender = loadTemplates("home.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title":                "ooni-orchestrate",
			"componentName":        "ooni-orchestrate",
			"componentDescription": LongDescription,
		})
	})
	err = apiv1.BindAPI(router, authMiddleware)
	if err != nil {
		ctx.WithError(err).Error("failed to BinAPI")
		return nil
	}
	return router
}

// Start starts the events backend including the web handlers
func Start() {

	Addr := fmt.Sprintf("%s:%d", viper.GetString("api.address"),
		viper.GetInt("api.port"))
	ctx.Infof("starting on %s", Addr)

	router := SetupRouter(viper.GetString("database.url"))
	if router == nil {
		panic("failed to start")
	}

	s := &http.Server{
		Addr:    Addr,
		Handler: router,
	}
	gracehttp.Serve(s)
}
