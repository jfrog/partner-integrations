package server

import (
	"encoding/json"
	"net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Good struct {
	Id string `json:"id"`
	Status string `json:"status"`
}

type Bad struct {
	FailureId string `json:"failureId"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type Error struct {
	Error string `json:"error"`
}

// GoodHandler godoc
// @Summary Example successful endpoint
// @Description Example successful endpoint
// @Produce json
// @Security ApiKeyAuth
// @Security OAuth2Application[write, admin]
// @Success 200 {object} Good
// @Failure 500 {object} Error
// @Router /good [get]
func GoodHandler(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpan("good.handler.request", tracer.ResourceName("/good"))
	defer span.Finish()
	good := Good{Id: "1", Status: "success"}
	body, err := json.Marshal(good)
	if err != nil {
		requestNotSupported(w,  err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
	return
}

// BadHandler godoc
// @Summary Example failure endpoint
// @Description Example failure endpoint
// @Produce json
// @Security ApiKeyAuth
// @Security OAuth2Application[write, admin]
// @Success 200 {object} Bad
// @Failure 500 {object} Error
// @Router /bad [get]
func BadHandler(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpan("bad.handler.request", tracer.ResourceName("/bad"))
	defer span.Finish()
	errJson := Error{Error: "HTTP 500 Error!"}
	errBody, _ := json.Marshal(errJson)
	errString := string(errBody[:])
	http.Error(w, errString, http.StatusInternalServerError)
	return
}


func requestNotSupported(w http.ResponseWriter, err string, httpCode int) {
	errJson := Error{Error: err}
	errBody, _ := json.Marshal(errJson)
	errString := string(errBody[:])
	http.Error(w, errString, httpCode)
}
