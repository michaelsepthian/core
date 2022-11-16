package updateproducts

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, productID uint64, name string, price int, stock int) (models.Products, *others.BaseResponse) {
	var errResponse *others.BaseResponse

	_, err := FindProductByID(db, productID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Product Cannot Found",
			StatusCode: http.StatusNotFound,
		}
		return models.Products{}, errResponse
	}

	updateProduct := models.Products{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	user, err := UpdateProduct(db, productID, updateProduct)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Failed To Update Data Product",
			StatusCode: http.StatusBadRequest,
		}
		return models.Products{}, errResponse
	}

	return user, errResponse
}
