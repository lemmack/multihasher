package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lemmack/multihasher/hash"
)

const portString string = ":8000"                  // Port the server will run on
const localClient string = "http://127.0.0.1:5500" // Address of a local client for local testing

type jsonHash struct {
	Hash string `json:"hash"`
}

// Handles the "/" (index) route
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index route response.")
}

// Handles the "/upload" route
func upload_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", localClient) // Allow CORS for a local client

	switch r.Method {
	case "POST":
		r.ParseMultipartForm(32 << 20) // Parses a request body as multipart/form-data

		file, _, err := r.FormFile("file") // Retrieve the file from form data

		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			log.Println("error:", err)
			return
		}
		defer file.Close()

		j, err := make_hash_json(file)

		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			log.Println("error:", err)
			return
		}

		log.Println("json: ", string(j))
		fmt.Fprintf(w, "%s", j) // Write the hash as a json string to the response
	default:
		fmt.Fprintf(w, "Sorry, only POST method supported")
	}
}

// Takes an io.Reader (such as a multipart.File), generates its FNV hash, and returns the hash in json format
func make_hash_json(f io.Reader) ([]byte, error) {
	rawHash := hash.ReaderToFNV(f)
	jh := jsonHash{Hash: rawHash}
	j, err := json.Marshal(jh)

	if err != nil {
		return nil, err
	}

	return j, nil
}

func Start() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/upload", upload_handler)

	log.Printf("Starting server on port %s", portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}
