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
