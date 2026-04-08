package models

import "gorm.io/gorm"

//Struct Product yang mapping ke table products di MySQL
type Product struct {
	gorm.Model
	Name        string  `gorm:"size:200;not null;index" json:"name"` //nama produk, ada index untuk pencarian
	Description string  `gorm:"type:text" json:"description"`        //deskripsi produk
	Price       float64 `gorm:"not null" json:"price"`               //harga produk, gunakan decimal untuk produksi
	Stock       int     `gorm:"default:0" json:"stock"`              //stok produk dengan default 0
	Category    string  `gorm:"size:100;index" json:"category"`      //kategori produk ada index untuk pencarian
	ImageURL    string  `gorm:"size:500" json:"image_url"`           //URL gambar produk
	IsActive    bool    `gorm:"default:true" json:"is_active"`       //untuk produk agar tampil atau tidak dengan value 1 untuk true dan 0 untuk false dan true dijadikan value default
}

//Request/Response DTO (Data Transfer Object) untuk validasi dari HTTP request dan response JSON
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=2,max=200"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	Category    string  `json:"category" binding:"required"`
	ImageURL    string  `json:"image_url"`
}

type UpdateProductRequest struct {
	Name        *string  `json:"name" binding:"omitempty,min=2"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
	Stock       *int     `json:"stock" binding:"omitempty,min=0"`
	Category    *string  `json:"category"`
	ImageURL    *string  `json:"image_url"`
}
