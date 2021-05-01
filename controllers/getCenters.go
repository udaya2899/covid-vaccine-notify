package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/udaya2899/covid-vaccine-notify/services"
)

func GetCenters(c *gin.Context) {
	districtID := c.Query("district_id")
	date := c.Query("date")

	fmt.Printf("district ID: %v \n date: %v", districtID, date)
	centers, err := services.GetCenters(districtID, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, centers)
}
