package main

import (
	"fmt"
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)


func Run(rl *fn.ResourceList) (bool, error) {
	for _, o := range rl.Items{
		fmt.Println(o.GetAPIVersion(), o.GetKind())
	}
	




	// fmt.Println("resourcelist", rl)
	// return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}