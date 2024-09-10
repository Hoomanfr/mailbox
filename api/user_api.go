package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thumperq/golib/logging"
	httpserver "github.com/thumperq/golib/servers/http"
	"github.com/thumperq/wms/mailbox/internal/app"
)

type UserApi struct {
	userApp app.UserApp
}

func SetupUserApi(userApp app.UserApp, engine *http.ServeMux) {
	userApi := UserApi{
		userApp: userApp,
	}
	userApi.InitializeRoutes(engine)
}

func (api UserApi) InitializeRoutes(engine *http.ServeMux) {
	engine.HandleFunc("POST wms/mailbox/v1/users", api.createUser)
	engine.HandleFunc("GET wms/mailbox/v1/users/{id}", api.createUser)
}

func (api UserApi) createUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request app.UserRequest
	if err := httpserver.ShouldBindJson(r, &request); err != nil {
		httpserver.Json(http.StatusBadRequest, w, httpserver.H{"error": err.Error()})
		return
	}
	id, err := api.userApp.CreateUser(ctx, request)
	if err != nil {
		httpserver.Json(http.StatusInternalServerError, w, httpserver.H{"error": err.Error()})
		return
	}
	logging.TraceLogger(ctx).Info().Msgf("create user %s", id)
	httpserver.Json(http.StatusCreated, w, gin.H{"id": id})
}

func (api UserApi) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	user, err := api.userApp.FindUserById(ctx, id)
	if err != nil {
		httpserver.Json(http.StatusInternalServerError, w, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		httpserver.Status(http.StatusNotFound, w)
		return
	}
	httpserver.Json(http.StatusOK, w, user)
}
