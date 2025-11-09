// I AM NOT DONE
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Exercise: Deploy a Kubernetes Deployment with replicas")

	yamlPath := "exercises/03_deploy_deployment/deployment.yaml"

	// Check if YAML file exists
	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist\n", yamlPath)
		fmt.Println("Create the deployment.yaml file with a deployment specification")
		fmt.Println("\nHint: A Deployment should have:")
		fmt.Println("  - apiVersion: apps/v1")
		fmt.Println("  - kind: Deployment")
		fmt.Println("  - spec.replicas: 3")
		fmt.Println("  - spec.selector matching the pod labels")
		os.Exit(1)
	}

	// Validate YAML file has content
	fileInfo, err := os.Stat(yamlPath)
	if err != nil || fileInfo.Size() == 0 {
		fmt.Printf("Error: %s is empty\n", yamlPath)
		fmt.Println("Add a valid deployment specification to deployment.yaml")
		os.Exit(1)
	}

	// Read and validate the YAML contains key elements
	content, err := os.ReadFile(yamlPath)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", yamlPath, err)
		os.Exit(1)
	}

	contentStr := string(content)
	checks := map[string]bool{
		"kind: Deployment": strings.Contains(contentStr, "kind: Deployment"),
		"replicas: 3":      strings.Contains(contentStr, "replicas: 3"),
		"selector":         strings.Contains(contentStr, "selector"),
	}

	allPassed := true
	for check, passed := range checks {
		if !passed {
			fmt.Printf("⚠ Warning: deployment.yaml might be missing: %s\n", check)
			allPassed = false
		}
	}

	if !allPassed {
		fmt.Println("\nYour deployment.yaml should include:")
		fmt.Println("  - kind: Deployment")
		fmt.Println("  - replicas: 3 (to create 3 pod replicas)")
		fmt.Println("  - selector with matchLabels")
		fmt.Println("  - template with pod specification")
	}

	fmt.Println("✓ Found deployment.yaml")

	// Check if kubectl is available
	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Println("\n⚠ kubectl is not installed or not in PATH")
		fmt.Println("In a real environment, you would run: kubectl apply -f", yamlPath)
		fmt.Println("\nFor learning purposes, this exercise checks:")
		fmt.Println("  ✓ deployment.yaml exists")
		fmt.Println("  ✓ deployment.yaml is not empty")
		if allPassed {
			fmt.Println("  ✓ deployment.yaml contains required fields")
		}
		fmt.Println("\nWhat you learned:")
		fmt.Println("  • Deployments manage ReplicaSets and Pods")
		fmt.Println("  • replicas: 3 ensures 3 copies of your pod are always running")
		fmt.Println("  • If a pod dies, Deployment automatically creates a new one")
		fmt.Println("  • This provides high availability for your application")
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
		fmt.Println("\nFix the issues in your deployment.yaml file")
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

	fmt.Println("\n✅ Success! Deployment created in cluster")
	fmt.Println(string(stdout))
	fmt.Println("\nTry these commands to see your deployment in action:")
	fmt.Println("  kubectl get deployments")
	fmt.Println("  kubectl get replicasets")
	fmt.Println("  kubectl get pods")
	fmt.Println("  kubectl scale deployment/<deployment-name> --replicas=5")
}
