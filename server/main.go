package main

import (
	"fmt"
	"net/http"
	"strconv"

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
		Description string `json:"description"`
		Total       uint   `json:"total"`
	}
	transformedBill struct {
		ID          uint    `json:"id"`
		Description string  `json:"description"`
		Total       float32 `json:"total"`
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
	var bills []billModel
	var _bills []transformedBill

	db.Find(&bills)
	if len(bills) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No bills found!"})
		return
	}

	for _, item := range bills {
		floatTotal := float32(item.Total)
		floatTotal /= 100
		_bills = append(_bills, transformedBill{ID: item.ID, Total: floatTotal, Description: item.Description})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _bills})
}

func getSingleBill(c *gin.Context) {

}

func postBill(c *gin.Context) {
	floatTotal, err := strconv.ParseFloat(c.PostForm("total"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": fmt.Errorf("Something went wrong: %s", err)})
		return
	}
	floatTotal *= 100
	total := uint(floatTotal)

	bill := billModel{Description: c.PostForm("description"), Total: total}
	db.Save(&bill)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Bill added!", "billId": bill.ID})
}

func deleteBill(c *gin.Context) {

}

func updateBill(c *gin.Context) {

}
