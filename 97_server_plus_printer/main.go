package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type server struct{}

//################## HTTP Handler ##############################################################################

/*func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	case "POST":
		// w.WriteHeader(http.StatusCreated)
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
		text := r.FormValue("text")

		t := time.Now().Format("2006-01-02 15:04:05")

		fmt.Println(text)
		printLKT(text + "\n\n         " + t)


	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))

	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}*/

func handlePic(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.Header.Get("Referer"), 302)

	fileName := "test.png"

	fmt.Println("Post success")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	tempFile, _ := os.Create(fileName)


	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		fmt.Printf("write error", err)
	}
	// PNG file created
	err = tempFile.Close()
	if err != nil {
		fmt.Printf("could not open", err)
	}

	printPic(fileName)



}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	//http.Handle("/", http.StripPrefix("/", h))
}

func getText(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	text := r.FormValue("text")

	t := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(text)
	printLKT(text + "\n\n\n         " + t + "\n", "test.txt")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/", getText).Methods("POST")
	r.HandleFunc("/upload", handlePic).Methods("POST")
	//err := http.ListenAndServeTLS(":8080", "https-server.crt", "https-server.key", nil);
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	//http.Handle("/", r)
	/*s := &server{}
	http.Handle("/", s)
	http.HandleFunc("/upload/", handlePic)
	log.Fatal(http.ListenAndServe(":8080", nil))*/
}

func printLKT(t, file string) {
	printText := []byte(t)
	err := ioutil.WriteFile("./" + file, printText, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// sh := `lp test.txt -d LKT`
	sh := "lp " + file + " -d LKT"
	args := strings.Split(sh, " ")

	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s \n", b)
}

func printPic(file string) {
	sh := "lp " + file + " -d LKT"
	args := strings.Split(sh, " ")

	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%s \n", b)
}