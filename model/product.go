package model

type ProductResponse struct{
	ID int `json:"id" db:"id"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Foto string `db:"foto" json:"foto"`
	Harga string `db:"harga" json:"harga"`
	Url string `db:"url" json:"url"`
}

type ProductRequest struct{
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Foto string `db:"foto" json:"foto"`
	Harga string `db:"harga" json:"harga"`
	Url string `db:"url" json:"url"`
}