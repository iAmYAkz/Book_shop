package handlers

import (
	"yakz/config"
	"yakz/models"

	"github.com/gofiber/fiber/v2"
)







//Gets//
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	config.GetDB().Find(&books)
	return c.JSON(books)
}

//GetById
func GetBook (c *fiber.Ctx)error {
	 db := config.GetDB()

	 id :=c.Params("id")
	 var book models.Book
	 
	 if err := db.First(&book,id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error" : "Book not found !!"})
	 }
	 return c.JSON(book)
}

//Create// 
func CreateBook(c *fiber.Ctx) error {
	db := config.GetDB()

	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		var books []models.Book
		if err := c.BodyParser(&books); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		db.Create(&books)
		return c.JSON(books)
	}
	
	db.Create(&book)
	return c.JSON(book)
}

//Update
func UpdateBook(c *fiber.Ctx)error{
	db := config.GetDB()

	id :=c.Params("id")
	 var book models.Book

	 if err := db.First(&book,id).Error; err !=nil{
		return c.Status(404).JSON(fiber.Map{"error": "Book not found !!"})
	 }

	 if err := c.BodyParser(&book); err !=nil{
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request !!"})
	 }
	 
	 db.Save(&book)
	 return c.JSON(book)
}

//Delete
func DeleteBook(c *fiber.Ctx)error{
	db := config.GetDB()

	id :=c.Params("id")
	 var book models.Book
	 if err := db.First(&book,id).Error; err !=nil{
		return c.Status(404).JSON(fiber.Map{"error": "Book not found !!"})
	 }

	 db.Delete(&book)
	 return c.JSON(fiber.Map{"Message": "Book Delete Successful"})
	 
}