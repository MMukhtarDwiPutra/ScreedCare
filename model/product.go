package model

type ProductResponse struct{
	ID int `json:"id" db:"id"`
	Foto string `db:"foto" json:"foto"`
	Deskripsi string `db:"deskripsi" json:"deskripsi"`
	Url string `db:"url" json:"url"`
}

type ProductRequest struct{
	Foto string `db:"foto" json:"foto"`
	Deskripsi string `db:"deskripsi" json:"deskripsi"`
	Url string `db:"url" json:"url"`
}