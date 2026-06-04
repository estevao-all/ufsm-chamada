package routes

import (
	"backend/utils"
	"bytes"
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"

	"github.com/liyue201/goqr"
)

type ImageDataResponse struct {
	Data string `json:"data"`
}

func read_image(imgData []byte) (string, error) {

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return "", err
	}

	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		return "", err
	}

	if len(qrCodes) == 0 {
		return "", errors.New("No QR code found in image.")
	}

	return string(qrCodes[0].Payload), nil
}

func HandleQRCodeUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, "Method not allowed")
		return
	}

	// (32MB max file size)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusBadRequest, "Error parsing form: "+err.Error())
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusBadRequest, "Error retrieving image: "+err.Error())
		return
	}
	defer file.Close()

	imageBytes, err := io.ReadAll(file)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading image: "+err.Error())
		return
	}

	qrResult, err := read_image(imageBytes)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error decoding image: "+err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, ImageDataResponse{Data: qrResult}) // replace with a database write
}
