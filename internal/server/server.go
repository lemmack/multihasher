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
	Fnv    string `json:"fnv"`
	Md5    string `json:"md5"`
	Ripemd string `json:"ripemd160"`
}

// Handles the "/" (index) route
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index route response.")
}

// Handles the "/upload" route
func upload_handler(w http.ResponseWriter, r *http.Request) {

	// Allow CORS for a local client if one is specified
	if localClient != "" {
		w.Header().Set("Access-Control-Allow-Origin", localClient)
	}

	switch r.Method {
	case "POST":
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("file")

		if err != nil {
			fmt.Fprintf(w, "error forming file: %s", err)
			log.Println("error forming file:", err)
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Fprintf(w, "error forming byteslice from file: %s", err)
			log.Println("error forming byteslice from file:", err)
			return
		}

		j, err := make_hash_json(fileBytes)

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
	rawFnv, err := hash.BytesToFNV(b)

	if err != nil {
		return nil, err
	}

	rawMd5 := hash.BytesToMD5(b)
	rawRipemd, err := hash.BytesToRIPEMD(b)

	if err != nil {
		return nil, err
	}

	hashStruct := Hashes{Fnv: rawFnv, Md5: rawMd5, Ripemd: rawRipemd}
	jsonBytes, err := json.Marshal(hashStruct)

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// Sets up the routes and starts the server
func Start(ps string, lc string) {
	portString, localClient = ps, lc

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/upload", upload_handler)

	log.Printf("Starting server on port %s", portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}
