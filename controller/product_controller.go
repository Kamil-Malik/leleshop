package controller

import (
	"leleshop/dto/product"
	"leleshop/dto/response"
	"leleshop/service"
	"leleshop/util"
	"net/http"
	"strconv"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func AddProduct(ctx *gin.Context) {
	var productDto product.ProductDto
	contentType := ctx.Request.Header.Get(util.ContentType)

	// Get JSON or Body Form from HTTP Request
	if contentType == util.ApplicationJson {
		if err := ctx.ShouldBindJSON(&productDto); err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				response.ErrorResponse{
					Status:  false,
					Message: err.Error(),
				},
			)
			return
		}
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: "This endpoint only accept a JSON",
			},
		)
		return
	}

	// Validate Struct with validator
	if _, err := valid.ValidateStruct(&productDto); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			},
		)
		return
	}

	if err := service.AddProduct(&productDto); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: "Request model didn't satisfy the require product",
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		response.Response{
			Status:  true,
			Message: "Ok",
			Data:    "Product created",
		},
	)
}

func GetProducts(ctx *gin.Context) {
	pageNumberQuery := ctx.Query("page")
	pageSizeQuery := ctx.Query("page_size")

	pageNumber, pageNumberErr := strconv.Atoi(pageNumberQuery)
	if pageNumberErr != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: "Page number should be a valid integer",
			},
		)
		return
	}

	pageSize, pageSizeError := strconv.Atoi(pageSizeQuery)
	if pageSizeError != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: "Page size should be a valid integer",
			},
		)
		return
	}

	products, pagination, err := service.GetProductsPagination(pageNumber, pageSize)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.PaginationResponse{
			Status:         true,
			Message:        "Ok",
			Data:           products,
			PaginationItem: pagination,
		},
	)
}
