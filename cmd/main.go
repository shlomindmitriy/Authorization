package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Authorization/internal/database/inmemory"
	"github.com/Authorization/internal/models"
	"github.com/Authorization/internal/service"
)

func main() {
	// Поднять http server
	// создать две эндпоинта
	// 1) v1/auth/login    GET
	// 2) v1/auth/register POST
	db := inmemory.NewInMemory(make(map[string]models.UserInfo))
	svc := service.NewService(db)
	ha := handler{
		svc: svc,
	}

	http.HandleFunc("/v1/auth/login", ha.login)
	http.HandleFunc("/v1/auth/register", ha.register)
	http.HandleFunc("/v1/auth/delete", ha.delete)

	// http.HandleFunc("/v1/user/info", nil)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type handler struct {
	svc *service.Service
}

func (h *handler) login(w http.ResponseWriter, r *http.Request) {
	logname := r.URL.Query().Get("logname")
	password := r.URL.Query().Get("password")

	a, err := h.svc.Login(context.Background(), logname, password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	buf, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	w.Write(buf)
	w.WriteHeader(http.StatusOK)

}
func (h *handler) register(w http.ResponseWriter, r *http.Request) {

	// h.svc.Register()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	req := models.UserInfo{}

	err = json.Unmarshal(b, &req)
	if err != nil {
		panic(err)
	}
	fmt.Println(req)
	err = h.svc.Register(context.Background(), req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	logname := r.URL.Query().Get("logname")

	h.svc.Delete(context.Background(), logname)

}
