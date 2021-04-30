package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/majustman/todo-app"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
