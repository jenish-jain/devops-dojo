// I AM NOT DONE
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Deploying a pod using kubectl")

	// Check if kubectl is available
	if _, err := exec.LookPath("kubectl"); err != nil {
		fmt.Println("Error: kubectl is not installed or not in PATH")
		fmt.Println("Please install kubectl to run this exercise")
		fmt.Println("Visit: https://kubernetes.io/docs/tasks/tools/")
		os.Exit(1)
	}

	// Check if YAML file exists
	yamlPath := "exercises/01_create_pod/pod.yaml"
	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist\n", yamlPath)
		os.Exit(1)
	}

	app := "kubectl"
	arg0 := "apply"
	arg1 := "-f"
	arg2 := yamlPath

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error executing kubectl:")
		fmt.Println(string(stdout))
		os.Exit(1)
	}

	fmt.Println("Success!")
	fmt.Println(string(stdout))
}
