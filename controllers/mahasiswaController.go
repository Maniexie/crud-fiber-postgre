package controllers

import (
	"log"

	"github.com/Maniexie/crud-fiber-postgres/database"
	"github.com/Maniexie/crud-fiber-postgres/models/entities"
	"github.com/Maniexie/crud-fiber-postgres/models/req"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// func MahasiswaControllerIndex(c *fiber.Ctx) error {
// 	return c.JSON(fiber.Map{
// 		"message": "Selamat datang",
// 	})
// }

func MahasiswaControllerShow(c *fiber.Ctx) error {
	var mahasiswa []entities.Mahasiswa
	err := database.DB.Find(&mahasiswa).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(mahasiswa)
}

func MahasiswaControllerCreate(c *fiber.Ctx) error {
	mahasiswa := new(req.MahasiswaReq)
	if err := c.BodyParser(mahasiswa); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(mahasiswa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama / Email tidak sesuai",
			"error":   err.Error(),
			"test":    "saya ingin ngetes kk",
		})
	}

	// result := database.DB.Where("nim = ?", mahasiswa.Nim).First(&mahasiswa)
	// if result.RowsAffected > 0 {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Nim sudah ada"})
	// }

	newMahasiswa := entities.Mahasiswa{
		Nim:   mahasiswa.Nim,
		Name:  mahasiswa.Name,
		Email: mahasiswa.Email,
	}

	if err := database.DB.Create(&newMahasiswa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan data baru",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil Membuat Data Baru",
		"data":    newMahasiswa,
	})

}

// Mencari data berdasarkan id
func MahasiswaControllerFind(c *fiber.Ctx) error {
	var mahasiswa []entities.Mahasiswa
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak Boleh Kosong",
		})
		return nil
	}

	if err := database.DB.Where("id = ?", id).First(&mahasiswa).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data mahasiswa tidak di temukan",
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data mahasiswa",
		"data":    mahasiswa,
	})
}

// update data
func MahasiswaControllerUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	mahasiswa := new(entities.Mahasiswa)

	if err := c.BodyParser(mahasiswa); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Where("id = ?", id).Updates(&mahasiswa)
	return c.JSON(mahasiswa)

}

func MahasiswaControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var mahasiswa []entities.Mahasiswa

	result := database.DB.Delete(&mahasiswa, id)

	if result.RowsAffected == 0 {
		return c.JSON(map[string]string{
			"message": "Data mahasiswa tidak ada , Silahkan cek kembali",
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "Data berhasil di hapus",
	})
}
