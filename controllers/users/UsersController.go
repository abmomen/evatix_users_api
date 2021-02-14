package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abmomin/users_api/domain/users"
	"github.com/abmomin/users_api/services"
	"github.com/abmomin/users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	fmt.Println(userID)
	if userErr != nil {
		return 0, errors.NewBadRequestError(fmt.Sprintf("user id: %d should be a number", userID))
	}
	return userID, nil
}

//CreateUser creates a new user in the database.
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

//GetUser returns a user with asked details
func Get(c *gin.Context) {
	userID, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context)  {
	userID, idErr := getUserId("userId")
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

	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch
	result,  err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context)  {
	userID, idErr := getUserId("userId")
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func Search(c *gin.Context)  {
	status := c.Query("status")
	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, users)
}

//SearchUser finds a user in database.
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
