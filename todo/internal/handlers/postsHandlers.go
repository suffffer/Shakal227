package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suffffer/Shakal227/internal/database"
	"strconv"
)

func CreatePost(c *fiber.Ctx) error {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	var newPost struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&newPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	post := db.Todo{
		Title:       newPost.Title,
		Description: newPost.Description,
		UserID:      uint(userID),
	}

	if err := db.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	return c.JSON(post)
}

func GetPosts(c *fiber.Ctx) error {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	var posts []db.Todo
	if err := db.DB.Where("user_id = ?", uint(userID)).Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch posts"})
	}

	return c.JSON(posts)
}

func UpdatePost(c *fiber.Ctx) error {
	userIDStr := c.Query("user_id")
	postIDStr := c.Params("id")

	if userIDStr == "" || postIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id and post_id are required"})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post_id"})
	}

	var updatedPost struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&updatedPost); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	var post db.Todo
	if err := db.DB.First(&post, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	if post.UserID != uint(userID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not authorized to update this post"})
	}

	post.Title = updatedPost.Title
	post.Description = updatedPost.Description
	if err := db.DB.Save(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update post"})
	}

	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	userIDStr := c.Query("user_id")
	postIDStr := c.Params("id")

	if userIDStr == "" || postIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id and post_id are required"})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user_id"})
	}

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post_id"})
	}

	var post db.Todo
	if err := db.DB.First(&post, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}

	if post.UserID != uint(userID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not authorized to delete this post"})
	}

	if err := db.DB.Delete(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	return c.JSON(fiber.Map{"message": "Post deleted"})
}

func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post ID",
		})
	}

	var post db.Todo
	if err := db.DB.First(&post, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
