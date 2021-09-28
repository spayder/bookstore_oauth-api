package app

import (
	"githab.com/spayder/bookstore_oauth-api/src/domain/access_token"
	"githab.com/spayder/bookstore_oauth-api/src/domain/repository/db"
	"githab.com/spayder/bookstore_oauth-api/src/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Handle() {
	service := access_token.NewService(db.NewRepository())
	handler := http.NewHandler(service)

	router.GET("/oauth/access_token/:access_token_id", handler.GetById)

	router.Run(":8092")
}