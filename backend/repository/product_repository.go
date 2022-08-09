package repository

import(
	"ScreedCare/backend/helper"
	"ScreedCare/model"
	"database/sql"
)

type ProductRepository interface{
	GetAllProduct() []model.ProductResponse
	AddProduct(p model.ProductRequest)
	DeleteProduct(id int)
	FindProduct(id int) model.ProductResponse
	UpdateProduct(p model.ProductRequest, id int)
}

type productRepository struct{
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository{
	return &productRepository{db}
}

func (r *productRepository) GetAllProduct() []model.ProductResponse{
	rows, err := r.db.Query("SELECT * FROM product")
	helper.PanicIfError(err)

	var products []model.ProductResponse
	for rows.Next(){
		var p model.ProductResponse
		rows.Scan(&p.ID, &p.NamaBarang, &p.Foto, &p.Harga, &p.Url)

		products = append(products, p)
	}

	return products
}

func (r *productRepository) AddProduct(p model.ProductRequest) {
	r.db.Query("INSERT INTO product (nama_barang, foto, url, harga) VALUES (?, ?, ?)",p.NamaBarang, p.Foto, p.Url, p.Harga)
}

func (r *productRepository) DeleteProduct(id int){
	r.db.Query("DELETE FROM product WHERE id = ?",id)
}

func (r *productRepository) FindProduct(id int) model.ProductResponse{
	var p model.ProductResponse
	row := r.db.QueryRow("SELECT * FROM product WHERE id = ?", id)
	
	row.Scan(&p.ID, &p.NamaBarang, &p.Foto, &p.Harga, &p.Url)

	return p
}

func (r *productRepository) UpdateProduct(p model.ProductRequest, id int) {
	r.db.Query("UPDATE product SET nama_barang = ?, foto = ?, harga = ?, url = ? WHERE id = ?",p.NamaBarang, p.Foto, p.Harga, p.Url, id)
}