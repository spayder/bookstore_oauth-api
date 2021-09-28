package http

import (
	"githab.com/spayder/bookstore_oauth-api/src/domain/access_token"
	"githab.com/spayder/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := c.Param("access_token_id")

	accessToken, err := h.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var token access_token.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Code, restErr)
		return
	}

	if err := h.service.Create(token); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}