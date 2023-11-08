package app

import (
	"Backend/internal/database"
	"Backend/internal/models"
	"context"
)

type Merch struct {
}

/**
 * start
 * @api {products}
 */

func CreateProduct(product *models.MerchProduct) error {
	_, err := database.DB.Exec(context.Background(), `
		INSERT INTO merch.product (title, description, category_id, primary_image_id, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)`,
		product.Title, product.Description, product.CategoryID, product.PrimaryImageID, product.CreatedAt, product.UpdatedAt)
	return err
}

func UpdateProduct(productID int, updatedProduct *models.MerchProduct) error {
	_, err := database.DB.Exec(context.Background(), `
		UPDATE merch.product SET title = $1, description = $2, category_id = $3, primary_image_id = $4, updated_at = $5
		WHERE id = $6`,
		updatedProduct.Title, updatedProduct.Description, updatedProduct.CategoryID, updatedProduct.PrimaryImageID, updatedProduct.UpdatedAt, productID)
	return err
}

func DeleteProduct(productID int) error {
	_, err := database.DB.Exec(context.Background(), `
		DELETE FROM merch.product WHERE id = $1`, productID)
	return err
}

func ListProducts() ([]*models.MerchProduct, error) {
	rows, err := database.DB.Query(context.Background(), `
		SELECT id, title, description, category_id, primary_image_id, created_at, updated_at FROM merch.product`)
	if err != nil {
		return nil, err
	}

	var products []*models.MerchProduct
	for rows.Next() {
		var product models.MerchProduct
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.CategoryID, &product.PrimaryImageID, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func GetProductByID(productID int) (*models.MerchProduct, error) {
	var product models.MerchProduct

	// Query the product data
	err := database.DB.QueryRow(context.Background(), `
        SELECT 
            p.id, p.title, p.description, p.category_id, p.primary_image_id, p.created_at, p.updated_at
        FROM merch.product p
        WHERE p.id = $1`, productID).
		Scan(&product.ID, &product.Title, &product.Description, &product.CategoryID, &product.PrimaryImageID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Query and associate sizes
	sizes, err := GetSizeProduct(productID)
	if err != nil {
		return nil, err
	}
	product.Size = sizes

	// Query and associate colors
	colors, err := GetColorProduct(productID)
	if err != nil {
		return nil, err
	}
	product.Color = colors

	// Query and associate prices
	prices, err := GetProductPrice(productID)
	if err != nil {
		return nil, err
	}
	product.Price = prices

	return &product, nil
}

func GetSizeProduct(productID int) ([]*models.MerchSize, error) {
	rows, err := database.DB.Query(context.Background(), `
		SELECT id, product_id, size, created_at, updated_at FROM merch.size WHERE product_id = $1`, productID)
	if err != nil {
		return nil, err
	}

	var sizes []*models.MerchSize
	for rows.Next() {
		var size models.MerchSize
		err := rows.Scan(&size.ID, &size.ProductID, &size.Name, &size.CreatedAt, &size.UpdatedAt)
		if err != nil {
			return nil, err
		}

		sizes = append(sizes, &size)
	}

	return sizes, nil
}

func GetColorProduct(productID int) ([]*models.MerchColor, error) {
	rows, err := database.DB.Query(context.Background(), `
		SELECT id, product_id, color, created_at, updated_at FROM merch.color WHERE product_id = $1`, productID)
	if err != nil {
		return nil, err
	}

	var colors []*models.MerchColor
	for rows.Next() {
		var color models.MerchColor
		err := rows.Scan(&color.ID, &color.ProductID, &color.Name, &color.CreatedAt, &color.UpdatedAt)
		if err != nil {
			return nil, err
		}

		colors = append(colors, &color)
	}

	return colors, nil
}

func GetProductPrice(productID int) ([]*models.MerchPrice, error) {
	rows, err := database.DB.Query(context.Background(), `
        SELECT id, product_id, price, created_at, updated_at FROM merch.price WHERE product_id = $1`, productID)
	if err != nil {
		return nil, err
	}

	var prices []*models.MerchPrice
	for rows.Next() {
		var price models.MerchPrice
		err := rows.Scan(&price.ID, &price.ProductID, &price.Price, &price.CreatedAt, &price.UpdatedAt)
		if err != nil {
			return nil, err
		}

		prices = append(prices, &price)
	}

	return prices, nil
}

/**
 * end
 * @api {products}
 */

/**
 * start database
 * @api {categories}
 */

func CreateCategory(category *models.MerchCategory) error {
	_, err := database.DB.Exec(context.Background(), `
		INSERT INTO merch.category (name, created_at, updated_at) 
		VALUES ($1, $2, $3)`,
		category.Name, category.CreatedAt, category.UpdatedAt)
	return err
}

func UpdateCategory(categoryID int, updatedCategory *models.MerchCategory) error {
	_, err := database.DB.Exec(context.Background(), `
		UPDATE merch.category SET name = $1, updated_at = $2
		WHERE id = $3`,
		updatedCategory.Name, updatedCategory.UpdatedAt, categoryID)
	return err
}

func DeleteCategory(categoryID int) error {
	_, err := database.DB.Exec(context.Background(), `
		DELETE FROM merch.category WHERE id = $1`, categoryID)
	return err
}

func ListCategories() ([]*models.MerchCategory, error) {
	rows, err := database.DB.Query(context.Background(), `
		SELECT id, name, created_at, updated_at FROM merch.category`)
	if err != nil {
		return nil, err
	}

	var categories []*models.MerchCategory
	for rows.Next() {
		var category models.MerchCategory
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}

		categories = append(categories, &category)
	}

	return categories, nil
}

func GetCategoryByID(categoryID int) (*models.MerchCategory, error) {
	var category models.MerchCategory
	err := database.DB.QueryRow(context.Background(), `
		SELECT id, name, created_at, updated_at FROM merch.category WHERE id = $1`, categoryID).
		Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

/**
 * end database
 * @api {categories}
 */

/**
 * start database
 * @api {sizes}
 */

func CreateSize(size *models.MerchSize) error {
	_, err := database.DB.Exec(context.Background(), `
		INSERT INTO merch.size (product_id, size, created_at, updated_at) 
		VALUES ($1, $2, $3, $4)`,
		size.ProductID, size.Name, size.CreatedAt, size.UpdatedAt)
	return err
}

func UpdateSize(sizeID int, updatedSize *models.MerchSize) error {
	_, err := database.DB.Exec(context.Background(), `
		UPDATE merch.size SET product_id = $1, size = $2, updated_at = $3
		WHERE id = $4`,
		updatedSize.ProductID, updatedSize.Name, updatedSize.UpdatedAt, sizeID)
	return err
}

func DeleteSize(sizeID int) error {
	_, err := database.DB.Exec(context.Background(), `
		DELETE FROM merch.size WHERE id = $1`, sizeID)
	return err
}

func GetSizeByID(sizeID int) (*models.MerchSize, error) {
	var size models.MerchSize
	err := database.DB.QueryRow(context.Background(), `
		SELECT id, product_id, size, created_at, updated_at FROM merch.size WHERE id = $1`, sizeID).
		Scan(&size.ID, &size.ProductID, &size.Name, &size.CreatedAt, &size.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &size, nil
}

/**
 * end database
 * @api {sizes}
 */

/**
 * start database
 * @api {colors}
 */

func CreateColor(color *models.MerchColor) error {
	_, err := database.DB.Exec(context.Background(), `
		INSERT INTO merch.color (product_id, color, created_at, updated_at) 
		VALUES ($1, $2, $3, $4)`,
		color.ProductID, color.Name, color.CreatedAt, color.UpdatedAt)
	return err
}

func UpdateColor(colorID int, updatedColor *models.MerchColor) error {
	_, err := database.DB.Exec(context.Background(), `
		UPDATE merch.color SET product_id = $1, color = $2, updated_at = $3
		WHERE id = $4`,
		updatedColor.ProductID, updatedColor.Name, updatedColor.UpdatedAt, colorID)
	return err
}

func DeleteColor(colorID int) error {
	_, err := database.DB.Exec(context.Background(), `
		DELETE FROM merch.color WHERE id = $1`, colorID)
	return err
}

func GetColorByID(colorID int) (*models.MerchColor, error) {
	var color models.MerchColor
	err := database.DB.QueryRow(context.Background(), `
		SELECT id, product_id, color, created_at, updated_at FROM merch.color WHERE id = $1`, colorID).
		Scan(&color.ID, &color.ProductID, &color.Name, &color.CreatedAt, &color.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &color, nil
}

/**
 * end database
 * @api {colors}
 */

/**
 * start database
 * @api {transactions}
 */

func CreateTransaction(transaction *models.MerchTransaction) error {
	_, err := database.DB.Exec(context.Background(), `
		INSERT INTO merch.transaction (user_id, product_id, coupon_id, proof_of_payment, status, paid_at, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		transaction.UserID, transaction.ProductID, transaction.CouponID, transaction.ProofOfPayment, transaction.Status, transaction.PaidAt, transaction.CreatedAt, transaction.UpdatedAt)
	return err
}

func ListTransactions() ([]*models.MerchTransaction, error) {
	rows, err := database.DB.Query(context.Background(), `
		SELECT id, user_id, product_id, coupon_id, proof_of_payment, status, paid_at, created_at, updated_at FROM merch.transaction`)
	if err != nil {
		return nil, err
	}

	var transactions []*models.MerchTransaction
	for rows.Next() {
		var transaction models.MerchTransaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.CouponID, &transaction.ProofOfPayment, &transaction.Status, &transaction.PaidAt, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

func GetTransaction(transactionID int) (*models.MerchTransaction, error) {
	var transaction models.MerchTransaction
	err := database.DB.QueryRow(context.Background(), `
		SELECT id, user_id, product_id, coupon_id, proof_of_payment, status, paid_at, created_at, updated_at FROM merch.transaction WHERE id = $1`, transactionID).
		Scan(&transaction.ID, &transaction.UserID, &transaction.ProductID, &transaction.CouponID, &transaction.ProofOfPayment, &transaction.Status, &transaction.PaidAt, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

/**
 * end database
 * @api {transactions}
 */
