package utility

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Mkdir(path string, perm os.FileMode) error {

	err := os.Mkdir(path, perm)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("directory %s already exists", path)
		}
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}

func WriteResponse2File(filename string, destDirectory string, response *http.Response) string {

	contentType := response.Header.Get("content-type")
	outputPath := destDirectory + filename

	switch contentType {
	case "audio/mpeg":
		outputPath = outputPath + ".mp3"
	case "audio/wav":
		outputPath = outputPath + ".wav"
	case "image/jpeg":
		outputPath = outputPath + ".jpg"
	case "image/png":
		outputPath = outputPath + ".png"
	case "application/octet-stream":
		outputPath = outputPath + ".lrc"
	}

	file, err := os.Create(outputPath)
	if err != nil {
		print(err)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("ERROR Copying")
	}

	file.Close()
	response.Body.Close()

	// Special case for wav files (which needs to be created beforehand)
	if contentType == "audio/wav" {
		return WavToMp3(outputPath)
	}

	return outputPath
}
