package inventory

import (
	"strconv"
	"tes1/model"
	"tes1/varglobal"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
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

	var inventories []model.Inventory
	query := varglobal.DB.Model(&model.Inventory{})

	// Filter by search keyword
	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("name LIKE ?", searchTerm)
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

func PostInventory(c *gin.Context) {
	var tambah model.Inventory

	if err := c.ShouldBindJSON(&tambah); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := varglobal.DB.Create(&tambah).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to retrieve item",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "item added successfully",
		"book":    tambah,
	})

}

func UpdateInventory(c *gin.Context) {
	var updateData model.Inventory

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var existingInventaries model.Inventory

	if err := varglobal.DB.First(&existingInventaries, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	// Lakukan update
	if err := varglobal.DB.Model(&existingInventaries).Updates(updateData).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Book updated successfully",
		"book":    existingInventaries,
	})
}

func DeleteInventory(c *gin.Context) {
	var DeleteData model.Inventory

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var existingInventaries model.Inventory
	if err := varglobal.DB.First(&existingInventaries, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}

	// Lakukan update
	if err := varglobal.DB.Model(&existingInventaries).Where("id = ?", idParam).Delete(DeleteData).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Book deleted successfully",
		"book":    existingInventaries,
	})
}

func GetInventoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var inventory model.Inventory
	if err := varglobal.DB.First(&inventory, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(200, gin.H{
		"book": inventory,
	})
}
