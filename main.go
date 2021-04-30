package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type MyUrl struct {

	// ID of user
	ID int32 `json:"id,omitempty" gorm:"primary_key"`

	// user login
	LongUrl string `json:"longUrl"`

	// HASH of user password. Use SHA2 algorithm
	ShortUrl string `json:"shortUrl,omitempty"`

}



var Urls []MyUrl
type Articles []Article
var articles Articles
var db *gorm.DB
var templates *template.Template

func dbinit() {
	fmt.Println("Connecting to databse ...")
	dsn := "root@tcp(127.0.0.1:3306)/url?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Print("Connection to dababase successful!")
	}
}

func homepage(w http.ResponseWriter, r *http.Request)  {
	//t,_:=template.ParseFiles("Index.html")
	//ExecuteTemplate(w,t)

	templates.ExecuteTemplate(w, "Index.html", nil)
}

func handleRequests(){

	myRouter:=mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/homepage", homepage)
	myRouter.HandleFunc("/create", createEndpoint).Methods("POST")
	myRouter.HandleFunc("/{id}", RedirectToRoot).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",myRouter))
}

func RedirectToRoot(writer http.ResponseWriter, request *http.Request) {
params := mux.Vars(request)
	var url MyUrl
	log.Println("finding original url...")
	s := params["id"]
result := db.Table("url").Find(&url,s)
	if result.Error != nil{
		writer.WriteHeader(500)
	}
	log.Println("redirecting...")
http.Redirect(writer,request,url.LongUrl,301)
}

func createEndpoint(w http.ResponseWriter, request *http.Request) {

	var url MyUrl
	var url2 MyUrl

	_ = json.NewDecoder(request.Body).Decode(&url)
	//var n1qlParams []interface{}
	//n1qlParams = append(n1qlParams,url.LongUrl)
	url.LongUrl=request.FormValue("url")
	query := db.Table("url").Find(&url,"LongUrl=?")
	if query.RowsAffected != 0{
		log.Print("url found in database!")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(url)
	}else{
		log.Print("creating new url...")
		//w.WriteHeader(404)
		var ID int32
		query :=db.Table("url").Last(&url2)
		if query.Error != nil{
			w.WriteHeader(500)
		}
		ID=url2.ID+1
		url.ShortUrl= "http://localhost:8080/"+fmt.Sprint(ID)
		result := db.Table("url").Create(&url)
		if result.Error != nil{
			w.WriteHeader(500)
		}else{
			w.WriteHeader(201)
		}
		log.Print("url created!")
		json.NewEncoder(w).Encode(url.ShortUrl)
	}

}


func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	fmt.Println("Rest API v2.0 - Mux Routers")
	dbinit()
	handleRequests()
}




