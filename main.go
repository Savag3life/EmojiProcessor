package main

import (
	"archive/zip"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"os"
	"strconv"

	_ "github.com/nfnt/resize"
)

func main() {
	http.HandleFunc("/", uploadForm)
	http.HandleFunc("/process", processImage)
	http.HandleFunc("/download", downloadZip)
	http.ListenAndServe(":8080", nil)
}

func uploadForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<html>
	<head><title>PNG Emote Processor</title>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">
	</head>
	<body class="container mt-5">
	<h2>Upload PNG Image</h2>
	<form action="/process" method="post" enctype="multipart/form-data">
		<div class="mb-3">
			<label class="form-label">Select PNG:</label>
			<input class="form-control" type="file" name="image" required>
		</div>
		<div class="mb-3">
			<label class="form-label">Emote ID:</label>
			<input class="form-control" type="text" name="emote_id" required>
		</div>
		<div class="mb-3">
			<label class="form-label">Emoji Size (px):</label>
			<input class="form-control" type="number" name="emoji_size" value="32" required>
		</div>
		<button class="btn btn-primary" type="submit">Process</button>
	</form>
	</body>
	</html>
	`)
}

func processImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	emoteID := r.FormValue("emote_id")
	emojiSize, err := strconv.Atoi(r.FormValue("emoji_size"))
	if err != nil || emojiSize <= 0 {
		http.Error(w, "Invalid emoji size", http.StatusBadRequest)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Error decoding image", http.StatusInternalServerError)
		return
	}

	bounds := img.Bounds()
	rows := bounds.Dy() / emojiSize
	cols := bounds.Dx() / emojiSize
	zipPath := "output.zip"

	zipFile, err := os.Create(zipPath)
	if err != nil {
		http.Error(w, "Failed to create ZIP file", http.StatusInternalServerError)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	var emojiGrid string
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			rect := image.Rect(c*emojiSize, r*emojiSize, (c+1)*emojiSize, (r+1)*emojiSize)
			subImg := img.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(rect)

			filename := fmt.Sprintf("%s_%d_%d.png", emoteID, r, c)
			f, err := zipWriter.Create(filename)
			if err != nil {
				http.Error(w, "Error adding file to ZIP", http.StatusInternalServerError)
				return
			}

			png.Encode(f, subImg)
			emojiGrid += fmt.Sprintf(":%s_%d_%d:", emoteID, r, c)
		}
		emojiGrid += "<br>"
	}

	fmt.Fprintf(w, `
	<html>
	<head><title>PNG Emote Processor</title>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css">
	</head>
	<body class="container mt-5">
	<h2>Emoji Grid</h2>
	<div class="mb-3">%s</div>
	<a class="btn btn-primary" href="/download" download>Download ZIP</a>
	</body>
	</html>
	`, emojiGrid)
}

func downloadZip(w http.ResponseWriter, r *http.Request) {
	zipPath := "output.zip"
	w.Header().Set("Content-Disposition", "attachment; filename=output.zip")
	w.Header().Set("Content-Type", "application/zip")
	http.ServeFile(w, r, zipPath)
}
