package model

type ProductData struct{
	Data []ProductRequest `json:"data"`
}

type ProductResponse struct{
	ID int `json:"id" db:"id"`
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Foto string `db:"foto" json:"foto"`
	Harga int `db:"harga" json:"harga"`
	HargaStr string `json:"harga_str"`
	LinkShopee string `db:"link_shopee" json:"link_shopee"`
	LinkTokopedia string `db:"link_tokopedia" json:"link_tokopedia"`
	LinkSociolla string `db:"link_sociolla" json:"link_sociolla"`
}

type ProductRequest struct{
	NamaBarang string `json:"nama_barang" db:"nama_barang"`
	Foto string `db:"foto" json:"foto"`
	Harga int `db:"harga" json:"harga"`
	LinkShopee string `db:"link_shopee" json:"link_shopee"`
	LinkTokopedia string `db:"link_tokopedia" json:"link_tokopedia"`
	LinkSociolla string `db:"link_sociolla" json:"link_sociolla"`
}