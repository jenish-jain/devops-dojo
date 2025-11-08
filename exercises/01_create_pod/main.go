// I AM NOT DONE
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Deploying a pod using kubectl")

	yamlPath := "exercises/01_create_pod/pod.yaml"

	// Check if YAML file exists
	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist\n", yamlPath)
		fmt.Println("Create the pod.yaml file with a proper pod definition")
		os.Exit(1)
	}

	// Validate YAML file has content
	fileInfo, err := os.Stat(yamlPath)
	if err != nil || fileInfo.Size() == 0 {
		fmt.Printf("Error: %s is empty\n", yamlPath)
		fmt.Println("Add a valid pod specification to pod.yaml")
		os.Exit(1)
	}

	fmt.Println("✓ Found pod.yaml")

	// Check if kubectl is available
	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Println("\n⚠ kubectl is not installed or not in PATH")
		fmt.Println("In a real environment, you would run: kubectl apply -f", yamlPath)
		fmt.Println("\nFor learning purposes, this exercise checks:")
		fmt.Println("  ✓ pod.yaml exists")
		fmt.Println("  ✓ pod.yaml is not empty")
		fmt.Println("\nTo complete this exercise in a real cluster:")
		fmt.Println("  1. Install kubectl: https://kubernetes.io/docs/tasks/tools/")
		fmt.Println("  2. Setup local cluster: minikube start OR kind create cluster")
		fmt.Println("  3. Run this exercise again")
		fmt.Println("\n✅ Exercise validated successfully (kubectl not available - dry run mode)")
		return
	}

	// kubectl is available, try to apply
	fmt.Printf("✓ Found kubectl at: %s\n", kubectlPath)

	cmd := exec.Command("kubectl", "apply", "-f", yamlPath, "--dry-run=client")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("\n❌ Error validating with kubectl:")
		fmt.Println(string(stdout))
		fmt.Println("\nFix the issues in your pod.yaml file")
		os.Exit(1)
	}

	fmt.Println("\n✓ YAML validation passed")
	fmt.Println(string(stdout))

	// Try to actually apply if cluster is available
	fmt.Println("\nAttempting to apply to cluster...")
	cmd = exec.Command("kubectl", "apply", "-f", yamlPath)
	stdout, err = cmd.CombinedOutput()

	if err != nil {
		fmt.Println("\n⚠ Could not apply to cluster (cluster might not be running):")
		fmt.Println(string(stdout))
		fmt.Println("\nYour YAML is valid, but ensure your cluster is running:")
		fmt.Println("  minikube status  OR  kubectl cluster-info")
		fmt.Println("\n✅ Exercise validated successfully (dry-run mode)")
		return
	}

	fmt.Println("\n✅ Success! Pod deployed to cluster")
	fmt.Println(string(stdout))
}
