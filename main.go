package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bigchange/go-pro/myproject/jwt"
	"github.com/bigchange/go-pro/myproject/utils"
	"github.com/gin-gonic/gin"
	// "github.com/bigchange/go-pro/myproject/db"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跨域配置
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// handler request
func helloGinHandler(c *gin.Context) {
	c.JSON(200, "hello world!!")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Gin Web FrameWork
func GinInit() {
	r := gin.Default()
	configJsonPath := flag.String("config_json_path", "", "the config.json path")
	flag.Parse()
	config := &utils.LLBConfig{DbName: "postgres_test",
		DbHost: "localhost", DbUser: "postgres", DbPass: "postgres@"}
	if len(*configJsonPath) > 0 {
		configData, err := ioutil.ReadFile(*configJsonPath)
		check(err)
		err = json.Unmarshal(configData, &config)
		check(err)
	}
	utils.Init()
	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "casem zone",
		Key:        []byte("this is casemind go api"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userName string, password string, c *gin.Context) (string, bool) {
			/**
			user, err := userModel.SigninEx(userName, password)
			if err != nil {
				return "", false
			}
			return strconv.Itoa(user.ID), true
			*/
			return "", false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Token",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
	r.Use(CORSMiddleware())
	r.POST("/api/login", authMiddleware.LoginHandler)
	apiAuth := r.Group("/api/auth")
	apiAuth.Use(authMiddleware.MiddlewareFunc())
	apiAuth.GET("/refresh_token", authMiddleware.RefreshHandler)

	// db.Init()

	apiUser := r.Group("/api/user")
	{
		apiUser.GET("/index", helloGinHandler)
	}

	r.LoadHTMLGlob("./public/*.html")

	r.Static("/static", "./casem/case_go/public/static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "index.html", gin.H{})
	})

	r.Run(":20989")
}

func main() {
}
