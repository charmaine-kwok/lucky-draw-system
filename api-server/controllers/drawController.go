package controllers

import (
	"errors"
	"fmt"
	"lucky-draw/initializers"
	"lucky-draw/models"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type prizeResponse struct {
	Prize string `json:"prize" example:"Congratulations! You have won Buy 1 Get 1 Free Coupon!"`
}

type errorResponse struct {
	Error string `json:"error" example:"Some error message..."`
}

var internalServerError = errorResponse{Error: "Internal server error"}

// Allow customer to enter draw
//
//	@Summary		Enter draw
//	@Tags			Customer
//	@Description	Enter draw
//	@Produce		json
//	@Param			customerId	path		string			true	"Customer ID"
//	@Success		200			{object}	prizeResponse	"Prize"
//	@Failure		400			{object}	errorResponse	"Invalid customer id"
//	@Failure		409			{object}	errorResponse	"Conflict"
//	@Failure		500			{object}	errorResponse	"Internal server error"
//	@Router			/draw/{customerId} [get]
func Draw(c *gin.Context) {

	// Get customer Id from the customerId parameter
	customerId, err := strconv.Atoi(c.Param("customerId"))
	if err != nil {
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

	// Check if user have participated today
	if customer.Drawed {
		// Remind user to join tomorrow
		c.JSON(http.StatusConflict, errorResponse{Error: "You have entered the draw today, please join again tomorrow"})
		return
	}

	// Update drawed as true
	initializers.DB.Model(&customer).Where("id = ?", customerId).Update("drawed", true)

	// Draw prize

	var prizes []models.Prize

	var prize *models.Prize

	// Get all prizes
	result = initializers.DB.Find(&prizes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, internalServerError)
		return
	}

	// Generate a random number between 0 and 1
	randomNumber := rand.Float64()
	fmt.Println("randomNumer", randomNumber)

	prob := float64(0)

	// Check the random number against the probabilities
	for _, v := range prizes {
		prob += math.Round(v.Probability*1000) / 1000
		fmt.Println(prob)
		fmt.Println("ran", randomNumber)

		if randomNumber < prob {
			prize = &v
			break
		}
	}
	fmt.Println("prize", *prize)

	dailyQuota := prize.Dailyquota
	fmt.Println("dailyQuota", dailyQuota)

	totalQuota := prize.Totalquota
	fmt.Println("totalQuota", totalQuota)

	// check if the prize has unlimited quota or not
	// if not unlimited
	if dailyQuota != 9999 || totalQuota != 9999 {

		// Check if the prize have reached its daily/total quota or not
		if totalQuota == 0 {
			c.JSON(http.StatusConflict, errorResponse{Error: "Sorry the prize has exceed its total quota. Please enter the draw again tomorrow."})
			return
		}
		if dailyQuota == 0 {
			c.JSON(http.StatusConflict, errorResponse{Error: "Sorry the prize has exceed its daily quota. Please enter the draw again tomorrow."})
			return
		}

		// Reduce the prize's daily quota by 1
		result = initializers.DB.Model(&prize).Where("category = ?", prize.Category).Update("dailyquota", dailyQuota-1)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, internalServerError)
			return
		}

		// Reduce the prize's total quota by 1
		result = initializers.DB.Model(&prize).Where("category = ?", prize.Category).Update("totalquota", totalQuota-1)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, internalServerError)
			return
		}
	}

	var response string

	if prize.Category == "No prize" {
		response = "Sorry you didn't win any prize. Please enter the draw tomorrow"
	} else {
		response = "Congratulations! You have won " + prize.Category + "!"
	}

	// Respond with the prize
	c.JSON(http.StatusOK, prizeResponse{
		Prize: response,
	})
}
