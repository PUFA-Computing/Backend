package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProductID(c *gin.Context) (int, error) {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return 0, err
	}

	return productID, nil
}

func GetCategoryID(c *gin.Context) (int, error) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}

func GetSizeID(c *gin.Context) (int, error) {
	sizeID, err := strconv.Atoi(c.Param("sizeID"))
	if err != nil {
		return 0, err
	}

	return sizeID, nil
}

func GetColorID(c *gin.Context) (int, error) {
	colorID, err := strconv.Atoi(c.Param("colorID"))
	if err != nil {
		return 0, err
	}

	return colorID, nil
}

func GetTransactionID(c *gin.Context) (int, error) {
	transactionID, err := strconv.Atoi(c.Param("transactionID"))
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}
