package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create_list(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) get_all_lists(c *gin.Context) {

}

func (h *Handler) get_list(c *gin.Context) {

}

func (h *Handler) update_list(c *gin.Context) {

}

func (h *Handler) delete_list(c *gin.Context) {

}
