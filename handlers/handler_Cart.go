package handlers

import (
	"strconv"
	"yakz/config"
	"yakz/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AddToCart - เพิ่มหนังสือลงตะกร้า
func AddtoCart(c *fiber.Ctx) error {
	db := config.GetDB()

	var cartItem models.CartBook

	if err := c.BodyParser(&cartItem); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})

	}

	// ดึง user_id จาก JWT
	userData := c.Locals("user")
	if userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := userData.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	cartItem.UserID = uint(claims["user_id"].(float64))

	// ตรวจสอบว่าหนังสือมีอยู่หรือไม่
	var book models.Book

	if err := db.First(&book, cartItem.BookID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not Found"})
	}
	// ค้นหาว่าหนังสือนี้มีอยู่ในตะกร้าหรือยัง
	var existingItem models.CartBook
	if err := db.Preload("Book").Where("user_id = ? AND book_id = ?", cartItem.UserID, cartItem.BookID).First(&existingItem).Error; err == nil {

		// ถ้ามีอยู่แล้ว, เพิ่มจำนวน
		existingItem.Qty += cartItem.Qty
		db.Save(&existingItem)
		return c.JSON(existingItem)
	}

	// ถ้ายังไม่มี, ให้สร้างใหม่
	db.Create(&cartItem)

	//  ดึงข้อมูลอัปเดต พร้อม preload "Book"
	db.Preload("Book").First(&cartItem, cartItem.ID)

	return c.JSON(cartItem)
}

func GetCart(c *fiber.Ctx) error {

	db := config.GetDB()

	userData := c.Locals("user")
	if userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := userData.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var cartItem []models.CartBook

	db.Preload("Book").Where("user_id = ?", userID).Find(&cartItem)
	return c.JSON(cartItem)
}

func DeleteCart(c *fiber.Ctx) error {
	db := config.GetDB()

	userData := c.Locals("user")
	if userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := userData.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	bookID, err := strconv.Atoi(c.Params("book_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	var cartItem models.CartBook

	if err := db.Where("user_id = ? AND book_id = ?", userID, bookID).First(&cartItem).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": " Item nof found in Cart"})
	}

	db.Delete(&cartItem)
	return c.JSON(fiber.Map{"message": "Item removed from cart"})
}

func UpdatrCartQty(c *fiber.Ctx) error {
	db := config.GetDB()

	userData := c.Locals("user")
	if userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := userData.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	bookID, err := strconv.Atoi(c.Params("book_id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	var request struct {
		Qty int `json:"qty"`
	}
	if err := c.BodyParser(&request); err != nil || request.Qty < 1 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid quantity"})
	}

	var cartItem models.CartBook
	if err := db.Where("user_id = ? AND book_id = ?", userID, bookID).First(&cartItem).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}

	cartItem.Qty = request.Qty 
	db.Save(&cartItem)
	return c.JSON(cartItem)
}


func ClearCart(c *fiber.Ctx)error {
	db := config.GetDB()

	userData := c.Locals("user")
	if userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := userData.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))


	if err := db.Where("user_id = ?",userID).Delete(&models.CartBook{}).Error; err !=nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to clear cart"})
	}

	return c.JSON(fiber.Map{"Message": "Cart cleared"})
}

func GetCartTotal(c *fiber.Ctx)error {
	db := config.GetDB()

	userData := c.Locals("user")
	if userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	user := userData.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var total float64
	if err := db.Model(&models.CartBook{}).
	Select("SUM(books.price * cart_books.qty)").
	Joins("JOIN books ON books.id = cart_books.book_id").
	Where("cart_books.user_id = ?",userID).
	Scan(&total).Error; err !=nil {
		// return c.Status(500).JSON(fiber.Map{"error": "Failed to calculate total"})
		return err
	}
	return c.JSON(fiber.Map{"Total Price":total})
}