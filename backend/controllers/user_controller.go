package controllers

import (
	"Golang-App/interfaces"
	"Golang-App/models"
	searchobjects "Golang-App/models/search_objects"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController[models.User, models.UserInsert, models.UserUpdate, searchobjects.BaseSearchObject]
}

// NewUserController constructor
func NewUserController(service interfaces.IUserService) *UserController {
	return &UserController{
		BaseController: BaseController[models.User, models.UserInsert, models.UserUpdate, searchobjects.BaseSearchObject]{
			Service: service,
		},
	}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Adds a new user to the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.UserInsert true "User data"
// @Success      201 {object} models.User
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	ctrl.BaseController.Create(c)
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Returns a paginated list of users
// @Tags         users
// @Produce      json
// @Param        page     query    int    false  "Page number"
// @Param        page_size query    int    false  "Number of items per page (page size)"
// @Success      200 {object} models.PagedResult[models.User]
// @Failure      500 {object} map[string]string
// @Router       /users [get]
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	ctrl.BaseController.GetAll(c)
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Returns a user by its ID
// @Tags         users
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} models.User
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /users/{id} [get]
func (ctrl *UserController) GetUserByID(c *gin.Context) {
	ctrl.BaseController.GetByID(c)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Updates a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        user body models.UserUpdate true "Updated user data"
// @Success      200 {object} models.User
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	ctrl.BaseController.Update(c)
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID
// @Tags         users
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	ctrl.BaseController.Delete(c)
}
