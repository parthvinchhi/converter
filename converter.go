package converter

import (
	"log"
	"os/exec"
)

func Mp4ToMp3(inputFile, outputFile string) {

	// Construct the ffmpeg command
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-q:a", "0", "-map", "a", outputFile)

	// Run the command
	err := cmd.Run()
	if err != nil {
		log.Fatalf("ffmpeg command failed with %s\n", err)
	}

	log.Println("Conversion completed successfully")
}
