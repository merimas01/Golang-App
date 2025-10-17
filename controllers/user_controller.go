package controllers

import (
	"Golang-App/models"
	"Golang-App/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
}

// @Summary      Create new user
// @Description  Adds a new user to MySQL
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.UserInsert true "User"
// @Success      201 {object} models.User
// @Router       /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var input models.UserInsert

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to create user and get the full User model
	user, err := ctrl.Service.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the saved user (with ID)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary      Update existing user
// @Description  Updates a user by ID in MySQL
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int                true  "User ID"
// @Param        user body      models.UserUpdate  true  "User Update Data"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	// Get ID from path
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Bind JSON input
	var input models.UserUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to update user
	user, err := ctrl.Service.UpdateUser(&input, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the updated user
	c.JSON(http.StatusOK, user)
}

// @Summary      Get all users
// @Tags         users
// @Produce      json
// @Success      200 {array} models.User
// @Router       /users [get]
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieve a single user by their ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")         // Get the "id" from the URL
	id, err := strconv.Atoi(idParam) // Convert string to int
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := ctrl.Service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

// GetUserByID godoc
// @Summary      Delete user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")         // Get the "id" from the URL
	id, err := strconv.Atoi(idParam) // Convert string to int
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := ctrl.Service.DeleteUser(uint(id)); err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
