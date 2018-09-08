package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB

const (
	host   = "localhost"
	port   = 5432
	user   = "kadikis"
	dbname = "bill_tracker"
)

type (
	billModel struct {
		gorm.Model
		Description string `json:description`
		Total       uint   `json:total`
	}
)

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&billModel{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/bills")
	{
		v1.POST("/", postBill)
		v1.GET("/", getBills)
		v1.GET("/:id", getSingleBill)
		v1.PUT("/:id", updateBill)
		v1.DELETE("/:id", deleteBill)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func getBills(c *gin.Context) {

}

func getSingleBill(c *gin.Context) {

}

func postBill(c *gin.Context) {

}

func deleteBill(c *gin.Context) {

}

func updateBill(c *gin.Context) {

}
