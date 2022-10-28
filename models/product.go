package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id       uint     `form:"id" json:"id" validate:"required"`
	Name     string  `form:"name" json:"name" validate:"required"`
	Quantity int     `form:"quantity" json:"quantity" validate:"required"`
	Image 	string `form:"image" json:"image" validate:"required"`
	Price    float64 `form:"price" json:"price" validate:"required"`
}

func CreateProduct(db *gorm.DB, newProduct *Product) (err error) {
	err = db.Create(newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProducts(db *gorm.DB, products *[]Product) (err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProductById(db *gorm.DB, products *Product, id int) (err error) {
	err = db.Where("id=?", id).First(products).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, products *Product) (err error) {

	

	db.Save(products)
	
	return nil
}

func DeleteProduct(db *gorm.DB, product *Product, id int) (err error) {
	err = db.Where("id=?", id).Delete(product).Error
	if err != nil {
		return err
	}
	return nil
}