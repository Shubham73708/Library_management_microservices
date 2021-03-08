package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/toast.v1"
)

type Guruom struct {
	Id      int
	Name    string
	Email   string
	Comment string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Shubh@123"
	dbName := "employee"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Index", nil)

}

func Service(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Service", nil)

}

func About(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "About", nil)

}

func Product(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Product", nil)

}

func Contact(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Contact", nil)

}

func Contactus(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		comment := r.FormValue("comment")

		insForm, err := db.Prepare("INSERT INTO Guruom(name,email,comment) VALUES(?,?,?)")
		if err != nil {
			http.Redirect(w, r, "/contact", 301)
			if err != nil {
				log.Fatalln(err)
			}
		}
		insForm.Exec(name, email, comment)
		log.Println("INSERT: Name: " + name + " | Email: " + email)
		http.Redirect(w, r, "/signin", 301)
		notification := toast.Notification{
			Message: "succesfully ping us, we will reach you soon.",
		}
		err = notification.Push()
		if err != nil {
			log.Fatalln(err)
		}
	}
	defer db.Close()
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/service", Service)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/product", Product)
	http.HandleFunc("/contactus", Contactus)

	http.ListenAndServe(":8080", nil)
}
