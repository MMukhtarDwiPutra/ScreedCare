package repository

import(
	"ScreedCare/model"
	"gorm.io/gorm"
)

type ProductRepository interface{
	GetAllProduct() []model.ProductResponse
	AddProduct(p model.ProductRequest)
	DeleteProduct(id int)
	FindProduct(id int) model.ProductResponse
	UpdateProduct(p model.ProductRequest, id int)
}

type productRepository struct{
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository{
	return &productRepository{db}
}

func (r *productRepository) GetAllProduct() []model.ProductResponse{
	// rows, err := r.db.Query("SELECT * FROM product")
	// helper.PanicIfError(err)

	var products []model.ProductResponse
	r.db.Table("product").Find(&products)

	// for rows.Next(){
	// 	var p model.ProductResponse
	// 	rows.Scan(&p.ID, &p.NamaBarang, &p.Foto, &p.Harga, &p.LinkShopee, &p.LinkTokopedia, &p.LinkSociolla)

	// 	products = append(products, p)
	// }

	return products
}

func (r *productRepository) AddProduct(p model.ProductRequest) {
	// r.db.Query("INSERT INTO product (nama_barang, foto, link_shopee, harga, link_tokopedia, link_sociolla) VALUES (?, ?, ?, ?, ?, ?)",p.NamaBarang, p.Foto, p.LinkShopee, p.Harga, p.LinkTokopedia, p.LinkSociolla)
	r.db.Table("product").Create(&p)
	
}

func (r *productRepository) DeleteProduct(id int){
	// r.db.Query("DELETE FROM product WHERE id = ?",id)
	var p model.ProductResponse
	r.db.Table("product").Where("id = ?",id).Delete(&p)
}

func (r *productRepository) FindProduct(id int) model.ProductResponse{
	var p model.ProductResponse
	// row := r.db.QueryRow("SELECT * FROM product WHERE id = ?", id)
	
	// row.Scan(&p.ID, &p.NamaBarang, &p.Foto, &p.Harga, &p.LinkShopee, &p.LinkTokopedia, &p.LinkSociolla)
	r.db.Table("product").Where("id = ?",id).First(&p)
	return p
}

func (r *productRepository) UpdateProduct(p model.ProductRequest, id int) {
	// _, err := r.db.Query("UPDATE product SET nama_barang = ?, foto = ?, harga = ?, link_shopee = ?, link_tokopedia = ?, link_sociolla = ? WHERE id = ?",p.NamaBarang, p.Foto, p.Harga, p.LinkShopee, p.LinkTokopedia, p.LinkSociolla, id)

	r.db.Table("product").Where("id = ?", id).Updates(&p)
}