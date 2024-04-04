package main

import (
	"go-mem-profile/controller"
	"go-mem-profile/repo"
	"go-mem-profile/service"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// database.InitiliazeDatabase()

	// db := database.GetDatabaseInstance()

	repository := repo.NewRepository(nil)
	service := service.NewService(repository)
	handler := controller.NewHandler(service)

	// Add profiling endpoints
	// http.HandleFunc("/debug/pprof/", pprof.Index)
	// http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Add your application endpoints
	http.HandleFunc("/pointer", handler.GetDataPointer)
	http.HandleFunc("/value", handler.GetDataValue)

	http.ListenAndServe(":8080", nil)
}
