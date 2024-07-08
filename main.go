package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func setWallPaper(imagePath string) error {
	// Get the absolute path of the image
	absPath, err := filepath.Abs(imagePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Run the applescript to set the wallpaper
	cmd := exec.Command("osascript", "./set_wallpaper.scpt", absPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set wallpaper: %v, output: %s", err, output)
	}
	fmt.Printf("AppleScript output: %s\n", output)
	return nil
}

func main() {
	imagePath := "./python.jpg"

	// Check if the file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		log.Fatalf("Image file does not exist: %s", imagePath)
	}

	err := setWallPaper(imagePath)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Wallpaper set to", imagePath)
	}
}