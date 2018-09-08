package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "kadikis"
	dbname = "bill_tracker"
)

func init() {
	var psqlInfo string
	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
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
