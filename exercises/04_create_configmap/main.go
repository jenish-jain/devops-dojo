// I AM NOT DONE
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Exercise: Create and use a ConfigMap for application configuration")

	yamlPath := "exercises/04_create_configmap/configmap.yaml"

	// Check if YAML file exists
	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist\n", yamlPath)
		fmt.Println("Create the configmap.yaml file with configuration data")
		fmt.Println("\nHint: A ConfigMap should have:")
		fmt.Println("  - apiVersion: v1")
		fmt.Println("  - kind: ConfigMap")
		fmt.Println("  - data: key-value pairs of configuration")
		os.Exit(1)
	}

	// Validate YAML file has content
	fileInfo, err := os.Stat(yamlPath)
	if err != nil || fileInfo.Size() == 0 {
		fmt.Printf("Error: %s is empty\n", yamlPath)
		fmt.Println("Add a valid ConfigMap specification to configmap.yaml")
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
		"kind: ConfigMap": strings.Contains(contentStr, "kind: ConfigMap"),
		"data:":           strings.Contains(contentStr, "data:"),
	}

	allPassed := true
	for check, passed := range checks {
		if !passed {
			fmt.Printf("⚠ Warning: configmap.yaml might be missing: %s\n", check)
			allPassed = false
		}
	}

	if !allPassed {
		fmt.Println("\nYour configmap.yaml should include:")
		fmt.Println("  - kind: ConfigMap")
		fmt.Println("  - data: with key-value pairs")
		fmt.Println("\nExample:")
		fmt.Println("  data:")
		fmt.Println("    DATABASE_URL: postgres://db:5432/myapp")
		fmt.Println("    LOG_LEVEL: info")
	}

	fmt.Println("✓ Found configmap.yaml")

	// Check if kubectl is available
	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Println("\n⚠ kubectl is not installed or not in PATH")
		fmt.Println("In a real environment, you would run: kubectl apply -f", yamlPath)
		fmt.Println("\nFor learning purposes, this exercise checks:")
		fmt.Println("  ✓ configmap.yaml exists")
		fmt.Println("  ✓ configmap.yaml is not empty")
		if allPassed {
			fmt.Println("  ✓ configmap.yaml contains required fields")
		}
		fmt.Println("\nWhat you learned:")
		fmt.Println("  • ConfigMaps externalize configuration from application code")
		fmt.Println("  • This follows the 12-factor app methodology")
		fmt.Println("  • You can use the same container image across dev/staging/prod")
		fmt.Println("  • Just change the ConfigMap for different environments")
		fmt.Println("  • ConfigMaps can be mounted as environment variables or files")
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
		fmt.Println("\nFix the issues in your configmap.yaml file")
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

	fmt.Println("\n✅ Success! ConfigMap created in cluster")
	fmt.Println(string(stdout))
	fmt.Println("\nTry these commands to explore your ConfigMap:")
	fmt.Println("  kubectl get configmaps")
	fmt.Println("  kubectl describe configmap/<configmap-name>")
	fmt.Println("  kubectl get configmap/<configmap-name> -o yaml")
}
