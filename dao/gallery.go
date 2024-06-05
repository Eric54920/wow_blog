package dao

import (
	"wow_blog/models"
	"wow_blog/utils"
)

func GetGallery() (gallerys []*models.Gallery, err error) {
	sqlStr := `SELECT
		g.id,
		g.title,
		g.intro,
		g.create_time,
		gd.url AS cover 
	FROM
		gallery g
		LEFT JOIN gallery_detail gd ON g.cover = gd.id 
	ORDER BY
		create_time DESC;`
	gallerys = make([]*models.Gallery, 0, 3)
	err = utils.DB.Select(&gallerys, sqlStr)
	return
}

func GetGalleryDetail(gid string) (gallery_detail *models.GalleryDetail, err error) {
	sqlStr := `SELECT
		g.id,
		g.title,
		g.intro,
		gd.url AS cover,
		g.create_time 
	FROM
		gallery g
		JOIN gallery_detail gd ON g.id = gd.gallery_id 
	WHERE
		g.id =?`
	gallery := models.GalleryDetail{}

	_ = utils.DB.Get(&gallery, sqlStr, gid)

	imgStr := `SELECT url FROM gallery_detail WHERE gallery_id=?`
	images := make([]*models.Image, 0, 2)
	err = utils.DB.Select(&images, imgStr, gid)

	gallery_detail = &models.GalleryDetail{}
	gallery_detail.Id = gallery.Id
	gallery_detail.Title = gallery.Title
	gallery_detail.Cover = gallery.Cover
	gallery_detail.Intro = gallery.Intro
	gallery_detail.Create_time = gallery.Create_time
	gallery_detail.Images = images
	return
}
