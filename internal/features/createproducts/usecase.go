package createproducts

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, name string, price int, stock int) (models.Products, *others.BaseResponse) {
	var errorResponse *others.BaseResponse
	newProduct := models.Products{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	product, err := CreateProduct(db, newProduct)
	if err != nil {
		errorResponse = &others.BaseResponse{
			Message:    "Failed To Save Product",
			StatusCode: http.StatusInternalServerError,
		}
		return models.Products{}, errorResponse
	}

	return product, errorResponse
}
