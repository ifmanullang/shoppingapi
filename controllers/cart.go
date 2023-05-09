package controllers

import (
	"irfan/shoppingcartapi/database"
	"irfan/shoppingcartapi/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type TransaksiController struct {
	Db *gorm.DB
}

// login dahulu sebelum transaksi
func (controller *TransaksiController) GetTransaksi(c *fiber.Ctx) error {
	userId := c.Query("userid")
	user_id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(fiber.Map{
			"message": "Anda harus login",
		})
	}

	var transaksi []models.Transaksi
	err = models.ViewTransaksi(controller.Db, &transaksi, user_id)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(transaksi)
}
// melakukan pembelian dengan id user dan id product yg akan di beli
func (controller *TransaksiController) AddtoTransaksi(c *fiber.Ctx) error {
	userId := c.Query("userid")
	productId := c.Query("productid")

	product_id,_ := strconv.Atoi(productId)
	user_id,_ := strconv.Atoi(userId)
  
	var transaksi models.Transaksi
	var product models.Product
	var find models.Transaksi

	err2 := models.ReadProductById(controller.Db, &product, product_id)

	if err2 != nil {
		c.JSON(fiber.Map{
			"message": "Product ini tidak ada",// jika salah memasukkan id product
			"status": 500,
		})
	}

	err4 := models.FindTransaksi(controller.Db, &find, product_id, user_id)
	if err4 != nil {
		c.JSON(fiber.Map{
			"message": "Tidak ada transaksi",// belum melakukan transaksi
			"status": 500,
		})
	}

	if find.Id != 0 {
		find.Quantity = find.Quantity + 1
		find.Total = find.Total + product.Price

		models.UpdateTransaksi(controller.Db, &find)

		return c.JSON(find)		
	} else {
		transaksi.ProductId = product_id
		transaksi.UserId = user_id
		transaksi.Quantity = 1
		transaksi.Total = float64(transaksi.Quantity)*product.Price

		err3 := models.AddtoTransaksi(controller.Db, &transaksi)

		if err3 != nil {
			c.JSON(fiber.Map{
				"message": "Gagal melakukan transaksi",
				"status": 500,
			})
		}

		return c.JSON(transaksi)
	}
	
}

func InitTransaksiController() *TransaksiController {
	db := database.InitDb()

	db.AutoMigrate(&models.Transaksi{})

	return &TransaksiController{Db: db}
}