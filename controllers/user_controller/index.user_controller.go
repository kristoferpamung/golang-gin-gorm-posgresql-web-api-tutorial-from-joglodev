package user_controller

import (
	"joglo-dev/database"
	"joglo-dev/models"
	"joglo-dev/requests"
	"joglo-dev/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {

	users := new([]models.User)

	if err := database.DB.Table("users").Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserByID(ctx *gin.Context) {

	id := ctx.Param("id")
	user := new(responses.UserResponse)

	err := database.DB.Table("users").Where("id=?", id).Find(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})
		return
	}

	if user.Name == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "data not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data transmitted.",
		"data":    user,
	})

}

func StoreUser(ctx *gin.Context) {
	var userReq requests.UserRequest

	if errReq := ctx.ShouldBindJSON(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	userEmailExist := new(models.User)

	database.DB.Table("users").Where("email=?", userReq.Email).First(&userEmailExist)

	if userEmailExist.Email != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "email already used.",
		})

		return
	}

	user := new(models.User)

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Password = &userReq.Password
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errDb := database.DB.Table("users").Create(&user).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't create data.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    user,
	})
}

func UpdateUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	userReq := new(requests.UserRequest)

	userEmailExist := new(models.User)

	if errReq := ctx.ShouldBindJSON(&userReq); errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	errDb := database.DB.Table("users").Where("id=?", id).Find(&user).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})

		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found.",
		})
		return
	}

	errUserEmailExist := database.DB.Table("users").Where("email=?", userReq.Email).Find(&userEmailExist).Error

	if errUserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})

		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email already used.",
		})
		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Password = &userReq.Password
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id=?", id).Updates(&user).Error

	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't update data.",
		})
		return
	}

	userResponse := responses.UserResponse{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		BornDate: user.BornDate,
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "data updated successfully.",
		"data":    userResponse,
	})
}

func DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	errFind := database.DB.Table("users").Where("id=?", id).Find(&user).Error

	if errFind != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
		})
		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found.",
		})
		return
	}

	errDb := database.DB.Table("users").Unscoped().Where("id=?", id).Delete(&models.User{}).Error

	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error.",
			"error":   errDb.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "data deleted successfully.",
	})
}

// Paginate
func GetUserPaginate(ctx *gin.Context) {
	page := ctx.Query("page")

	if page == "" {
		page = "1"
	}

	perPage := ctx.Query("per_page")

	if perPage == "" {
		perPage = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)

	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]models.User)

	if err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}
