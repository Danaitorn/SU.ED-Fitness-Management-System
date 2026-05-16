package utils

import (
	"mime/multipart"
	"net/http"
	"os"
)
func UploadToSupabase(file multipart.File, filename string) (string, error) {

	url := "https://sbxdfneqhidfxeisfuvv.supabase.co/storage/v1/object/public/news/" + filename

	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_SERVICE_KEY"))
	req.Header.Set("Content-Type", "image/jpeg")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	publicURL := "https://sbxdfneqhidfxeisfuvv.supabase.co/storage/v1/object/public/news/" + filename

	return publicURL, nil
}