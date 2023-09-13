package main

import "fmt"

func main() {
	/* Define an array of size 4 that stores deployment options */
	var DeploymentOptions = []string{"r-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
	for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}
}
