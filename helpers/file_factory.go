package helpers

import (
	"encoding/base64"
	"net/http"
	"os"
)

/*
UploadFile - Função responsável por liberar o publish do item
*/
func UploadFile(w http.ResponseWriter, name, item string) string {
	dec, err := base64.StdEncoding.DecodeString(item)

	if err != nil {
		return ""
	}

	f, err := os.Create(name + ".pdf")

	if err != nil {
		return ""
	}
	f.Write(dec)
	f.Close()
	AWS().UploadFileToS3(f.Name())
	os.Remove(f.Name())

	return "https://s3.amazonaws.com/" + configuration.Amazon.Bucket + "/" + configuration.Amazon.Prefix + f.Name()
}
