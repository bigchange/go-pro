package main

import (
	"flag"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/bigchange/go-pro/myproject/utils"
	"github.com/bigchange/go-pro/myproject/db"
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
func helloHandler(c *gin.Context) {
	c.JSON(200, "hello world!!")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	r := gin.Default()
	var err error 
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
	r.Use(CORSMiddleware())
	db.Init()

	apiUser := r.Group("/api/user")
	{
		apiUser.GET("/index",helloHandler)
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
