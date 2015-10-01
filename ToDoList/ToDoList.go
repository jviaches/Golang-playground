package main


import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)


type ToDoItem struct  {
	Id int
	Caption string
	IsFinished bool
}

var TodoItemsSlice =  []ToDoItem{
}

func main() {
	fmt.Println("Initializing...")
	TodoItemsSlice = make([]ToDoItem,0)
	runServer()
}

func runServer(){
	fmt.Println("Starting sever...")

	// public views
	http.HandleFunc("/additem", addItemHandler )

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	captionFormValue := r.PostFormValue("caption")
	if  captionFormValue!= "" {

		fmt.Println("caption:", captionFormValue)

		newId := len(TodoItemsSlice) + 1
		p := &ToDoItem{Id:newId, Caption: captionFormValue, IsFinished:false}
		TodoItemsSlice = append(TodoItemsSlice, *p)
	}

	t, err := template.ParseFiles("views/main.gtpl")
	if err != nil {
		log.Fatal("Can not parse views/main.gtpl "+ err.Error())
	}
	t.Execute(w, TodoItemsSlice)
}


func addItemHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("addItemHandler")
	t, _ := template.ParseFiles("views/additem.gtpl")
	t.Execute(w, t)
}