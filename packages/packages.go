package main

// go mod init packages or github.com/marceljaworski/packages
import "packages/package1"

func main() {
	package1.Test()
	x := package1.TestType{}
	x.DoOnce()
	// x.b = false
	x.DoOnce()
}
