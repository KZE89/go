package main

import (
	"encoding/binary"
	"encoding/json"
	"log"
    "net/http"
	"reflect"
	"fmt"
	
	"model"
	"github.com/gorilla/mux"
)

var (
	Auth map[string]string = make(map[string]string)
	Work map[string]int = make(map[string]int)
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", mainPage)
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/login/pass", changePass).Methods("POST")
	r.HandleFunc("/login/dowork", doWork).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func init() {
	Work["admin"] = 1000000
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta name="theme-color" content="#375EAB">
		
			<title>main page</title>
		</head>
		<body>
			Page body and some more content
		</body>
		</html>`))
}

func login(w http.ResponseWriter, r *http.Request) {

	login := r.FormValue("login")
	pass := r.FormValue("pass")

	if Auth[login] == pass {
		w.WriteHeader(http.StatusOK)
	}
	
	user := new(model.User)
	err := model.Get(user, login, pass)

	if err == nil {
		Auth[login] = pass
		Work[login] = user.Worknumber
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	model.GormClose()
}

func changePass(w http.ResponseWriter, r *http.Request) {
	
	login := r.FormValue("login")
	pass := r.FormValue("pass")

	newPass := r.FormValue("newPass")

	if Auth[login] != pass {
		w.WriteHeader(http.StatusBadRequest)
	}

	user := new(model.User)
	err := model.Get(user, login, pass)
	
	if err == nil {
		err := model.Save(user, newPass)
		if err == nil {
			Auth[login] = newPass
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}


type DTO struct {
	BigNumber 	  int64
	Text      string
}


func doWork(w http.ResponseWriter, r *http.Request) {
	var value DTO
	login := r.FormValue("login")
	if Work[login] <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal([]byte(r.FormValue("value")), &value)
	
	v := reflect.ValueOf(&value).Elem()
	
	for i := 0; i < v.NumField(); i++ {
		w.Write(reverse(v.Field(i)))
	}
}

func reverse(val reflect.Value) []byte {
	switch val.Kind().String() {
		case "int64":
			//Для отладки
			fmt.Println("int64 = %d", uint64(val.Interface().(int64)))
			fmt.Println("int64 = %d", uint64(9223372036854775807 - val.Interface().(int64)))
			result := make([]byte, 8)
			binary.LittleEndian.PutUint64(result, uint64(9223372036854775807 - val.Interface().(int64)))
			return result
		case "int32":
			//Для отладки
			fmt.Println("int32 = %d", uint32(val.Interface().(int32)))
			fmt.Println("int32 = %d", uint32(2147483647 - val.Interface().(int32)))
			result := make([]byte, 4)
			binary.LittleEndian.PutUint32(result, uint32(2147483647 - val.Interface().(int32)))
			return result
		case "string":
			var result string
			for i := len(val.Interface().(string))-1; i >= 0; i-- {
				result += string(val.Interface().(string)[i])
			}
			//Для отладки
			fmt.Println("string = %s",string(val.Interface().(string)))
			fmt.Println("string = %s",result)
			return []byte(result)
	}
	return nil
}
