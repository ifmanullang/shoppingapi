package main

import (
	"irfan/shoppingcartapi/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	store := session.New()
	app := fiber.New()

	productAPI := controllers.InitProductController()
	auth := controllers.InitAuthController(store)
	transaksi := controllers.InitTransaksiController()

	p := app.Group("/")
	p.Get("/products", productAPI.IndexProducts)
	p.Get("/productsdetail/:id", productAPI.DetailProduct)
	p.Post("/addproducts", productAPI.CreateProduct)
	p.Put("/editproducts/:id", productAPI.EditProduct)
	p.Delete("/deleteproducts/:id", productAPI.DeleteProduct)

	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)

	app.Get("/transaksi", transaksi.GetTransaksi)
	app.Post("/transaksi", transaksi.AddtoTransaksi)


	app.Listen(":3000")
}