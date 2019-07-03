package main
import (
        "fmt"
        "net/http"
        "gopkg.in/yaml.v2"
)

type User struct {
        Name string
        Email string
        ID int
}

func userRouter(w http.ResponseWriter, r *http.Request) {
        ourUser := User{}
        ourUser.Name = "Bill Smith"
        ourUser.Email = "bill.smith@example.com"
        ourUser.ID = 100

        output, _ := yaml.Marshal(&ourUser)
        fmt.Fprintln(w, string(output))
}

func main() {
        fmt.Println("Starting YAML server")
        http.HandleFunc("/user", userRouter)
        http.ListenAndServe(":8080", nil)
}
