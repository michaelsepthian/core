package getallproducts

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB) ([]Response, *others.BaseResponse) {
	var errorResponse *others.BaseResponse

	product, err := GetAllProducts(db)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Empty Data",
			StatusCode: http.StatusNotFound,
		}
		return []Response{}, errorResponse
	}

	responseList := MapProductsToResponse(product)

	return responseList, errorResponse

}

func MapProductsToResponse(products []models.Products) []Response {
	var responselist []Response
	for _, product := range products {
		allProducts := Response{
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}
		responselist = append(responselist, allProducts)
	}
	return responselist
}
