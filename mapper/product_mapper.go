package mapper

import (
	dto "leleshop/dto/product"
	entity "leleshop/entity/product"
)

func ToProductImageDto(product entity.ProductImageEntity) dto.ProductImageDto {
	return dto.ProductImageDto{
		ID:        product.ID,
		ProductID: product.ProductID,
		ImageURL:  product.ImageURL,
	}
}

// ToProductImageEntity converts ProductImageDto to ProductImageEntity
func ToProductImageEntity(product dto.ProductImageDto) entity.ProductImageEntity {
	return entity.ProductImageEntity{
		ID:        product.ID,
		ProductID: product.ProductID,
		ImageURL:  product.ImageURL,
	}
}

func ToProductDto(entity entity.ProductEntity) dto.ProductDto {
	dto := dto.ProductDto{
		ID:          entity.ID,
		SellerID:    entity.SellerID,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		Quantity:    entity.Quantity,
		CategoryID:  entity.CategoryID,
		IsSold:      entity.IsSold,
	}

	// Convert ProductImageEntity slice to ProductImageDto slice
	for _, img := range entity.Images {
		dto.Images = append(dto.Images, ToProductImageDto(img))
	}

	return dto
}

// ToProductEntity converts ProductDto to ProductEntity
func ToProductEntity(dto dto.ProductDto) entity.ProductEntity {
	entity := entity.ProductEntity{
		ID:          dto.ID,
		SellerID:    dto.SellerID,
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Quantity:    dto.Quantity,
		CategoryID:  dto.CategoryID,
		IsSold:      dto.IsSold,
	}

	// Convert ProductImageDto slice to ProductImageEntity slice
	for _, img := range dto.Images {
		entity.Images = append(entity.Images, ToProductImageEntity(img))
	}

	return entity
}
