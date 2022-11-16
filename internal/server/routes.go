package server

import (
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/cancelorder"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/createorder"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/createpayment"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/createproducts"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/createusers"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/deleteusers"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/finishtransaction"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/getallorders"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/getallproducts"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/getallusers"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/updateproducts"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/updateusers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) http.Handler {
	router := httprouter.New()

	router.POST("/create-users", createusers.New(db))
	router.POST("/create-products", createproducts.New(db))
	router.POST("/orders", createorder.New(db))
	router.POST("/create-payments", createpayment.New(db))

	router.GET("/users", getallusers.New(db))
	router.GET("/products", getallproducts.New(db))
	router.GET("/orders", getallorders.New(db))

	router.PUT("/update-users/:userId", updateusers.New(db))
	router.PUT("/update-products/:productId", updateproducts.New(db))

	router.PATCH("/transaction/:transactionId", finishtransaction.New(db))

	router.DELETE("/delete-users/:userId", deleteusers.New(db))
	router.DELETE("/cancel-orders", cancelorder.New(db))
	return router
}
