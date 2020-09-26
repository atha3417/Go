package main;

import (
	"fmt"
	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
);

func main() {
	connectToDB();

	// router := gin.Default();
	// v1 := router.Group("/api/v1") 
	// {
	// 	articles := v1.Group("/articles")
	// 	{
	// 		articles.GET("/", getHome)
	// 		articles.GET("/:title", getArticle)
	// 		articles.POST("/", postArticle)
	// 	}
	// }
	// router.Run();
}

func connectToDB() {
	db, error := gorm.Open(mysql.New(mysql.Config{
	DSN: "root:root@tcp(127.0.0.1:3306)/learn-gin?charset=utf8&parseTime=True&loc=Local",
	DefaultStringSize: 256,
	DisableDatetimePrecision: true,
	DontSupportRenameIndex: true,
	DontSupportRenameColumn: true,
	SkipInitializeWithVersion: false,
	}), &gorm.Config{});
	if error != nil {
		panic("Failed to connect to database");
	}
	fmt.Println(db);
}

// func getHome(c *gin.Context) {
// 	c.JSON(200, gin.H {
// 		"status": "OK",
// 		"message": "Ini Halaman Home...",
// 	});
// }

// func getArticle(c *gin.Context) {
// 	title := c.Param("title");
// 	c.JSON(200, gin.H {
// 		"status": "OK",
// 		"message": title,
// 	});
// }

// func postArticle(c *gin.Context) {
// 	title := c.PostForm("title");
// 	desc := c.PostForm("desc");

// 	c.JSON(200, gin.H {
// 		"status": "OK",
// 		"title": title,
// 		"desc": desc,
// 	});
// }
