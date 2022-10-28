package controllers

import (
	"fmt"
	"irfan/shoppingcartapi/database"
	"irfan/shoppingcartapi/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController struct {
	Db *gorm.DB
}

// menambahkan product
func (controller *ProductController) CreateProduct(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["image"]

		for _, file := range files {
			var data models.Product
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			if err := c.BodyParser(&data); err != nil {
				return c.JSON(fiber.Map{
					"message": "Semua data harus di isi",
				})
			}

			if err := c.SaveFile(file, fmt.Sprintf("./public/images/%s", file.Filename)); err != nil {
				return err
			}

			data.Image = file.Filename

			err := models.CreateProduct(controller.Db, &data)

			if err != nil {
				return c.JSON(data)
			}

			c.JSON(data)
		}
		return c.JSON(fiber.Map{
			"message": "Product berhasil ditambahkan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product gagal ditambahkan",
	})
}

func (controller *ProductController) IndexProducts(c *fiber.Ctx) error {
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(products)
}

// Detail Product
func (controller *ProductController) DetailProduct(c *fiber.Ctx) error {
	var product models.Product
	var id, err = c.ParamsInt("id")

	if err != nil {
		return err
	}

	err2 := models.ReadProductById(controller.Db, &product, id)

	if err2 != nil {
		return c.SendStatus(500)
	}

	return c.JSON(product)
}

// Update Product
func (controller *ProductController) EditProduct(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["image"]

		for _, file := range files {
			var data models.Product

			id := c.Params("id")
			idn, _ := strconv.Atoi(id)

			err := models.ReadProductById(controller.Db, &data, idn)
			if err != nil {
				return c.SendStatus(500) // http 500 internal server error
			}

			var myform models.Product

			if err := c.BodyParser(&myform); err != nil {
				return c.JSON(fiber.Map{
					"message": "Data product harus lengkap",
				})
			}

			data.Name = myform.Name
			data.Quantity = myform.Quantity
			data.Price = myform.Price

			// simpan product yang di update
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

			if err := c.SaveFile(file, fmt.Sprintf("./public/images/%s", file.Filename)); err != nil {
				return err
			}

			data.Image = file.Filename

			err = models.UpdateProduct(controller.Db, &data)

			if err != nil {
				return c.JSON(data)
			}

			c.JSON(data)
		}
		return c.JSON(fiber.Map{
			"message": "Product berhasil di perbaharui!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "error",
	})
}

// menghapus proudct
func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	var product models.Product
	var id, err = c.ParamsInt("id")

	if err != nil {
		return err
	}

	err2 := models.DeleteProduct(controller.Db, &product, id)

	if err2 != nil {
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil Menghapus!",
	})
}

func InitProductController() *ProductController {
	db := database.InitDb()
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db}
}
