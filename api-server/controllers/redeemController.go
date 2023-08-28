package controllers

import (
	"errors"
	"lucky-draw/initializers"
	"lucky-draw/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Customer enter mobile phone number to redeem prize
//
//	@Summary		Redeem prize
//	@Tags			Customer
//	@Description	Redeem prize
//	@Accept			json
//	@Param			customerId	path		string			true	"Customer ID"
//	@Param			body		body		models.Mobile	true	"Mobile"
//	@Success		200			{object}	models.Mobile	"Mobile"
//	@Failure		400			{object}	errorResponse	"Invalid customer id"
//	@Failure		500			{object}	errorResponse	"Internal server error"
//	@Router			/redeem/{customerId} [post]
func Redeem(c *gin.Context) {

	var body *models.Mobile

	// Bind the JSON request body to the body variable
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "Please send us your mobile number"})
		return
	}

	// Get customer Id from the customerId parameter
	customerId, err := strconv.Atoi(c.Param("customerId"))

	if err != nil {
		// If there is an error converting the customerId to an integer, return a bad request error
		c.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid customer id"})
		return
	}

	var customer *models.Customer

	// Get the customer from database
	result := initializers.DB.Where("id = ?", customerId).First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid customer id"})
		} else {
			c.JSON(http.StatusInternalServerError, internalServerError)
		}
		return
	}

	if customer == nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid customer id"})
		return
	}

	// Check if body.Mobile is an 8-digit integer
	mobileInt, err := strconv.Atoi(body.Mobile)
	if err != nil || mobileInt < 10000000 || mobileInt > 99999999 {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "Invalid mobile number. Please enter an 8-digit integer."})
		return
	}

	mobile := &models.Mobile{
		Customerid: customerId,
		Mobile:     body.Mobile,
	}

	// Create a row in mobile table
	result = initializers.DB.Table("mobile").Create(&mobile)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, internalServerError)
		return
	}

	// Return the mobile struct as JSON in the response
	c.JSON(http.StatusOK, mobile)
}
