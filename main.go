package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/toast.v1"
)

type Employee struct {
	Id         int
	Bookid     int
	Issuerid   int
	Issuedate  string
	Returndate string
}

type User struct {
	Id       int
	Name     string
	Email    string
	Phone    string
	Password string
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
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id, bookid, issuerid int
		var issuedate, returndate string
		err = selDB.Scan(&id, &bookid, &issuerid, &issuedate, &returndate)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Bookid = bookid
		emp.Issuerid = issuerid
		emp.Issuedate = issuedate
		emp.Returndate = returndate
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id, bookid, issuerid int
		var issuedate, returndate string
		err = selDB.Scan(&id, &bookid, &issuerid, &issuedate, &returndate)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Bookid = bookid
		emp.Issuerid = issuerid
		emp.Issuedate = issuedate
		emp.Returndate = returndate
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id, bookid, issuerid int
		var issuedate, returndate string
		err = selDB.Scan(&id, &bookid, &issuerid, &issuedate, &returndate)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Bookid = bookid
		emp.Issuerid = issuerid
		emp.Issuedate = issuedate
		emp.Returndate = returndate
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		bookid := r.FormValue("bookid")
		issuerid := r.FormValue("issuerid")
		issuedate := r.FormValue("issuedate")
		returndate := r.FormValue("returndate")
		insForm, err := db.Prepare("INSERT INTO Employee(bookid, issuerid,issuedate,returndate) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(bookid, issuerid, issuedate, returndate)
		log.Println("INSERT: bookid: " + bookid + " | issuedate: " + issuedate)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)

	notification := toast.Notification{
		Message: "Issuer info added succesfully.",
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		bookid := r.FormValue("bookid")
		issuerid := r.FormValue("issuerid")
		issuedate := r.FormValue("issuedate")
		returndate := r.FormValue("returndate")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET bookid=?, issuerid=? ,issuedate=?,returndate=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(bookid, issuerid, issuedate, returndate, id)
		log.Println("UPDATE: bookid: " + bookid + " | issuedate: " + issuedate)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)

	notification := toast.Notification{
		Message: "Issuer info updated succesfully.",
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)

	notification := toast.Notification{
		Message: "Issuer info deleted succesfully.",
	}
	err = notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Signup", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		password := r.FormValue("password")
		insForm, err := db.Prepare("INSERT INTO User(name,email,phone,password) VALUES(?,?,?,?)")
		if err != nil {
			http.Redirect(w, r, "/signup", 301)
			notification := toast.Notification{
				Message: "Signup unsuccesfull....",
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
			}
		}
		insForm.Exec(name, email, phone, password)
		log.Println("INSERT: Name: " + name + " | Email: " + email)
		http.Redirect(w, r, "/signin", 301)
		notification := toast.Notification{
			Message: "Signup succesfull, Please login.",
		}
		err = notification.Push()
		if err != nil {
			log.Fatalln(err)
		}
	}
	defer db.Close()
}

func Signin(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Signin", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		EmailFromDB := ""
		passwordFromDB := ""

		err := db.QueryRow("select email,password FROM User WHERE email = ?", email).Scan(&EmailFromDB, &passwordFromDB)
		if err != nil {
			panic(err.Error())
		} else {
			if password == passwordFromDB {

				log.Println("Login succesfull: Email: " + EmailFromDB)

				http.Redirect(w, r, "/", 301)

				notification := toast.Notification{
					Message: "Login succesfull...",
				}
				err := notification.Push()
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Println("Login unsuccesfull, data does not match")
				http.Redirect(w, r, "/signin", 301)
				notification := toast.Notification{
					Message: "Login unsuccesfull, email and pasword does not match.",
				}
				err := notification.Push()
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}
	defer db.Close()
}

func Logout(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/signin", 301)

}

func main() {
	log.Println("Server started on: http://localhost:9090")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	http.ListenAndServe(":9090", nil)
}
