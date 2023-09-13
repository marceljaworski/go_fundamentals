package main

import "packages/package1"

func main() {
	package1.Test()
	x := package1.TestType{}
	x.DoOnce()
	x.DoOnce()
}
