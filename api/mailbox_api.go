package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thumperq/wms/mailbox/internal/application"
)

type MailboxApi struct {
	appFactory application.AppFactory
}

func SetupMailboxApi(appFactory application.AppFactory, engine *gin.Engine) {
	mailboxApi := MailboxApi{
		appFactory: appFactory,
	}
	mailboxApi.InitializeRoutes(engine)
}

func (api MailboxApi) InitializeRoutes(engine *gin.Engine) {
	engine.Group("/v1/mailboxes").
		POST("/", api.createMailbox).
		GET("/user/:user_id", api.getMailboxByUserId)
}

func (api MailboxApi) createMailbox(c *gin.Context) {
	var request application.MailboxRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := api.appFactory.MailboxApp.CreateMailbox(c, request)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (api MailboxApi) getMailboxByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	mailboxes, err := api.appFactory.MailboxApp.UserMailboxes(c, userId)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if mailboxes == nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, mailboxes)
}
