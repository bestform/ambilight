package main

import (
	"flag"
	"log"
	"os"

	"github.com/bestform/imagecolor"

	"github.com/bestform/gohue"

	"github.com/fsnotify/fsnotify"
)

func main() {
	pathPtr := flag.String("path", "", "The path to the WoW screenshot folder")
	usernamePtr := flag.String("username", "", "The username for your phillips hue network")
	ipPtr := flag.String("ip", "", "The ip for your philips hue network")
	deleteAfterProcessing := flag.Bool("delete", false, "Delete the image after it has been processed")
	flag.Parse()

	if "" == *pathPtr {
		log.Fatal("Please provide a path to watch")
	}
	if "" == *usernamePtr {
		log.Fatal("Please provide a username")
	}
	if "" == *ipPtr {
		log.Fatal("Please provide an ip")
	}
	if _, err := os.Stat(*pathPtr); err != nil {
		log.Fatalf("Path %s does not exist. Exiting.", *pathPtr)
	}

	log.Println("Watching ", *pathPtr)

	client := gohue.NewClient(*usernamePtr, *ipPtr)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("create: ", event.Name)
					sendColorFromFile(event.Name, &client)
					if *deleteAfterProcessing {
						log.Println("deleting: ", event.Name)
						os.Remove(event.Name)
					}
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(*pathPtr)
	if err != nil {
		log.Fatal(err)
	}

	<-done
}

func sendColorFromFile(filename string, client *gohue.Client) {
	i := imagecolor.NewImagecolor(filename)
	r, g, b, err := i.AverageColor()
	if err != nil {
		log.Println("error: ", err)
	}

	for _, light := range (*client).GetLights() {
		light.SetColorRGB(r, g, b)
	}
}
