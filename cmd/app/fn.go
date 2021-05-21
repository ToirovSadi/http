package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ToirovSadi/http/pkg/banners"
)

func uploadImage(item *banners.Banner, r *http.Request) (string, error) {
	if len(r.PostFormValue("image")) == 0 {
		return item.Image, nil
	}
	err := r.ParseMultipartForm(10 << 20) // size of coming data(image)
	if err != nil {
		log.Println("Err:app:uploadImage(): ", err)
		return "", err
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		log.Println("Err:app:uploadImage(): ", err)
		return "", err
	}
	defer file.Close()
	imageName := string(strconv.Itoa(int(item.ID)) + "." + getExtension(header.Filename))

	// tempFile, err := ioutil.TempFile("web/banners", imageName)
	// if err != nil {
	// 	log.Println("Err:app:uploadImage(): ", err)
	// 	return "", err
	// }
	// defer tempFile.Close()

	tempFile, err := os.Create("web/banners/" + imageName)
	if err != nil {
		return "", err
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Err:app:uploadImage(): ", err)
		return "", err
	}

	tempFile.Write(fileBytes)
	log.Println("imageName:", imageName)
	return imageName, nil
}

func getExtension(imageName string) string {
	return strings.Split(imageName, ".")[1]
}

func writeJson(w http.ResponseWriter, item interface{}) error {
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "appication/json")
	_, err = w.Write(data)
	return err
}
