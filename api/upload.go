package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

var validContentTypes map[string]struct{} = map[string]struct{}{
	"image/png":  {},
	"image/jpeg": {},
	"image/webp": {},
	"image/gif":  {},
}

type uploadResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (c *ApiClient) uploadFile(filePath string) (string, error) {
	var fileName string
	var fileBytes []byte
	var contentType string
	var err error

	if strings.Contains(filePath, "http") || strings.Contains(filePath, "https") {
		resp, err := http.Get(filePath)

		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		fileBytes, err = io.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		contentType = http.DetectContentType(fileBytes)

		splitPath := strings.Split(filePath, "/")
		fileName = splitPath[len(splitPath)-1]
	} else {
		fileBytes, err = os.ReadFile(filePath)

		if err != nil {
			return "", err
		}

		contentType = http.DetectContentType(fileBytes)

		splitPath := strings.Split(filePath, "/")
		fileName = splitPath[len(splitPath)-1]
	}

	if _, ok := validContentTypes[contentType]; !ok {
		return "", fmt.Errorf("error: invalid content type for uploaded media")
	}

	if len(fileBytes) > 5*1024*1024 {
		return "", fmt.Errorf("error: file size exceeds 5MB limit")
	}

	if !strings.Contains(fileName, ".") {
		switch contentType {
		case "image/png":
			fileName += ".png"
		case "image/jpeg":
			fileName += ".jpg"
		case "image/webp":
			fileName += ".webp"
		case "image/gif":
			fileName += ".gif"
		}
	}

	fmt.Println(contentType)

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	mimeHeaders := textproto.MIMEHeader{
		"Content-Disposition": []string{fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", fileName)},
		"Content-Type":        []string{contentType},
	}

	part, err := writer.CreatePart(mimeHeaders)

	if err != nil {
		return "", err
	}

	bytesWritten, err := part.Write(fileBytes)

	if err != nil {
		return "", err
	}

	if bytesWritten != len(fileBytes) {
		return "", fmt.Errorf("error: failed to write all bytes to multipart form")
	}

	err = writer.Close()

	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, "https://four.meme/meme-api/v1/private/token/upload", &body)

	if err != nil {
		return "", err
	}

	request.Header = http.Header{
		"accept":          []string{"application/json, text/plain, */*"},
		"accept-language": []string{"en-US,en;q=0.9"},
		"content-type":    []string{writer.FormDataContentType()},
		"origin":          []string{"https://four.meme"},
		"user-agent":      []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36"},
		"meme-web-access": []string{c.authToken},
	}

	response, err := c.httpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	var jsonResponse uploadResponse

	err = json.NewDecoder(response.Body).Decode(&jsonResponse)

	if err != nil {
		return "", err
	}

	if jsonResponse.Code != 0 || jsonResponse.Msg != "success" {
		return "", fmt.Errorf("error from server when uploading file: %s", jsonResponse.Msg)
	}

	return jsonResponse.Data, nil
}
