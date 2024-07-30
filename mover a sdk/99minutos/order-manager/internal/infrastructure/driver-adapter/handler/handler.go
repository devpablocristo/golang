package handler

import (
	"encoding/json"
	"log"
	"net/http"

	cmsapi "github.com/devpablocristo/99minutos/commons/api"
	port "github.com/devpablocristo/99minutos/order-manager/internal/application/port"
	domain "github.com/devpablocristo/99minutos/order-manager/internal/domain"
)

type Handler struct {
	orderManager port.OrderManager
	valUser      port.ValidateUser
	userRepo     port.UserRepo
	ordeRepo     port.OrderRepo
}

func NewHandler(om port.OrderManager, vu port.ValidateUser, ur port.UserRepo, or port.OrderRepo) *Handler {
	return &Handler{
		orderManager: om,
		valUser:      vu,
		userRepo:     ur,
		ordeRepo:     or,
	}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	defer body.Close()

	shippingRequest := &ShippingRequest{}
	err := json.NewDecoder(body).Decode(shippingRequest)
	if err != nil {
		responseErr := cmsapi.InvalidJSON("CreateOrder", "handler", err)
		w.WriteHeader(responseErr.StatusCode)
		err = json.NewEncoder(w).Encode(cmsapi.FailResponse(responseErr))
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(responseErr.Error())
		return
	}

	ctx := r.Context()

	role, err := h.valUser.Execute(ctx, shippingRequest.Email, shippingRequest.Password)
	if err != nil {
		responseErr := cmsapi.NewAPIError(http.StatusForbidden, "invalid-user", "CreateOrder", "handler", err)
		log.Println(responseErr)
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(responseErr.StatusCode, responseErr.Message))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		return
	}

	if role != domain.INTERNAL && role != domain.CUSTOMER {
		responseErr := cmsapi.NewAPIError(http.StatusForbidden, "invalid-user", "CreateOrder", "handler", err)
		log.Println(responseErr)
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(responseErr.StatusCode, responseErr.Message))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		return
	}

	ur, err := h.userRepo.FindByEmail(ctx, shippingRequest.Email)
	if err != nil {
		responseErr := cmsapi.NewAPIError(http.StatusForbidden, "invalid-user", "CreateOrder", "handler", err)
		log.Println(responseErr)
		w.WriteHeader(http.StatusUnauthorized)
		err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(responseErr.StatusCode, responseErr.Message))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		return
	}

	shippingRequest.CustomerID = ur.UUID
	order, err := h.orderManager.CreateOrder(ctx, shippingRequest.toOrderDomain())
	if err != nil {
		responseErr := cmsapi.NewAPIError(http.StatusForbidden, "canceled-order", "CreateOrder", "application", err)
		log.Println(responseErr)
		w.WriteHeader(responseErr.StatusCode)
		err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(responseErr.StatusCode, responseErr.Message))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	msg := order
	err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(http.StatusCreated, msg))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}
