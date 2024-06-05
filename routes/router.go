package routes

import (
	"wow_blog/controllers/gallery"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/gallery_list", gallery.GetGallery)
		api.GET("/gallery_detail", gallery.GetGalleryDetail)
	}
}
