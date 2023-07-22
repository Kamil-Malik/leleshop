package service

import (
	"errors"
	"fmt"
	dto "leleshop/dto/product"
	"leleshop/dto/response"
	entity "leleshop/entity/product"
	"leleshop/mapper"
	"math"

	"gorm.io/gorm"
)

// ProductDto and ProductImageDto structs (as shown earlier)

// AddProduct adds a new product to the database
func AddProduct(db *gorm.DB, productDto *dto.ProductDto) error {

	// Return error if the product image is empty
	if len(productDto.Images) == 0 {
		return errors.New("product image cannot be empty")
	}

	// Convert ProductDto to ProductEntity (you can also create a function to do this conversion)
	productEntity := mapper.ToProductEntity(*productDto)

	// Save the product to the database
	if err := db.Create(&productEntity).Error; err != nil {
		return err
	}

	// You can also add the product images here if needed
	for _, img := range productDto.Images {
		productImageEntity := mapper.ToProductImageEntity(img)
		productImageEntity.ProductID = productEntity.ID // Set the foreign key (ProductID) to link the product image to the product
		if err := db.Create(&productImageEntity).Error; err != nil {
			// If there's an error, you may choose to handle it or return it
			return err
		}
	}

	return nil
}

// GetProductByID retrieves a single product by its ID
func GetProductByID(db *gorm.DB, productID uint) (*dto.ProductDto, error) {
	var productEntity entity.ProductEntity

	// Find the product by its ID
	result := db.First(&productEntity, productID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("product not found")
		}
		return nil, result.Error
	}

	// Convert ProductEntity to ProductDto (you can also create a function to do this conversion)
	productDto := dto.ProductDto{
		ID:          productEntity.ID,
		SellerID:    productEntity.SellerID,
		Name:        productEntity.Name,
		Description: productEntity.Description,
		Price:       productEntity.Price,
		Quantity:    productEntity.Quantity,
		CategoryID:  productEntity.CategoryID,
		IsSold:      productEntity.IsSold,
	}

	// Retrieve product images here if needed

	return &productDto, nil
}

// GetAllProductsForUser retrieves all products for a specific user ID
func GetAllProductsForUser(db *gorm.DB, userID uint) ([]dto.ProductDto, error) {
	var products []entity.ProductEntity

	// Find all products for the specific user ID
	if err := db.Preload("Images").Where("seller_id = ?", userID).Find(&products).Error; err != nil {
		return nil, errors.New("no products found for the user")
	}

	// Convert ProductEntity slice to ProductDto slice (you can also create a function to do this conversion)
	var productDtos []dto.ProductDto
	for _, productEntity := range products {
		// Convert ProductEntity to ProductDto (you can create a mapper function for this)
		productDto := dto.ProductDto{
			ID:          productEntity.ID,
			SellerID:    productEntity.SellerID,
			Name:        productEntity.Name,
			Description: productEntity.Description,
			Price:       productEntity.Price,
			Quantity:    productEntity.Quantity,
			CategoryID:  productEntity.CategoryID,
			IsSold:      productEntity.IsSold,
		}

		// Convert ProductImageEntity slice to ProductImageDto slice (you can create a mapper function for this)
		for _, img := range productEntity.Images {
			productDto.Images = append(productDto.Images, dto.ProductImageDto{
				ID:        img.ID,
				ProductID: img.ProductID,
				ImageURL:  img.ImageURL,
			})
		}

		productDtos = append(productDtos, productDto)
	}

	return productDtos, nil
}

// GetProductsPagination retrieves products along with their associated images in a pagination form
func GetProductsPagination(db *gorm.DB, page, pageSize int) (products []dto.ProductDto, paginationResponse response.PaginationItemResponse, err error) {

	// Calculate the offset for pagination
	offset := (page - 1) * pageSize

	// Query products along with their associated images using Preload
	result := db.Preload("Images").Limit(pageSize).Offset(offset).Find(&products)
	if result.Error != nil {
		return nil, response.PaginationItemResponse{}, result.Error
	}

	// Convert ProductEntity slice to ProductDto slice (you can also create a function to do this conversion)
	var productDtos []dto.ProductDto
	for _, productEntity := range products {
		productDto := dto.ProductDto{
			ID:          productEntity.ID,
			SellerID:    productEntity.SellerID,
			Name:        productEntity.Name,
			Description: productEntity.Description,
			Price:       productEntity.Price,
			Quantity:    productEntity.Quantity,
			CategoryID:  productEntity.CategoryID,
			IsSold:      productEntity.IsSold,
		}

		// Convert ProductImageEntity slice to ProductImageDto slice (you can create a mapper function for this)
		for _, img := range productEntity.Images {
			productDto.Images = append(productDto.Images, dto.ProductImageDto{
				ID:        img.ID,
				ProductID: img.ProductID,
				ImageURL:  img.ImageURL,
			})
		}

		productDtos = append(productDtos, productDto)
	}

	// Calculate pagination metadata
	var totalItems int64
	db.Model(&entity.ProductEntity{}).Count(&totalItems)
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	// Populate the PaginationItemResponse struct
	paginationResponse = response.PaginationItemResponse{
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		PageSize:    pageSize,
		CurrentPage: page,
	}

	return productDtos, paginationResponse, nil
}
