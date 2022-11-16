package finishtransaction

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gorm.io/gorm"
	"net/http"
)

func UseCase(db *gorm.DB, transactionID uint64, method string) (Response, *others.BaseResponse) {
	var errResponse *others.BaseResponse

	_, err := FindTransactionByID(db, transactionID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Transaction Not Found",
			StatusCode: http.StatusNotFound,
		}
		return Response{}, errResponse
	}

	_, err = CheckTransaction(db, transactionID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Transaction Already Paid",
			StatusCode: http.StatusNotAcceptable,
		}
		return Response{}, errResponse
	}

	err = UpdateOrder(db, transactionID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Failed To Update Order Status",
			StatusCode: http.StatusNotAcceptable,
		}
		return Response{}, errResponse
	}

	_, err = UpdateTransaction(db, transactionID, method)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Transaction Failed",
			StatusCode: http.StatusBadRequest,
		}
		return Response{}, errResponse
	}

	response, err := GetTransaction(db, transactionID)
	if err != nil {
		errResponse = &others.BaseResponse{
			Message:    "Cannot  Get Transaction",
			StatusCode: http.StatusFound,
		}
		return Response{}, errResponse
	}

	return response, errResponse
}
