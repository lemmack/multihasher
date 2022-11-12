package main

import (
	//"io/ioutil"
	//"github.com/lemmack/multihasher/hash"
	"github.com/lemmack/multihasher/internal/server"
)

// Consider adding a version variable which is assigned at runtime by a makefile
// var version string

//func uploadFile(w http.ResponseWriter, r *http.Request) {

// var maxFileSize int64 = 10 << 20
// err := r.ParseMultipartForm(maxFileSize)

// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// file, header, err := r.FormFile("myFile")

// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// defer file.Close()
// fmt.Printf("File Name: %+v\n", header.Filename)

// tempFile, err := ioutil.TempFile("temp-files", "upload-*")

// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// defer tempFile.Close()

// fileBytes, err := ioutil.ReadAll(file)

// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// tempFile.Write(fileBytes)

//}

func main() {
	server.Start()
}
