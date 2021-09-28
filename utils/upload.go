package utils

import (
	"fmt"
	"github.com/tievs/km/models"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func PostUpload(file *multipart.FileHeader) (models.File,error) {
	var File models.File
	src, err := file.Open()
	if err != nil {
		return File,err
	}
	defer src.Close()
	fileName := strings.ReplaceAll(file.Filename," ","-")
	File.FileName = fileName
	File.Url = "http://upload.tievs.com/km/"+fileName
	fmt.Println(fileName)
	// Destination
	dst, err := os.Create("./files/"+fileName)
	if err != nil {
		return File,err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return File,err
	}
	return File,nil
}

func PutUpload(file *multipart.FileHeader, timenow string) (models.File,error) {
	var File models.File
	folderNameWithSpace := strings.Split(timenow,".")
	folder := strings.ReplaceAll(folderNameWithSpace[0]," ","-")
	err := os.Mkdir("./files/"+folder, 0755)
	if err != nil {
		return File,err
	}
	src, err := file.Open()
	if err != nil {
		return File,err
	}
	defer src.Close()
	fileName := strings.ReplaceAll(file.Filename," ","-")
	File.FileName = fileName
	File.Url = "http://upload.tievs.com/km/"+folder+"/"+fileName
	fmt.Println(fileName)
	// Destination
	dst, err := os.Create("/var/www/html/km/"+folder+"/"+fileName)
	if err != nil {
		return File,err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return File,err
	}
	return File,nil
}
