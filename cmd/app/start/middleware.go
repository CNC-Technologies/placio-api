package start

import (
	"context"
	"fmt"

	"github.com/axiaoxin-com/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	//"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"

	csrf "github.com/utrack/gin-csrf"
	// "gorm.io/gorm/logger"
	// "gorm.io/gorm/logger"
)

func formatDate(t time.Time) string {
	return t.Format(time.RFC822)
}

func PrintHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Print request headers
		fmt.Println("Request Headers:")
		for k, v := range c.Request.Header {
			fmt.Println(k, ":", v)
		}

		// Process request
		c.Next()

		// Print response headers
		fmt.Println("Response Headers:")
		for k, v := range c.Writer.Header() {
			fmt.Println(k, ":", v)
		}
	}
}

func Middleware(app *gin.Engine) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge:     12 * time.Hour,
		AllowFiles: true,
		// AllowAllOrigins: true,
		// AllowWildcard: true,
	}))

	app.Use(func(c *gin.Context) {
		c.Header("Vary", "Origin")
	})

	app.Use(PrintHeaders())

	store := cookie.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("mysession", store))
	app.Use(csrf.Middleware(csrf.Options{
		Secret: "secret123",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	gin.SetMode(gin.DebugMode)
	// you can custom the config or use logging.GinLogger() by default config
	conf := logging.GinLoggerConfig{
		Formatter: func(c context.Context, m logging.GinLogDetails) string {
			return fmt.Sprintf("%s use %s request %s at %v, handler %s use %f seconds to respond it with %d",
				m.ClientIP, m.Method, m.RequestURI, m.ReqTime, m.HandlerName, m.Latency, m.StatusCode)
		},
		SkipPaths:     []string{},
		EnableDetails: false,
		TraceIDFunc:   func(context.Context) string { return "fake-uuid" },
	}
	app.Use(logging.GinLoggerWithConfig(conf))

	//r.Use(ratelimiter.GinMemRatelimiter(ratelimiter.GinRatelimiterConfig{
	//	// config: how to generate a limit key
	//	LimitKey: func(c *gin.Context) string {
	//		return c.ClientIP()
	//	},
	// "github.com/axiaoxin-com/ratelimiter"
	//	// config: how to respond when limiting
	//	LimitedHandler: func(c *gin.Context) {
	//		c.JSON(200, "too many requests!!!")
	//		c.Abort()
	//		return
	//	},
	//	// config: return ratelimiter token fill interval and bucket size
	//	TokenBucketConfig: func(*gin.Context) (time.Duration, int) {
	//		return time.Second * 1, 1
	//	},
	//}))

	//debug := true
	//if debug {
	//	app.Use(inspector.InspectorStats())
	//	app.GET("/_inspector", func(c *gin.Context) {
	//		c.JSON(200, inspector.GetPaginator())
	//	})
	//}

	app.Delims("{{", "}}")

	//app.SetFuncMap(template.FuncMap{
	//	"formatDate": formatDate,
	//})
	//
	//app.LoadHTMLFiles("inspector.html")
	//
	//app.Use(inspector.InspectorStats())
	//app.GET("/_inspector", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "inspector.html", map[string]interface{}{
	//		"title":      "Gin Inspector",
	//		"pagination": inspector.GetPaginator(),
	//	})
	//
	//})

	//app.Use(limit.MaxAllowed(20))
	//app.Use(helmet.Default())

	//p := ginprometheus.NewPrometheus("gin")
	//
	//p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
	//	url := c.Request.URL.Path
	//	for _, p := range c.Params {
	//		if p.Key == "name" {
	//			url = strings.Replace(url, p.Value, ":name", 1)
	//			break
	//		}
	//	}
	//	return url
	//}

	//app.Use(ginprom.PromMiddleware(nil))
	//
	//// register the `/metrics` route.
	//app.GET("/metrics-new", ginprom.PromHandler(promhttp.Handler()))

	//p.Use(app)

}
