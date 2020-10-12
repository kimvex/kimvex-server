package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

//IMG struct for response cloudinary
type IMG struct {
	AssetID          string   `json:"asset_id"`
	PublicID         string   `json:"public_id"`
	Version          int      `json:"version"`
	VersionID        string   `json:"version_id"`
	Signature        string   `json:"signature"`
	Width            int      `json:"width"`
	Height           int      `json:"height"`
	Format           string   `json:"format"`
	ResourceType     string   `json:"resource_type"`
	CreatedAt        string   `json:"created_at"`
	Tags             []string `json:"tags"`
	Bytes            int      `json:"bytes"`
	Type             string   `json:"type"`
	ETag             string   `json:"etag"`
	Placeholder      bool     `json:"placeholder"`
	URL              string   `json:"url"`
	SecureURL        string   `json:"secure_url"`
	AccessMode       string   `json:"access_mode"`
	OriginalFilename string   `json:"original_filename"`
}

//UploadImg for upload images to cloudinary
func UploadImg(file *multipart.FileHeader, folder string) IMG {
	var CloudinaryResponse IMG
	sr, yr := file.Open()
	fmt.Println(sr, yr)
	// out, _ := os.Open(file.Filename)
	// defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("upload_preset", "wnnj4a59")
	_ = writer.WriteField("folder", folder)
	part, _ := writer.CreateFormFile("file", file.Filename)
	io.Copy(part, sr)
	writer.Close()

	r, _ := http.NewRequest("POST", "https://api.cloudinary.com/v1_1/h27hacklab/image/upload", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	// request.Header.Set("Content-Type", "image/png")
	r.Header.Add("api_key", "766496458317643")
	client := &http.Client{}

	response, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// content, err := ioutil.ReadAll(response.Body)
	json.NewDecoder(response.Body).Decode(&CloudinaryResponse)

	if err != nil {
		log.Fatal(err)
	}

	return CloudinaryResponse
}
