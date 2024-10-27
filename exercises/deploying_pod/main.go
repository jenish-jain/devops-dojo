// I AM NOT DONE
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Deploying a pod")

	app := "kubectl"

	arg0 := "apply"
	arg1 := "-f"
	arg2 := "exercises/deploying_pod/pod.yaml"

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("Command Successfully Executed")
	fmt.Println(string(stdout))
	fmt.Println("pod deployed")
}
