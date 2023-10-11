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

func GetCategoryByID(c *gin.Context) (int, error) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}
