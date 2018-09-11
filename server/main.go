package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
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
	//TODO: use the binding for a function to seperate the from validator from the function
	billBinding struct {
		Description string `form:"description" json:"description" binding:"required"`
		Total       string `form:"total" json:"total" binding:"required"`
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

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&billModel{})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

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
		floatTotal := makeFloatTotal(item.Total)
		_bills = append(_bills, transformedBill{ID: item.ID, Total: floatTotal, Description: item.Description})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _bills})
}

func getSingleBill(c *gin.Context) {
	bill, err := getSingleBillModel(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No bill found!"})
		return
	}

	floatTotal := makeFloatTotal(bill.Total)

	_bill := transformedBill{ID: bill.ID, Description: bill.Description, Total: floatTotal}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _bill})
}

func makeFloatTotal(total uint) float32 {
	floatTotal := float32(total)
	return floatTotal / 100
}

func makeIntTotal(floatTotal float32) uint {
	floatTotal *= 100
	return uint(floatTotal)
}

func postBill(c *gin.Context) {
	var billBinding billBinding

	if err := c.ShouldBind(&billBinding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if len(billBinding.Description) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Missing description!"})
		return
	}

	floatTotal, err := strconv.ParseFloat(billBinding.Total, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": fmt.Errorf("Something went wrong: %s", err), "error": err})
		return
	}

	total := makeIntTotal(float32(floatTotal))

	bill := billModel{Description: c.PostForm("description"), Total: total}
	db.Save(&bill)

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Bill added!", "billId": bill.ID})
}

func deleteBill(c *gin.Context) {
	bill, err := getSingleBillModel(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err})
		return
	}
	db.Delete(&bill)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Bill deleted!"})
}

func updateBill(c *gin.Context) {
	var billBinding billBinding

	if err := c.ShouldBind(&billBinding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if len(billBinding.Description) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Missing description!"})
		return
	}

	bill, err := getSingleBillModel(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err})
	}

	floatTotal, err := strconv.ParseFloat(billBinding.Total, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": fmt.Errorf("Something went wrong: %s", err)})
		return
	}
	total := makeIntTotal(float32(floatTotal))

	bill.Description = billBinding.Description
	bill.Total = total
	db.Save(&bill)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Bill updated!"})
}

func getSingleBillModel(billId string) (billModel, error) {
	var bill billModel
	db.First(&bill, billId)

	if bill.ID == 0 {
		return bill, errors.New("No bill found!")
	}

	return bill, nil
}
