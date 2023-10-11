package merch

import (
	"Backend/internal/models"
	"Backend/internal/services"
	"Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	MerchService      *services.MerchService
	PermissionService *services.PermissionService
}

func NewMerchHandler(merchService *services.MerchService, permissionService *services.PermissionService) *Handler {
	return &Handler{
		MerchService:      merchService,
		PermissionService: permissionService,
	}
}

/**
 * start
 * @api {products}
 */

func (h *Handler) CreateProduct(c *gin.Context) {
	var newProduct models.MerchProduct
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.MerchService.CreateProduct(&newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, newProduct)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	productID, err := utils.GetProductID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedProduct models.MerchProduct
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.MerchService.UpdateProduct(productID, &updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	productID, err := utils.GetProductID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.MerchService.DeleteProduct(productID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (h *Handler) ListProducts(c *gin.Context) {
	products, err := h.MerchService.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProductByID(c *gin.Context) {
	productID, err := utils.GetProductID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.MerchService.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) GetSizeProduct(c *gin.Context) {
	productID, err := utils.GetProductID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.MerchService.GetSizeProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) GetColorProduct(c *gin.Context) {
	productID, err := utils.GetProductID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.MerchService.GetColorProduct(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) GetProductPrice(c *gin.Context) {
	productID, err := utils.GetProductID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.MerchService.GetProductPrice(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

/**
 * end
 * @api {products}
 */

/**
 * start
 * @api {categories}
 */

func (h *Handler) CreateCategory(c *gin.Context) {
	var newCategory models.MerchCategory
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.MerchService.CreateCategory(&newCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, newCategory)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	categoryID, err := utils.GetCategoryByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedCategory models.MerchCategory
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.MerchService.UpdateCategory(categoryID, &updatedCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	categoryID, err := utils.GetCategoryByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.MerchService.DeleteCategory(categoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (h *Handler) ListCategories(c *gin.Context) {
	categories, err := h.MerchService.ListCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) GetCategoryByID(c *gin.Context) {
	categoryID, err := utils.GetCategoryByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.MerchService.GetCategoryByID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

/**
 * end
 * @api {categories}
 */
