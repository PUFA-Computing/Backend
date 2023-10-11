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
	err := database.DB.QueryRow(context.Background(), `
		SELECT id, title, description, category_id, primary_image_id, created_at, updated_at FROM merch.product WHERE id = $1`, productID).
		Scan(&product.ID, &product.Title, &product.Description, &product.CategoryID, &product.PrimaryImageID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

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

func GetProductPrice(productID int) (*models.MerchPrice, error) {
	var price models.MerchPrice
	err := database.DB.QueryRow(context.Background(), `
		SELECT id, product_id, price, created_at, updated_at FROM merch.price WHERE product_id = $1`, productID).
		Scan(&price.ID, &price.ProductID, &price.Price, &price.CreatedAt, &price.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &price, nil
}

/**
 * end
 * @api {products}
 */
