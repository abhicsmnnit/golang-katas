package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type Logger interface {
	Log(msg string)
}

type LoggerAdapter func(string)

func (lg LoggerAdapter) Log(msg string) {
	lg(msg)
}

type DataStore interface {
	UserNameForId(userId string) (string, bool)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Abhinav",
			"2": "Siddhanta",
			"3": "Ashraya",
		},
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type SimpleLogic struct {
	logger    Logger
	dataStore DataStore
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		logger:    l,
		dataStore: ds,
	}
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.logger.Log("Saying hello to " + userId)
	name, ok := sl.dataStore.UserNameForId(userId)
	if !ok {
		return "", errors.New("user not found with id " + userId)
	}
	return "Hello, " + name, nil
}

type Controller struct {
	logger Logger
	logic  Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("Saying hello")

	userID := r.URL.Query().Get("user-id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user-id is a required param"))

		return
	}

	msg, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))

		return
	}

	w.Write([]byte(msg))
}

func NewController(logger Logger, logic Logic) Controller {
	return Controller{
		logger: logger,
		logic:  logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
