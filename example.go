package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "http://192.168.1.191:12375"
	client, _ := docker.NewClient(endpoint)
	// list of images
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		fmt.Println(img.ID, img.RepoTags)
	}
	// top
	top, _ := client.TopContainer("fc5bfefc175c", "aux")
	fmt.Println(top.Titles)
	for _, p := range top.Processes {
		fmt.Println(p)
	}
	statsOption := docker.StatsOptions{}
	stats := make(chan *docker.Stats)
	statsOption.ID = "fc5bfefc175c"
	statsOption.Stream = true
	statsOption.Stats = stats
	done := make(chan bool)
	statsOption.Done = done
	errC := make(chan error, 1)
	go func() {
		errC <- client.Stats(statsOption)
		close(errC)
	}()
	for {
		s, ok := <-stats
		if !ok {
			break
		}
		fmt.Println("Network Stats", s.Network)
		fmt.Println("Memory Stats", s.MemoryStats)
	}
}
