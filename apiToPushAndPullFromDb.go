package main

import (
        "database/sql"
        "encoding/json"
        "fmt"
        _ "github.com/go-sql-driver/mysql"
        "github.com/gorilla/mux"
        "net/http"
//      "log"
)

var database *sql.DB

type Users struct {
        Users []User `json:"users"`
}
type User struct {
        ID int "json:id"
        Name string "json:username"
        Email string "json:email"
        First string "json:first"
        Last string "json:last"
}

type CreateResponse struct {
        Error string "json:error"
}

func UserCreate(w http.ResponseWriter, r *http.Request){
        NewUser := User{}
        NewUser.Name = r.FormValue("user")
        NewUser.Email = r.FormValue("email")
        NewUser.First = r.FormValue("first")
        NewUser.Last = r.FormValue("last")
        output, err := json.Marshal(NewUser)
        fmt.Println(string(output))
        if err !=nil {
                fmt.Println("something went wrong")
        }
        Response := CreateResponse{}
        sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "'"
        q, err := database.Exec(sql)
        if err != nil {
                Response.Error = err.Error()
                fmt.Println(err)
        }
        fmt.Println(q)
        createOutput, _:= json.Marshal(Response)
        fmt.Fprintln(w, string(createOutput))
}

func UserRetrive(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Pragma", "no-cache")
        rows, _ := database.Query("select * from users Limit 10")
        Response := Users{}

        for rows.Next() {
                user := User{}
                rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)
                Response.Users = append(Response.Users, user)
        }
        output, _ := json.Marshal(Response)
        fmt.Fprintln(w, string(output))
}

func main(){
        db,err := sql.Open("mysql", "kanshu:Feb2019$@(localhost:3307)/test?charset=utf8")
        if err != nil {
                fmt.Println(err)
        }
        database = db

        routes := mux.NewRouter()
        routes.HandleFunc("/api/users", UserCreate).Methods("POST")
        routes.HandleFunc("/api/users", UserRetrive).Methods("GET")
        http.Handle("/", routes)
        http.ListenAndServe(":8080", nil)
}
