package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lemmack/multihasher/hash"
)

var portString string
var localClient string

type Hashes struct {
	Fnv string `json:"fnv"`
	Md5 string `json:"md5"`
}

// Handles the "/" (index) route
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index route response.")
}

// Handles the "/upload" route
func upload_handler(w http.ResponseWriter, r *http.Request) {

	if localClient != "" {
		w.Header().Set("Access-Control-Allow-Origin", localClient) // Allow CORS for a local client
	}

	switch r.Method {
	case "POST":
		r.ParseMultipartForm(32 << 20) // Parses the request body as multipart/form-data

		file, _, err := r.FormFile("file") // Retrieve the file from the form data

		if err != nil {
			fmt.Fprintf(w, "error forming file: %s", err)
			log.Println("error forming file:", err)
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file) // Reads all the contents of the file into a byte slice

		if err != nil {
			fmt.Fprintf(w, "error forming byteslice from file: %s", err)
			log.Println("error forming byteslice from file:", err)
			return
		}

		j, err := make_hash_json(fileBytes) // Create a json object (as a byte slice) containing the hashes of the file

		if err != nil {
			fmt.Fprintf(w, "error making hash json: %s", err)
			log.Println("error making hash json:", err)
			return
		}

		log.Println("json: ", string(j))
		fmt.Fprintf(w, "%s", j) // Write the json as a string to the response
	default:
		fmt.Fprintf(w, "Sorry, only POST method supported")
	}
}

// Takes a byte slice b and returns a json object (as a byte slice) containing multiple hashes of the data in b
func make_hash_json(b []byte) ([]byte, error) {
	rawFnv, err := hash.BytesToFNV(b) // Generate FNV hash

	if err != nil {
		return nil, err
	}

	rawMd5 := hash.BytesToMd5(b) // Generate md5 hash

	// Form the json object containing the generated hashes
	jh := Hashes{Fnv: rawFnv, Md5: rawMd5}
	j, err := json.Marshal(jh)

	if err != nil {
		return nil, err
	}

	return j, nil
}

// Sets up the routes and starts the server
func Start(ps string, lc string) {
	portString, localClient = ps, lc

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/upload", upload_handler)

	log.Printf("Starting server on port %s", portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}
