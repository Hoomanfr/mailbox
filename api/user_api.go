package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thumperq/wms/mailbox/internal/application"
)

type UserApi struct {
	appFactory application.AppFactory
}

func SetupUserApi(appFactory application.AppFactory, engine *gin.Engine) {
	userApi := UserApi{
		appFactory: appFactory,
	}
	userApi.InitializeRoutes(engine)
}

func (api UserApi) InitializeRoutes(engine *gin.Engine) {
	engine.Group("/v1/users").
		POST("/", api.createUser).
		GET("/:id", api.getUser)
}

func (api UserApi) createUser(c *gin.Context) {
	var request application.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := api.appFactory.UserApp.CreateUser(c, request)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api UserApi) getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := api.appFactory.UserApp.FindUserById(c, id)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}
