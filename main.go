package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
	)

//declaration for todolist
type TodoList struct{
	ID string `json:"id,omitEmpty"`
	Todo string `json:"todo,omitEmpty"`
}

//array of todolist
var todoList []TodoList


//index route
func Index(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("TodoList API using GO"))
}

//getAll route : getAll todolist
func GetAll(w http.ResponseWriter,req *http.Request){
	json.NewEncoder(w).Encode(todoList)
}


//createOne route : createone todolist
func Createone(w http.ResponseWriter,req *http.Request){
	params := mux.Vars(req)
	var todo TodoList
	_ = json.NewDecoder(req.Body).Decode(&todo)
	todo.ID = params["id"]
	todoList = append(todoList,todo)
	json.NewEncoder(w).Encode(todoList)
}

//deleteOne route: deleteone todolist
func Deleteone(w http.ResponseWriter,req *http.Request){
	params := mux.Vars(req)
	for index,item := range todoList {
		if item.ID == params["id"] {
			todoList = append(todoList[:index],todoList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todoList)
}


func main(){
	router := mux.NewRouter()
	todoList = append(todoList,TodoList{ID:"1",Todo:"Test data1"})
	todoList = append(todoList,TodoList{ID:"2",Todo:"Test data2"})
	router.HandleFunc("/",Index).Methods("GET");
	router.HandleFunc("/getAll",GetAll).Methods("GET")
	router.HandleFunc("/createOne/{id}",Createone).Methods("POST")
	router.HandleFunc("/deleteOne/{id}",Deleteone).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8888",router))
}