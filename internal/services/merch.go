package services

import (
	"Backend/internal/database/app"
	"Backend/internal/models"
)

type MerchService struct {
}

func NewMerchService() *MerchService {
	return &MerchService{}
}

/**
 * start
 * @api {products}
 */

func (ms *MerchService) CreateProduct(merch *models.MerchProduct) error {
	if err := app.CreateProduct(merch); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) UpdateProduct(productID int, updatedProduct *models.MerchProduct) error {
	existingProduct, err := app.GetProductByID(productID)
	if err != nil {
		return err
	}

	updatedProduct.CreatedAt = existingProduct.CreatedAt

	if err := app.UpdateProduct(productID, updatedProduct); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) DeleteProduct(productID int) error {
	if err := app.DeleteProduct(productID); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) ListProducts() ([]*models.MerchProduct, error) {
	products, err := app.ListProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ms *MerchService) GetProductByID(productID int) (*models.MerchProduct, error) {
	product, err := app.GetProductByID(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ms *MerchService) GetSizeProduct(productID int) ([]*models.MerchSize, error) {
	sizes, err := app.GetSizeProduct(productID)
	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (ms *MerchService) GetColorProduct(productID int) ([]*models.MerchColor, error) {
	colors, err := app.GetColorProduct(productID)
	if err != nil {
		return nil, err
	}

	return colors, nil
}

func (ms *MerchService) GetProductPrice(productID int) (*models.MerchPrice, error) {
	price, err := app.GetProductPrice(productID)
	if err != nil {
		return nil, err
	}

	return price, nil
}

/**
 * end services
 * @api {products}
 */

/**
 * start services
 * @api {categories}
 */

func (ms *MerchService) CreateCategory(merch *models.MerchCategory) error {
	if err := app.CreateCategory(merch); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) UpdateCategory(categoryID int, updatedCategory *models.MerchCategory) error {
	existingCategory, err := app.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	updatedCategory.CreatedAt = existingCategory.CreatedAt

	if err := app.UpdateCategory(categoryID, updatedCategory); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) DeleteCategory(categoryID int) error {
	if err := app.DeleteCategory(categoryID); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) ListCategories() ([]*models.MerchCategory, error) {
	categories, err := app.ListCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (ms *MerchService) GetCategoryByID(categoryID int) (*models.MerchCategory, error) {
	category, err := app.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	return category, nil
}

/**
 * end services
 * @api {categories}
 */

/**
 * start services
 * @api {sizes}
 */

func (ms *MerchService) CreateSize(merch *models.MerchSize) error {
	if err := app.CreateSize(merch); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) UpdateSize(sizeID int, updatedSize *models.MerchSize) error {
	existingSize, err := app.GetSizeByID(sizeID)
	if err != nil {
		return err
	}

	updatedSize.CreatedAt = existingSize.CreatedAt

	if err := app.UpdateSize(sizeID, updatedSize); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) DeleteSize(sizeID int) error {
	if err := app.DeleteSize(sizeID); err != nil {
		return err
	}

	return nil
}

/**
 * end services
 * @api {sizes}
 */

/**
 * start services
 * @api {colors}
 */

func (ms *MerchService) CreateColor(merch *models.MerchColor) error {
	if err := app.CreateColor(merch); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) UpdateColor(colorID int, updatedColor *models.MerchColor) error {
	existingColor, err := app.GetColorByID(colorID)
	if err != nil {
		return err
	}

	updatedColor.CreatedAt = existingColor.CreatedAt

	if err := app.UpdateColor(colorID, updatedColor); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) DeleteColor(colorID int) error {
	if err := app.DeleteColor(colorID); err != nil {
		return err
	}

	return nil
}

/**
 * end services
 * @api {colors}
 */

/**
 * start services
 * @api {transactions}
 */

func (ms *MerchService) ListTransactions() ([]*models.MerchTransaction, error) {
	transactions, err := app.ListTransactions()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (ms *MerchService) CreateTransaction(merch *models.MerchTransaction) error {
	if err := app.CreateTransaction(merch); err != nil {
		return err
	}

	return nil
}

func (ms *MerchService) GetTransaction(transactionID int) (*models.MerchTransaction, error) {
	transaction, err := app.GetTransaction(transactionID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

/**
 * end services
 * @api {transactions}
 */
