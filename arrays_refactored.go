package main

import "fmt"

func main() {
	/* Define an array and let the compiler count its size */
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}
}
