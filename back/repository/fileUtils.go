package repository

import (
	"log"
	"mime/multipart"
)

func Upload(path string, file multipart.File) (string,error) {



	log.Println(path)
	return "",nil
}
