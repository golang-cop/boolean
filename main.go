package main

import (
	"fmt"

	Boolean "github.com/golang-cop/boolean/src"
)

func main() {
	boolean := Boolean.True()
	//fmt.Println(boolean.ToGoString())
	//*
	isEqual := boolean.Equal(Boolean.False())
	inspect := isEqual.Inspect()
	fmt.Println(
		"Boolean.True().Equal(Boolean.False()).Inspect().ToGoString(): ",
		inspect.ToGoString(),
	)
	//*/
}
