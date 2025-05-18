package snippetControllers

import (
	"strconv"
	"vault-dev/config"
	"vault-dev/models"

	"github.com/gin-gonic/gin"
)

func PostSnippet(c *gin.Context) {
	var snippet models.SnippetModel

	if err := c.ShouldBindJSON(&snippet); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&snippet).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create snippet"})
		return
	}
	c.JSON(201, snippet)
}

func GetSnippets(c *gin.Context) {
	var snippets []models.SnippetModel

	// Build base query
	query := config.DB.Model(&models.SnippetModel{})

	// Pagination: parse page and limit
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Filters
	if language := c.Query("lang"); language != "" {
		query = query.Where("language ILIKE ?", "%"+language+"%")
	}

	if username := c.Query("author"); username != "" {
		query = query.Where("author ILIKE ?", "%"+username+"%")
	}

	if title := c.Query("title"); title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	if tags := c.Query("tags"); tags != "" {
		query = query.Where("tags ILIKE ?", "%"+tags+"%")
	}

	if folders := c.Query("folders"); folders != "" {
		query = query.Where("folders ILIKE ?", "%"+folders+"%")
	}
	if user_id := c.Query("user_id"); user_id != "" {
		query = query.Where("user_id = ?", user_id)
	}

	// Apply pagination and execute query
	if err := query.Limit(limit).Offset(offset).Find(&snippets).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch paginated snippets"})
		return
	}

	c.JSON(200, gin.H{
		"snippets": snippets,
		"page":     page,
		"limit":    limit,
		"total":    len(snippets),
	})
}

func GetFavSnippets(c *gin.Context) {
	var snippets []models.SnippetModel

	// Pagination setup
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Build query
	query := config.DB.Model(&models.SnippetModel{}).Where("favorite = ?", true)

	// Optional user_id filter
	if userIDStr := c.Query("user_id"); userIDStr != "" {
		if userID, err := strconv.Atoi(userIDStr); err == nil {
			query = query.Where("user_id = ?", userID)
		} else {
			c.JSON(400, gin.H{"error": "Invalid user_id"})
			return
		}
	}

	// Execute query with pagination
	if err := query.Limit(limit).Offset(offset).Find(&snippets).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch favorite snippets"})
		return
	}

	c.JSON(200, gin.H{
		"page":     page,
		"limit":    limit,
		"snippets": snippets,
	})
}
func UpdateSnippet(c *gin.Context) {
	var snippet models.SnippetModel
	user_id := c.Query("user_id")
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	if user_id == "" {
		c.JSON(400, gin.H{"error": "user_id is required"})
		return
	}
	if err := config.DB.First(&snippet, id, user_id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Snippet not found"})
		return
	}

	var updatedSnippet models.SnippetModel

	if err := c.ShouldBindJSON(&updatedSnippet); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	snippet.Title = updatedSnippet.Title
	snippet.Code = updatedSnippet.Code
	snippet.Language = updatedSnippet.Language
	snippet.Description = updatedSnippet.Description
	snippet.Tags = updatedSnippet.Tags
	snippet.Folders = updatedSnippet.Folders
	snippet.Username = updatedSnippet.Username
	snippet.Likes = updatedSnippet.Likes
	snippet.Favorite = updatedSnippet.Favorite

	if err := config.DB.Save(&snippet).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update snippet"})
		return
	}
	c.JSON(200, snippet)
}

func DeleteSnippet(c *gin.Context) {
	id := c.Param("id")
	var snippet models.SnippetModel

	if err := config.DB.First(&snippet, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Snippet not found"})
		return
	}

	if err := config.DB.Delete(&snippet).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete snippet"})
		return
	}
	c.JSON(200, gin.H{"message": "Snippet deleted successfully"})
}
