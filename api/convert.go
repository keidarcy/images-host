package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/convert", convertHandler)
	http.ListenAndServe(":8080", nil)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a POST request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Only POST requests are allowed")
		return
	}

	// Parse the form data and retrieve the file
	err := r.ParseMultipartForm(32 << 20) // 32 MB
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Failed to parse form data")
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Failed to retrieve file from form data")
		return
	}
	defer file.Close()

	fmt.Println(header.Filename)

	// Check if the file is a HEIC file
	if !strings.HasSuffix(header.Filename, ".HEIC") {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "File is not a HEIC file")
		return
	}

	// Create a new JPG file
	newFileName := strings.TrimSuffix(header.Filename, ".heic") + ".jpg"
	newFile, err := os.Create(newFileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to create new JPG file")
		return
	}
	defer newFile.Close()

	// Read the HEIC file and decode it to an image
	heicData, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to read HEIC file")
		return
	}
	img, err := decodeHEIC(heicData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to decode HEIC file")
		return
	}

	// Encode the image to a JPG file
	err = jpeg.Encode(newFile, img, &jpeg.Options{Quality: 100})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to encode JPG file")
		return
	}

	// Set the response header and send the new JPG file as the response
	w.Header().Set("Content-Disposition", "attachment; filename="+newFileName)
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, newFileName)

	// Remove the new JPG file after it has been sent
	err = os.Remove(newFileName)
	if err != nil {
		fmt.Println("Failed to remove new JPG file")
	}
}
func decodeHEIC(data []byte) (*image.NRGBA, error) {
	// Save the HEIC data to a temporary file
	tmpFile, err := ioutil.TempFile("", "*.heic")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.Write(data)
	if err != nil {
		return nil, err
	}
	tmpFile.Close()

	// Use the sips command to convert the HEIC file to a JPEG file
	jpgFile := tmpFile.Name() + ".jpg"
	cmd := fmt.Sprintf("sips -s format jpeg %s --out %s", tmpFile.Name(), jpgFile)
	err = exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		return nil, err
	}

	// Read the JPEG file and decode it to an image
	imgFile, err := os.Open(jpgFile)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()
	img, err := jpeg.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	nrgba := image.NewNRGBA(img.Bounds())
	draw.Draw(nrgba, nrgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	// Remove the JPEG file after decoding it to an image
	err = os.Remove(jpgFile)
	if err != nil {
		fmt.Println("Failed to remove temporary JPEG file")
	}

	return nrgba, nil
}
