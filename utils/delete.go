package utils

import (
	"io/ioutil"
	"os"
	"strings"
	"tievs.com/km/models"
)
type Errors struct {
	E []error `json:"errors" bson:"errors"`
}

func Delete(item *models.Item) Errors {
	var e Errors
	urls := item.Urls
	for _, value := range urls{
		urlTrim := strings.TrimPrefix(value.Url,"http://upload.tievs.com/km/")
		folderName := strings.Split(urlTrim,"/")
		filePath := "/var/www/html/km/" + urlTrim
		err := os.Remove(filePath)
		if err != nil {
			e.E = append(e.E, err)
			continue
		}

		//Empty Folder Delete
		hasFolder := strings.Contains(folderName[0],"/")
		if hasFolder {
			folderPath := "/var/www/html/km/"+folderName[0]
			files, err := ioutil.ReadDir(folderPath)
			if err != nil {
				e.E = append(e.E, err)
				continue
			}
			if len(files) == 0 {
				err = os.Remove(folderPath)
				if err != nil {
					e.E = append(e.E, err)
					continue
				}
			}
		}

	}
	return e
}