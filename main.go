package main
#powered by https://github.com/pabloos 
import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var dir string

const (
	slash           = "/"
	filePermissions = 0755
)

func endsWithSlash(path string) bool {
	lastChar := path[len(path)-1:]
	if lastChar == slash {
		return true
	}

	return false
}

func correctPath(path string) string {
	if !endsWithSlash(path) {
		return fmt.Sprintf("%s%s", path, slash)
	}

	return path
}

func createDir(path string) {
	if !dirExists(path) {
		err := os.Mkdir(path, filePermissions)

		if err != nil {
			log.Print(err)
		}
	}
}

func dirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(dir+header.Filename, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.StringVar(&dir, "dir", "./src", "directorio donde se guardan los ficheros")
	flag.Parse()

	dir = correctPath(dir)
	createDir(dir)

	files := http.FileServer(http.Dir("public"))
	http.Handle("/", files)
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
