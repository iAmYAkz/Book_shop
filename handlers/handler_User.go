package handlers

import (
	"yakz/config"
	"yakz/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func DeleteUser(c *fiber.Ctx) error {
	db := config.GetDB()
	id := c.Params("id")

	if err := db.Delete(&models.User{},id).Error ; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"Message": "User deleted successfuly"})
}

func UpdateUserRole(c *fiber.Ctx)error {
	db := config.GetDB()
	id := c.Params("id")

	var input struct{
		Role string `json:"role"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	if err := db.First(&user,id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error":"User not found"})
	}

	user.Role = input.Role
	db.Save(&user)

	return c.JSON(fiber.Map{"Message": "User role update","user":user})
}

func ChangeUserPassword(c *fiber.Ctx)error {
	db := config.GetDB()
	id := c.Params("id")

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BodyParser(&input); err !=nil{
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	if err := db.First(&user,id).Error; err !=nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}


	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.OldPassword)); err !=nil{
		return c.Status(401).JSON(fiber.Map{"error": "Old password incorrect"})
	}

	hsahedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.NewPassword),bcrypt.DefaultCost)
	user.Password = string(hsahedPassword)
	db.Save(&user)

	return c.JSON(fiber.Map{"message": "Password updated successfully"})
}

func UpdateProfile(c * fiber.Ctx)error {
	db := config.GetDB()
	id := c.Params("id")

	var input struct{
		Name string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User 

	if err := db.First(&user ,id).Error ; err !=nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	user.Name = input.Name
	user.Email = input.Email

	db.Save(&user)

	return c.JSON(fiber.Map{"Message":"Profile update","user": user})
}