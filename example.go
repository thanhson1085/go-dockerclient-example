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
	top, _ := client.TopContainer("3ce0eefcced2", "aux")
	fmt.Println(top.Titles)
	for _, p := range top.Processes {
		fmt.Println(p)
	}
}
