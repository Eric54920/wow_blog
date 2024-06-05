package services

import (
	"log"
	"wow_blog/dao"
	"wow_blog/models"
)

func GetGallerys() (gallerys []*models.Gallery, err error) {
	gallerys, err = dao.GetGallery()
	if err != nil {
		log.Fatalf("Select Gallerys Err: %v", err)
		return
	}
	return
}

func GetGalleryDetail(gid string) (gallery_detail *models.GalleryDetail, err error) {
	gallery_detail, err = dao.GetGalleryDetail(gid)
	if err != nil {
		log.Fatalf("Select Gallery Detail Err: %v", err)
		return
	}
	return
}
