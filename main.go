package main

import (
	"fmt"
	"os"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)


func Run(rl *fn.ResourceList) (bool, error) {
	for _, o := range rl.Items{
		fmt.Println(o.GetAPIVersion(), o.GetKind())

		if o.GetAPIVersion() == "req.nephio.org/v1alpha1" && o.GetKind() == "Interface" {
			networkInstanceName, ok, err := o.NestedString("spec", "networkInstance", "name")
			fmt.Println(networkInstanceName)
		}
	}
	




	// fmt.Println("resourcelist", rl)
	// return true, nil
}

func main() {
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}