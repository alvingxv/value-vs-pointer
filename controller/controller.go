package controller

import (
	"fmt"
	"go-mem-profile/service"
	"net/http"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) handler {
	return handler{
		service: service,
	}
}

func (h *handler) GetDataPointer(w http.ResponseWriter, r *http.Request) {
	res := h.service.GetDataPointer()
	// a, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Fprintf(w, string(a))
	fmt.Fprintf(w, res)
}

func (h *handler) GetDataValue(w http.ResponseWriter, r *http.Request) {
	res := h.service.GetDataValue()

	// a, err := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Fprintf(w, string(a))
	fmt.Fprintf(w, res)
}
