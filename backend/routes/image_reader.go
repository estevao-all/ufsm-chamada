package routes

import (
	"backend/utils"
	"bytes"
	"errors"
	"image"
	_ "image/jpeg" // Register JPEG format
	_ "image/png"  // Register PNG format
	"io"
	"net/http"

	"github.com/liyue201/goqr"
)

type ImageDataResponse struct {
	Data string `json:"data"`
}

func read_image(imgData []byte) (string, error) {

	// Decode the raw data into a standard Go image.Image
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return "", err
	}

	// Recognize and parse the QR code patterns from the image
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		return "", err
	}

	// Print the payloads of any detected QR codes
	if len(qrCodes) == 0 {
		return "", errors.New("No qr code found in image")
	}

	return string(qrCodes[0].Payload), nil
}

// just reads the qrcodes for now but we probably want to
// store them in a db for later use.
func HandleQRCodeUpload(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		utils.WriteStatusAndLogInternally(w, http.StatusInternalServerError, "Method not allowed")
		return
	}

	// Parse multipart form (32MB max file size)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusBadRequest, "Error parsing form: "+err.Error())
		return
	}

	// Get the image file from form data
	file, _, err := r.FormFile("image") // "image" is the form field name
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusBadRequest, "Error retrieving image: "+err.Error())
		return
	}
	defer file.Close()

	// Read image bytes
	// Might be different from the actual esp implementation
	// good enough for now
	imageBytes, err := io.ReadAll(file)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error reading image: "+err.Error())
		return
	}

	// Decode QR code
	qrResult, err := read_image(imageBytes)
	if err != nil {
		utils.WriteStatusAndLogInternally(w,
			http.StatusInternalServerError, "Error decoding image: "+err.Error())
		return
	}

	utils.WriteJSON(w, http.StatusOK, ImageDataResponse{Data: qrResult})
}
