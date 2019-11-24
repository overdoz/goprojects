package main

import (
	"encoding/json"
	"fmt"
	"github.com/disintegration/imaging"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode/utf8"
)


type Message struct {
	Text string
}

func home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		dir, _ := os.Getwd()
		http.FileServer(http.Dir(dir)).ServeHTTP(w,r)

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		re := r.FormValue("type")


		switch re {
		case "text":
			decoder := json.NewDecoder(r.Body)
			var m Message
			err := decoder.Decode(&m)
			if err != nil {
				panic(err)
			}
			t := time.Now().Format("2006-01-02 15:04:05")

			sp := strings.Split(m.Text, " ")[0]

			if strings.Contains(sp, ":") {
				log.Print(findNames(m.Text))
				printLKT(findNames(m.Text) + "\n\n\n\n         " + t + "\n\n   ", "test.txt")
			} else {
				log.Print(m.Text)
				printLKT(m.Text + "\n\n\n\n         " + t + "\n\n   ", "test.txt")
			}
			http.Redirect(w, r, r.Header.Get("Referer"), 302)

		case "files":
			fileName := "test.png"

			log.Println("processing image...")

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

			// create png file
			tempFile, _ := os.Create(fileName)

			// read request body from client
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}

			// write incoming file to new file
			_, err = tempFile.Write(fileBytes)
			if err != nil {
				fmt.Printf("write error", err)
			}
			// close temporary file
			err = tempFile.Close()
			if err != nil {
				fmt.Printf("could not open", err)
			}
			// ------------------------------------------
			// Read image from file that already exists
			src, _ := imaging.Open(fileName)

			src = imaging.Resize(src, 800, 0, imaging.Lanczos)

			src = imaging.AdjustBrightness(src, 30)

			src = imaging.Grayscale(src)

			// src = imaging.AdjustContrast(src, -20)

			src = imaging.AdjustContrast(src, -20)

			src = imaging.AdjustGamma(src, 0.75)

			err = imaging.Save(src, "test.png")
			if err != nil {
				log.Fatalf("failed to save image: %v", err)
			}

			// print the png file
			printPic(fileName)

			// redirect to previous site
			http.Redirect(w, r, r.Header.Get("Referer"), 302)

		default:
			http.Redirect(w, r, r.Header.Get("Referer"), 302)
			log.Print("couldn't read query")
		}



	default:
		log.Print("couldn't handle request")
	}
}



func main() {
	port := ":8001"

	http.HandleFunc("/", home)

	log.Println("listening on port " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Print("line 112")
		log.Fatal(err)
	}
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

func findNames(s string) string {
	maxCharsPerLine := 35
	stringArray := strings.Split(s, " ")
	longestString := 0

	var in []int
	q := make(map[string]string)

	// find longstes name
	for i, s := range stringArray {
		if strings.Contains(s, ":") {
			in = append(in, i)
			if utf8.RuneCountInString(s) > longestString {
				longestString = utf8.RuneCountInString(s)
			}
		}
	}


	for i := 0; i < len(in); i++ {
		if i < len(in)-1 {
			currentName := in[i]
			nextName := in[i+1]

			// connect quote to person
			q[stringArray[currentName]] = strings.Join(stringArray[currentName+1:nextName], " ")

		} else {
			currentName := in[i]
			q[stringArray[currentName]] = strings.Join(stringArray[currentName+1:], " ")
		}
	}

	outputString := ""



	for i, s := range q {
		outputString = outputString + i + "\n" + formatString(s, maxCharsPerLine, longestString) + "\n\n"
	}

	return outputString
}

func formatString(text string, maxLen, longest int) string {
	output := ""
	col := 0
	max := maxLen - longest
	tempText := strings.Split(text, " ")
	// fmt.Println(tempText)

	for _, v := range tempText {
		if col == 0 {
			output = output + strings.Repeat(" ", longest)
			col += longest
		} else if (col + utf8.RuneCountInString(v)) < max {
			output = output + v + " "
			col += utf8.RuneCountInString(v) + 1
		} else {
			output = output + "\n" + strings.Repeat(" ", longest) + v + " "
			col = longest + utf8.RuneCountInString(v) + 1
		}
	}
	// fmt.Println(output)

	return output


	//output = output + strings.Repeat(" ", longest)
}