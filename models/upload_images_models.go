package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadImage(type_foto string, id_data string, files multipart.File, header *multipart.FileHeader) string {

	b := make([]rune, 40)
	bc := make([]rune, 16)
	t := time.Now()
	pathImage := ""
	pathId := []byte(id_data)
	pathIdMD5 := md5.Sum(pathId)
	var letterRunes = []rune("abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	filename := strings.Join(strings.Fields(header.Filename), "")

	for i := range b {

		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	for i := range bc {

		bc[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	filename = string(b) + t.Format("20060102150405.0") + string(bc) + filename
	folderID := hex.EncodeToString(pathIdMD5[:])

	switch type_foto {
	case "lead":
		pathImage = "./pbc/img/lead/" + folderID + "/"
	case "target":
		pathImage = "./pbc/img/target/" + folderID + "/"
	case "document":
		pathImage = "./pbc/img/document/" + folderID + "/"

	default:
		pathImage = "./prv/img/all/" + folderID + "/"

	}

	errMk := os.MkdirAll(pathImage, 0777)

	if errMk != nil {
		fmt.Println(errMk.Error())
		return ""
	} else {

		out, err := os.Create(pathImage + filename)

		if err != nil {
			fmt.Println(err.Error())
			return ""
		} else {

			defer out.Close()

			_, err = io.Copy(out, files)

			if err != nil {
				fmt.Println(err.Error())
				return ""
			}

			// return string(filepath.Dir(pathImage)+filepath.FromSlash("/"+filename))
			return string(filepath.FromSlash(folderID + "/" + filename))

		}

	}

	//return
}
