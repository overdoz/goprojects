package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"goprojects/25_mongodb/controllers"
	"net/http"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// connect to your local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

//func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	s := `<!DOCTYPE html>
//<html lang="en">
//<head>
//<meta charset="UTF-8">
//<title>Index</title>
//</head>
//<body>
//<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
//</body>
//</html>`
//
//	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte(s))
//}

