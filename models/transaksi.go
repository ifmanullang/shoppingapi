package models

import "gorm.io/gorm"


type Transaksi struct {
	gorm.Model
	Id       		uint     `json:"id" validate:"required"`
	ProductId     	int  `form:"productid" json:"productid" validate:"required"`
	UserId     		int  `form:"userid" json:"userid" validate:"required"`
	Quantity 		int     `form:"quantity" json:"quantity" validate:"required"`
	Total    		float64 `form:"total" json:"total" validate:"required"`
	Product 		Product	`gorm:"foreignkey:ProductId;references:Id"`
	User  			User `gorm:"foreignkey:UserId;references:Id"`
}

func ViewTransaksi(db *gorm.DB, transaksi *[]Transaksi, id int) (err error) {
	err = db.Where(&Transaksi{UserId: id,}).Preload("User").Preload("Product").Find(transaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func FindTransaksi(db *gorm.DB, transaksi *Transaksi, product int, user int) (err error) {
	err = db.Where(&Transaksi{ProductId: product, UserId: user,}).Preload("User").Preload("Product").Find(transaksi).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTransaksi(db *gorm.DB, transaksi *Transaksi) (err error) {
	db.Save(transaksi)
	
	return nil
}

func AddtoTransaksi(db *gorm.DB, data *Transaksi) (err error) {
	err = db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTransaksi(db *gorm.DB, data *Transaksi, id int) (err error) {
	err = db.Where("id=?", id).Delete(data).Error
	if err != nil {
		return err
	}
	return nil
}