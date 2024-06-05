package models

type Gallery struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Intro       string `json:"intro"`
	Cover       string `json:"cover"`
	Create_time string `json:"create_time"`
}

type Image struct {
	Url string `json:"url"`
}

type GalleryDetail struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Intro       string   `json:"intro"`
	Cover       string   `json:"cover"`
	Create_time string   `json:"create_time"`
	Images      []*Image `json:"images"`
}
