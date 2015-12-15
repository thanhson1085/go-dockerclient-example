package main

import (
	"fmt"
	"github.com/thanhson1085/go-dockerclient"
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
	top, _ := client.TopContainer("7d1e304579e4", "aux")
	fmt.Println(top.Titles)
	for _, p := range top.Processes {
		fmt.Println(p)
	}
	statsOption := docker.StatsOptions{}
	stats := make(chan *docker.Stats)
	statsOption.ID = "7d1e304579e4"
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
		fmt.Println("Network Stats", s.Networks["eth0"])
		fmt.Println("Memory Stats", s.MemoryStats)
	}
}
