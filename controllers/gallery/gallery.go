package gallery

import (
	"net/http"
	"wow_blog/services"

	"github.com/gin-gonic/gin"
)

func GetGallery(c *gin.Context) {
	gallerys, _ := services.GetGallerys()
	c.JSON(http.StatusOK, gallerys)
}

func GetGalleryDetail(c *gin.Context) {
	gid := c.Query("id")
	gallery_detail, _ := services.GetGalleryDetail(gid)
	c.JSON(http.StatusOK, gallery_detail)
}
