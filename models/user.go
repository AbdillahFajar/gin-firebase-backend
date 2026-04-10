package models

import "gorm.io/gorm"

// Membuat model User yang mapping ke table users di MySQL
// GORM otomatis plural nama struct User jadi nama tabel users
type User struct {
	gorm.Model
	FirebaseUID   string `gorm:"uniqueIndex;size:128;not null" json:"firebase_uid"` //UID dari firebase
	Email         string `gorm:"uniqueIndex;size:255;not null" json:"email"`        //Email user
	Name          string `gorm:"size:100" json:"name"`                              //Nama user
	Role          string `gorm:"size:20;default:user" json:"role"`                  //Hak akses: user/admin
	EmailVerified bool   `gorm:"default:false" json:"email_verified"`
	LastLoginAt   *int64 `gorm:"index" json:"last_login_at,omitempty"` //unix timestamp login terakhir
}

// gorm.Model otomatis memberikan fields:
// ID 	  uint           (primary key, auto-increment)
// CreatedAt time.Time (auto fill saat insert)
// UpdatedAt time.Time (auto fill saat update)
// DeletedAt gorm.DeletedAt (untuk soft delete, row tidak benar-benar dihapus)

// Struct tag "gorm" mengontrol perilaku GORM:
// uniqueIndex = buat unique index di kolom ini
// size:128 = varchar (128)
// not null = kolom wajib diisi atau tidak boleh NULL
// default:user = nilai default kolom

// Struct tag "json" mengontrol serialisasi JSON:
// json:"firebase_uid" = nama key di JSON response
// json:"-"            = field ini tidak dimasukkan ke JSON
// json:"..,omitempty" = skip field ini jika nilainya zero/nil atau default value
