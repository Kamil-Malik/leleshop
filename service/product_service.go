package service

import (
	"fmt"
	dto "leleshop/dto/product"
	entity "leleshop/entity/product"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProductDto and ProductImageDto structs (as shown earlier)

// AddProduct adds a new product to the database
func AddProduct(db *gorm.DB, productDto *dto.ProductDto) error {
	// Convert ProductDto to ProductEntity (you can also create a function to do this conversion)
	productEntity := entity.ProductEntity{
		SellerID:    productDto.SellerID,
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       productDto.Price,
		Quantity:    productDto.Quantity,
		CategoryID:  productDto.CategoryID,
		IsSold:      productDto.IsSold,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save the product to the database
	result := db.Create(&productEntity)
	if result.Error != nil {
		return result.Error
	}

	// You can also add the product images here if needed

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
	result := db.Where("seller_id = ?", userID).Find(&products)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no products found for the user")
		}
		return nil, result.Error
	}

	// Convert ProductEntity slice to ProductDto slice (you can also create a function to do this conversion)
	var productDtos []dto.ProductDto
	for _, productEntity := range products {
		productDtos = append(productDtos, dto.ProductDto{
			ID:          productEntity.ID,
			SellerID:    productEntity.SellerID,
			Name:        productEntity.Name,
			Description: productEntity.Description,
			Price:       productEntity.Price,
			Quantity:    productEntity.Quantity,
			CategoryID:  productEntity.CategoryID,
			IsSold:      productEntity.IsSold,
		})
	}

	// Retrieve product images here for each product if needed

	return productDtos, nil
}

// Example usage:
func main() {
	dsn := "host=localhost user=your_user password=your_password dbname=your_database port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Run migrations (if needed)
	db.AutoMigrate(&entity.ProductEntity{}, &entity.ProductEntity{})

	// Example: Add a product
	productDto := &dto.ProductDto{
		SellerID:    1,
		Name:        "Sample Product",
		Description: "This is a sample product",
		Price:       19.99,
		Quantity:    100,
		CategoryID:  1,
		IsSold:      false,
	}
	err = AddProduct(db, productDto)
	if err != nil {
		log.Fatal("failed to add product: ", err)
	}

	// Example: Get a product by ID
	productID := uint(1)
	product, err := GetProductByID(db, productID)
	if err != nil {
		log.Fatal("failed to get product: ", err)
	}
	fmt.Println("Product:", product)

	// Example: Get all products for a specific user ID
	userID := uint(1)
	products, err := GetAllProductsForUser(db, userID)
	if err != nil {
		log.Fatal("failed to get products for user: ", err)
	}
	fmt.Println("Products for User ID:", userID)
	for _, p := range products {
		fmt.Println(p)
	}
}
