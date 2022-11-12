package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lemmack/multihasher/hash"
)

const localClient string = "http://127.0.0.1:5500"
const portString string = ":8000"

type jsonHash struct {
	Hash string `json:"hash"`
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index route response.")
}

func upload_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", localClient)

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

		j, err := make_hash(file)

		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			log.Println("error:", err)
			return
		}

		log.Println("json: ", string(j))

		fmt.Fprintf(w, "%s", j)
	default:
		fmt.Fprintf(w, "Sorry, only POST method supported")
	}
}

func make_hash(f io.Reader) ([]byte, error) {
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
