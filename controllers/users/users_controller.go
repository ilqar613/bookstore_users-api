package users

import (
	//"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilqar613/bookstore_users-api/domain/users"
	"github.com/ilqar613/bookstore_users-api/services"
	"github.com/ilqar613/bookstore_users-api/utils/errors"
)

func GetUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//handle json error
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public")=="true"))
}

func Get(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public")=="true"))

}

func Update(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public")=="true"))
}

func Delete(c *gin.Context) {
	userId, idErr := GetUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	err := services.UsersService.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status:=c.Query("status")
	users,err:=services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
	}
	
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public")=="true"))
}
