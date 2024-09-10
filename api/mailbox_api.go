package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	httpserver "github.com/thumperq/golib/servers/http"
	"github.com/thumperq/wms/mailbox/internal/app"
)

type MailboxApi struct {
	mailboxApp app.MailboxApp
}

func SetupMailboxApi(mailboxApp app.MailboxApp, engine *http.ServeMux) {
	mailboxApi := MailboxApi{
		mailboxApp: mailboxApp,
	}
	mailboxApi.InitializeRoutes(engine)
}

func (api MailboxApi) InitializeRoutes(engine *http.ServeMux) {
	engine.HandleFunc("POST wms/mailbox/v1/mailboxes", api.createMailbox)
	engine.HandleFunc("GET wms/mailbox/v1/user/{user_id}", api.getMailboxByUserId)
}

func (api MailboxApi) createMailbox(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request app.MailboxRequest
	if err := httpserver.ShouldBindJson(r, &request); err != nil {
		httpserver.Json(http.StatusBadRequest, w, httpserver.H{"error": err.Error()})
		return
	}
	id, err := api.mailboxApp.CreateMailbox(ctx, request)
	if err != nil {
		httpserver.Json(http.StatusInternalServerError, w, httpserver.H{"error": err.Error()})
		return
	}
	httpserver.Json(http.StatusCreated, w, gin.H{"id": id})
}

func (api MailboxApi) getMailboxByUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := r.PathValue("user_id")
	mailboxes, err := api.mailboxApp.UserMailboxes(ctx, userId)
	if err != nil {
		httpserver.Json(http.StatusInternalServerError, w, gin.H{"error": err.Error()})
		return
	}
	if mailboxes == nil {
		httpserver.Status(http.StatusNotFound, w)
		return
	}
	httpserver.Json(http.StatusOK, w, mailboxes)
}
