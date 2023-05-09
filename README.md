# shoppingapi

### user 
melakukan registrasi => http://localhost:3000/register
          login      => http://localhost:3000/login
          
### product
  GET("http://localhost:3000/products") => menampilkan semua product
	Get("http://localhost:3000/productsdetail/:id") => menampilkan product berdasarkan id
	Post("http://localhost:3000/addproducts") => membuat product baru
	Put("http://localhost:3000/editproducts/:id") => mengedit data product
	Delete("http://localhost:3000/deleteproducts/:id") => menghapus product berdasarkan id

### transaksi
Get(http://localhost:3000/transaksi)
Post(http://localhost:3000/transaksi), masukkan user id dan id product untuk memasukkan product ke cart
