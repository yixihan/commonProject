package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.New()



	// 路由
	r.GET("/test/get", func(context *gin.Context) {
		fmt.Println("走 /test/get 接口")
		context.String(200, "Hello, Go-Gin!")
	})
	r.POST("/test/post", func(context *gin.Context) {
		fmt.Println("走 /test/post 接口")
		context.String(200, "Hello, Go-Gin!")
	})
	r.GET("/", func(context *gin.Context) {
		fmt.Println("走 / 路由")
		context.String(http.StatusOK, "欢迎访问~")
	})
	r.POST("/param/user/:name", func(context *gin.Context) {
		fmt.Println("走 /param/user/:name 路由")
		name := context.Param("name")
		fmt.Printf("user = %v\n", name)
		context.String(http.StatusOK, "欢迎 %v\n", name)
	})
	r.POST("/query/user", func(context *gin.Context) {
		fmt.Println("走 /query/user 路由")
		name := context.Query("name")
		fmt.Printf("user = %v\n", name)
		context.String(http.StatusOK, "欢迎 %v\n", name)
	})
	r.POST("/body/user", func(context *gin.Context) {
		fmt.Println("走 /body/user 路由")
		name := context.PostForm("name")
		password := context.DefaultPostForm("password", "123456789")

		fmt.Printf("name = %v, password = %v\n", name, password)
		context.JSON(http.StatusOK, gin.H{
			"name":     name,
			"password": password,
		})
	})
	r.POST("/map/user", func(context *gin.Context) {
		fmt.Println("走 /map/user 路由")

		nameMap := context.PostFormMap("name")
		idMap := context.QueryMap("id")
		fmt.Printf("nameMap = %v, idMap = %v\n", nameMap, idMap)
		context.JSON(http.StatusOK, gin.H{
			"nameMap": nameMap,
			"idMap":   idMap,
		})
	})
	r.POST("/array/user", func(context *gin.Context) {
		fmt.Println("走 /array/user 路由")

		nameArray := context.PostFormArray("name")
		idArray := context.QueryArray("id")
		fmt.Printf("nameArray = %v, idArray = %v\n", nameArray, idArray)
		context.JSON(http.StatusOK, gin.H{
			"nameArray": nameArray,
			"idArray":   idArray,
		})
	})
	r.POST("/redirect/in", func(context *gin.Context) {
		fmt.Println("走 /redirect/in 路由, 即将被重定向到 /redirect/out 路由")
		context.Redirect(http.StatusMovedPermanently, "/redirect/out")
	})
	r.GET("/redirect/out", func(context *gin.Context) {
		fmt.Println("走 /redirect/out 路由")
		context.String(http.StatusOK, "这是重定向后的页面")
	})

	// 分组路由
	defaultHandler := func(context *gin.Context) {
		path := context.FullPath()
		fmt.Printf("走 %v 路由\n", path)
		context.JSON(http.StatusOK, gin.H{
			"path": path,
		})
	}
	v1 := r.Group("/v1")
	{
		v1.POST("/test1", defaultHandler)
		v1.POST("/test2", defaultHandler)
		v1.POST("/test3", defaultHandler)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/test1", defaultHandler)
		v2.POST("/test2", defaultHandler)
		v2.POST("/test3", defaultHandler)
	}

	// 文件上传
	r.POST("/file/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		// 上传文件至指定目录
		if err := c.SaveUploadedFile(file, "D:/upload/"+file.Filename); err != nil {
			fmt.Println(err)
		}
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})
	r.POST("/file/upload2", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file[]"]

		for _, file := range files {
			// 上传文件至指定目录
			if err := c.SaveUploadedFile(file, "D:/upload/"+file.Filename); err != nil {
				fmt.Println(err)
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

	})

	// 中间件
	// 作用于全局
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 作用于单个路由
	r.GET("/mid/test01",MyBenchLogger())
	// 作用于一个组
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	_ = r.Run(":8080")
}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 给Context实例设置一个值
		c.Set("geektutu", "1111")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		log.Print(latency)
	}
}
