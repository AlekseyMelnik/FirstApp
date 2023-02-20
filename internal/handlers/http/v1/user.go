package v1

import (
	"FirstProject/internal/domain/entity"
	"FirstProject/internal/domain/service"
	"FirstProject/pkg/logging"

	handlers "FirstProject/internal/handlers/http"
	"FirstProject/internal/handlers/http/dto"
	"encoding/json"
	"net/http"

	//"context"
	"github.com/julienschmidt/httprouter"
	//"net/http"
)

const (
	userUrl = "/user"
)

type Service interface {
	CreateUser(user *service.CreateUserDTO) (*entity.User, error)
}

type userHandler struct {
	userService Service
}

func NewUserHandler(userService Service) handlers.Handler {
	return &userHandler{userService: userService}
}
func (h *userHandler) Register(router *httprouter.Router) {
	router.POST(userUrl, h.CreateUser)
}
func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	logger := logging.GetLogger()
	logger.Info("Create user in Handler")
	var d dto.CreateUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		logger.Fatal(err)
	}
	//MAPPING dto.CreateBookDTO --> book_usecase.CreateBookDTO
	seviceUserDto := service.CreateUserDTO{
		Email:    d.Email,
		Password: d.Password,
	}
	user, err := h.userService.CreateUser(&seviceUserDto)
	if err != nil {
		logger.Fatal(err)
		// TODO JSON RPC: TRANSPORT: 200, error: {msg, ..., dev_msg}
	}
	w.WriteHeader(http.StatusOK)
}
