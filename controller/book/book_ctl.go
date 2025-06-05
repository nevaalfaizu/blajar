package book

import (
	"strconv"
	"tes1/model"
	"tes1/varglobal"

	"github.com/gin-gonic/gin"
)

// var books = []model.Book{
// 	{ID: 1, Title: "Book One", Author: "Author One", Year: 2021},
// 	{ID: 2, Title: "Book Two", Author: "Author Two", Year: 2022},
// 	{ID: 3, Title: "Book Three", Author: "Author Three", Year: 2023},
// }

func GetBooks(c *gin.Context) {
	// Query parameter
	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "10")
	search := c.Query("search")
	yearParam := c.Query("year")

	page, _ := strconv.Atoi(pageParam)
	limit, _ := strconv.Atoi(limitParam)

	// Validasi agar tidak nol
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// Filter data sesuai search dan tahun (jika ada)
	var inventories []model.Book
	query := varglobal.DB.Model(&model.Book{})

	// Filter by search keyword
	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("title LIKE ?", searchTerm)
	}

	// Filter by year (jika ada)
	if yearParam != "" {
		query = query.Where("year = ?", yearParam)
	}

	// Hitung total
	var total int64
	query.Count(&total)

	// Pagination
	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&inventories).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve inventory"})
		return
	}

	// Response
	c.JSON(200, gin.H{
		"page":   page,
		"limit":  limit,
		"total":  total,
		"items":  inventories,
		"search": search,
		"year":   yearParam,
	})
}

func PostBook(c *gin.Context) {
	var tambah model.Book

	// Bind JSON dari request ke struct Book
	if err := c.ShouldBindJSON(&tambah); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := varglobal.DB.Create(&tambah).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to retrieve books",
		})
		return
	}

	// tambah.ID = len(tambah) + 1

	// tambah = append(tambah, tambah)

	c.JSON(201, gin.H{
		"message": "Book added successfully",
		"book":    tambah,
	})
}

func UpdateBook(c *gin.Context) {
	var updateData model.Book

	// Ambil dan validasi ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	// Bind data JSON ke struct
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Cari data buku yang akan diupdate
	var existingBook model.Book
	if err := varglobal.DB.First(&existingBook, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	// Lakukan update
	if err := varglobal.DB.Model(&existingBook).Updates(updateData).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Book updated successfully",
		"book":    existingBook,
	})
}

func DeleteBook(c *gin.Context) {
	var DeleteData model.Book

	// Ambil dan validasi ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	// Cari data buku yang akan diupdate
	var existingBook model.Book
	if err := varglobal.DB.First(&existingBook, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	// Lakukan update
	if err := varglobal.DB.Model(&existingBook).Where("id = ?", idParam).Delete(DeleteData).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Book deleted successfully",
		"book":    existingBook,
	})

}

func GetBookByID(c *gin.Context) {
	// Ambil ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	var book model.Book
	if err := varglobal.DB.First(&book, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(200, gin.H{
		"book": book,
	})
}
