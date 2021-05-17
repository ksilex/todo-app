package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.sign_up)
		auth.POST("/sign-in", h.sign_in)
	}

	api := router.Group("/api")
	{
		lists := api.Group("lists")
		{
			lists.POST("/", h.create_list)
			lists.GET("/", h.get_all_lists)
			lists.GET("/:id", h.get_list)
			lists.PUT("/:id", h.update_list)
			lists.DELETE("/:id", h.delete_list)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.create_item)
				items.GET("/", h.get_all_items)
				items.GET("/:item_id", h.get_item)
				items.PUT("/:item_id", h.update_item)
				items.DELETE("/:item_id", h.delete_item)
			}
		}
	}
	return router
}
