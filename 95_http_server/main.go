package main

import (
	"encoding/json"
	"github.com/disintegration/imaging"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const FILENAME_TXT = "test.txt"
const FILENAME_PNG = "test.png"
const LINE_WIDTH = 32



type Message struct {
	Text string
}


func home(w http.ResponseWriter, r *http.Request) {

	// http method
	switch r.Method {

	// serve React PWA
	// https://github.com/overdoz/airprinter
	case http.MethodGet:

		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		dir, _ := os.Getwd()
		http.FileServer(http.Dir(dir)).ServeHTTP(w,r)

	// two options to send files
	// @param text: send formatted text
	// @param file: send multipart/form-data
	// TODO: right now you either send a file or plain text, but not combined
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			log.Printf("ParseForm() err: %v", err)
			return
		}

		// POST Request params
		re := r.FormValue("type")

		// determine if type =? text or file
		switch re {

		// text will be formatted correctly
		// print timestamp at the end of the sheet
		case "text":
			decoder := json.NewDecoder(r.Body)
			var m Message
			err := decoder.Decode(&m)
			if err != nil {
				log.Fatal("JSON Decoder failed: ", err)
			}

			t := time.Now().Format("2006-01-02 15:04:05")

			sp := strings.Split(m.Text, " ")

			contains := false

			// if first string is a name followed by an ":", the text is declared as quote
			if strings.Contains(sp[0], ":") {
				g := m.Text

				for i, v := range sp {
					re, _ := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, v)
					if re {
						contains = true

						// extract given date
						ownDate := v

						// cut date from text
						splitted := strings.Split(g, " ")
						splitted = splitted[:i]
						g = strings.Join(splitted, " ")	//g[:i]

						output := findNames(g) + "\n\n\n\n                  " + ownDate

						log.Print("\n " + output)

						// send to printer
						printLKT(output, FILENAME_TXT)
					}
				}
				// print text with current timestamp
				if !contains {
					printLKT(findNames(g) + "\n\n\n\n         " + t, FILENAME_TXT)
				}
			} else {
				log.Print("\n " + m.Text)
				printLKT(m.Text + "\n\n\n\n         " + t + "\n\n   ", FILENAME_TXT)
			}
			// redirect to previous site
			http.Redirect(w, r, r.Header.Get("Referer"), 302)

		// POST Request params (type = "files")
		case "files":

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
			tempFile, _ := os.Create(FILENAME_PNG)

			// read request body from client
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println("Cannot read request body: ", err)
			}

			// write incoming file to new file
			_, err = tempFile.Write(fileBytes)
			if err != nil {
				log.Printf("write error", err)
			}
			// close temporary file
			err = tempFile.Close()
			if err != nil {
				log.Printf("could not open", err)
			}

			// Read image from file that already exists
			err = adjustImage(FILENAME_PNG)
			if err != nil {
				log.Fatalf("failed to save image after edit: %v", err)
			}

			// print the png file
			printPic(FILENAME_PNG)

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



// file should have the ending .txt
func printLKT(t, file string) {
	printText := []byte(t)

	err := ioutil.WriteFile("./" + file, printText, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// print command to LKT
	sh := "lp " + file + " -d LKT"

	args := strings.Split(sh, " ")

	cmd := exec.Command(args[0], args[1:]...)

	// execute command
	b, err := cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s \n", b)
}

// file should have the ending .png
func printPic(file string) {
/*	sh := "lp " + file + " -d LKT"

	args := strings.Split(sh, " ")

	cmd := exec.Command(args[0], args[1:]...)

	b, err := cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
	}
	log.Printf("%s \n", b)*/
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr:    "cm",
		Size:       gofpdf.SizeType{Wd: 7.1, Ht: 1},
	})

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		//
	}

	sh := "lp hello.pdf -d LKT"

	args := strings.Split(sh, " ")

	cmd := exec.Command(args[0], args[1:]...)

	_, err = cmd.CombinedOutput()

	if err != nil {
		log.Println(err)
	}

}

func findNames(s string) string {
	longestString := 0
	outputString := ""
	stringArray := strings.Split(s, " ")

	// slice without predefined length
	var in []int

	// [name]quote
	q := make(map[string]string)

	// find longest name
	for i, s := range stringArray {
		if strings.Contains(s, ":") {
			// save index of names
			in = append(in, i)
			if utf8.RuneCountInString(s) > longestString {
				longestString = utf8.RuneCountInString(s)
			}
		}
	}

	// connect quote to name
	for i := 0; i < len(in); i++ {
		if i < len(in)-1 {
			currentName := in[i]
			nextName := in[i+1]
			// connect quote to person
			// +1 to cut name at the beginning
			q[stringArray[currentName]] = strings.Join(stringArray[currentName+1:nextName], " ")
		} else {
			// if it's the last person, take the rest of the string
			currentName := in[i]
			q[stringArray[currentName]] = strings.Join(stringArray[currentName+1:], " ")
		}
	}

	// concat names and quotes
	for i, s := range q {
		outputString = outputString + i + "\n" + formatString(s, LINE_WIDTH, longestString) + "\n\n"
	}
	return outputString
}


// adjust width
func formatString(text string, maxLen, longest int) string {
	output := ""
	col := 0
	max := maxLen - longest
	tempText := strings.Split(text, " ")

	for _, v := range tempText {
		if col == 0 {
			output = output + strings.Repeat(" ", longest) + v + " "
			col += longest
		} else if (col + utf8.RuneCountInString(v)) < max {
			output = output + v + " "
			col += utf8.RuneCountInString(v) + 1
		} else {
			output = output + "\n" + strings.Repeat(" ", longest) + v + " "
			col = longest + utf8.RuneCountInString(v) + 1
		}
	}
	return output
}

// edit immage
func adjustImage(file string) error {
	src, _ := imaging.Open(file)
	src = imaging.Resize(src, 800, 0, imaging.Lanczos)
	src = imaging.AdjustBrightness(src, 30)
	src = imaging.Grayscale(src)
	src = imaging.AdjustContrast(src, -20)
	src = imaging.AdjustGamma(src, 0.75)
	return imaging.Save(src, FILENAME_PNG)
}

func main() {
	port := ":8001"

	http.HandleFunc("/", home)

	log.Println("listening on port " + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
