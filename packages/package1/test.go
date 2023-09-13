package package1

import "fmt"

type TestType struct {
	A string
	b bool
}

func (t *TestType) DoOnce() {
	if t.b {
		return
	}

	t.b = true
	fmt.Println("done")
}

func Test() {
	fmt.Println("klappt")
}
