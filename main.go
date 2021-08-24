package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       int
}

type Book struct {
	Id      int
	Author  string
	Name    string
	Volumes string
	Chapter string
}

var menus = []Menu{
	{
		Id:          1,
		Name:        "BB Corn",
		Description: "Giant breed of rare corn that was eaten by Gourmet Nobility as a snack long ago",
		Price:       4000,
	},
	{
		Id:          2,
		Name:        "Century Soup",
		Description: "A soup cooked with hundreds or even thousands of ingredients",
		Price:       10000,
	},
	{
		Id:          3,
		Name:        "Jewel Meat",
		Description: "Incandescent lamp-like radiance that dulls jewels and lights up a night sky",
		Price:       8000,
	},
}

var books = []Book{
	{Id: 1,
		Author:  "Masashi Kishimoto",
		Name:    "Naruto",
		Volumes: "72",
		Chapter: "700",
	},
	{Id: 2,
		Author:  "Tite Kubo",
		Name:    "Bleach",
		Volumes: "63",
		Chapter: "686",
	},
	{Id: 3,
		Author:  "Mitsutoshi Shimabukuro",
		Name:    "Toriko",
		Volumes: "42",
		Chapter: "396",
	},
	{Id: 4,
		Author:  "Eiichiro Oda",
		Name:    "One Piece",
		Volumes: "99",
		Chapter: "1022",
	},
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/404.html")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/hello.html")
}

func AllMenusHandler(w http.ResponseWriter, r *http.Request) {
	menuTemplate, err := template.ParseFiles("views/menus/index.html")
	fmt.Printf("%v\n", menuTemplate)
	fmt.Printf("%v\n", err)
	if err != nil {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	menuTemplate.Execute(w, menus)
}

func MenusHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	menuIndex, _ := strconv.ParseInt(params["id"], 0, 64)
	menuIndex -= 1

	menuTemplate, err := template.ParseFiles("views/menus/show.html")
	if err != nil || isOutOfRange(menuIndex) {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	fmt.Print("test")
	menuTemplate.Execute(w, menus[menuIndex])
}

func AllBooksHandler(w http.ResponseWriter, r *http.Request) {
	listTemplate, err := template.ParseFiles("views/books/index.html")
	if err != nil {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	listTemplate.Execute(w, books)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookIndex, _ := strconv.ParseInt(params["id"], 0, 64)
	bookIndex -= 1

	listTemplate, err := template.ParseFiles("views/books/show.html")
	if err != nil || isOutOfBookRange(bookIndex) {
		http.ServeFile(w, r, "public/500.html")
		return
	}
	fmt.Print("test")
	listTemplate.Execute(w, books[bookIndex])
}

func isOutOfRange(index int64) bool {
	return (index < 0 || index >= int64(len(menus)))
}

func isOutOfBookRange(index int64) bool {
	return (index < 0 || index >= int64(len(books)))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/home", HomeHandler)
	router.HandleFunc("/menus", AllMenusHandler)
	router.HandleFunc("/menus/{id:[0-9]+}", MenusHandler)
	// Please help implement one of following routing
	// router.HandleFunc("/products", AllProductsHandler)
	// router.HandleFunc("/products/{id:[0-9]+}", ProductsHandler)
	// or
	router.HandleFunc("/books", AllBooksHandler)
	router.HandleFunc("/books/{id:[0-9]+}", BooksHandler)
	// or
	// router.HandleFunc("/songs", AllSongsHandler)
	// router.HandleFunc("/songs/{id:[0-9]+}", SongsHandler)
	router.NotFoundHandler = http.HandlerFunc(NotFound)

	http.Handle("/", router)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("public/css"))))
	http.ListenAndServe(":8000", nil)
}
