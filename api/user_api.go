package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thumperq/golib/logging"
	"github.com/thumperq/wms/mailbox/internal/app"
)

type UserApi struct {
	userApp app.UserApp
}

func SetupUserApi(userApp app.UserApp, engine *gin.Engine) {
	userApi := UserApi{
		userApp: userApp,
	}
	userApi.InitializeRoutes(engine)
}

func (api UserApi) InitializeRoutes(engine *gin.Engine) {
	engine.Group("wms/mailbox/v1/users").
		POST("/", api.createUser).
		GET("/:id", api.getUser)
}

func (api UserApi) createUser(c *gin.Context) {
	var request app.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := api.userApp.CreateUser(c, request)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logging.TraceLogger(c.Request.Context()).Info().Msgf("create user %s", id)
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api UserApi) getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := api.userApp.FindUserById(c, id)
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
