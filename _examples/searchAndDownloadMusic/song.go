package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kingmariano/omnicron-go"
)

func main() {
	apiKey := "YOUR_API_KEY"
	client := omnicron.NewClient(apiKey, omnicron.WithBaseURL("https://omnicron-latest.onrender.com/"))

	matchingSongsResponse, err := client.MusicSearch(context.Background(), &omnicron.MusicSearchRequest{
		Song: "girl of my dreams rodwave", // example
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	// Song matches printed {
	// 	"response": [
	// 	  {
	// 		"shazam_url": "https://www.shazam.com/track/596602721/girl-of-my-dreams",
	// 		"song_image": "https://is1-ssl.mzstatic.com/image/thumb/Music116/v4/f4/51/17/f451178a-ab05-98b5-ff60-8502804db9b9/21UM1IM54282.rgb.jpg/400x400cc.jpg",
	// 		"song_name": "Girl Of My Dreams - Juice WRLD, SUGA \u0026 BTS"
	// 	  },
	// 	  {
	// 		"shazam_url": "https://www.shazam.com/track/511208818/girl-of-my-dreams",
	// 		"song_image": "https://is1-ssl.mzstatic.com/image/thumb/Music115/v4/4c/19/17/4c1917ae-0b9d-c0b6-3f17-b800c93d3214/808391079595.jpg/400x400cc.jpg",
	// 		"song_name": "Girl of My Dreams - Rod Wave"
	// 	  }

	// select one of the song from the output.
	songName := matchingSongsResponse.Path("response.1.song_name").Data().(string)
	downloadresponse, err := client.DownloadMusic(context.Background(), &omnicron.MusicRequest{
		Song: songName,
	})
	if err != nil {
		fmt.Printf("Error downloading song: %v\n", err)
		return
	}
	fmt.Printf("Selected song: %s\n", songName)
	// Marshal the response to JSON
	jsonData, err := json.MarshalIndent(downloadresponse, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	// Write JSON data to a file
	file, err := os.Create("file.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing JSON to file: %v\n", err)
		return
	}

	fmt.Println("JSON data written to file.json successfully")

}
