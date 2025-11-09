// I AM NOT DONE
package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Exercise: Create a Secret for sensitive data")

	yamlPath := "exercises/05_create_secret/secret.yaml"

	// Check if YAML file exists
	if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist\n", yamlPath)
		fmt.Println("Create the secret.yaml file with sensitive data")
		fmt.Println("\nHint: A Secret should have:")
		fmt.Println("  - apiVersion: v1")
		fmt.Println("  - kind: Secret")
		fmt.Println("  - type: Opaque (for generic secrets)")
		fmt.Println("  - data: base64-encoded key-value pairs")
		os.Exit(1)
	}

	// Validate YAML file has content
	fileInfo, err := os.Stat(yamlPath)
	if err != nil || fileInfo.Size() == 0 {
		fmt.Printf("Error: %s is empty\n", yamlPath)
		fmt.Println("Add a valid Secret specification to secret.yaml")
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
		"kind: Secret": strings.Contains(contentStr, "kind: Secret"),
		"type:":        strings.Contains(contentStr, "type:"),
		"data:":        strings.Contains(contentStr, "data:"),
	}

	allPassed := true
	for check, passed := range checks {
		if !passed {
			fmt.Printf("⚠ Warning: secret.yaml might be missing: %s\n", check)
			allPassed = false
		}
	}

	// Check if data is base64 encoded
	if strings.Contains(contentStr, "data:") {
		lines := strings.Split(contentStr, "\n")
		hasBase64 := false
		for _, line := range lines {
			if strings.Contains(line, ":") && !strings.Contains(line, "kind:") &&
				!strings.Contains(line, "type:") && !strings.Contains(line, "data:") &&
				!strings.Contains(line, "apiVersion:") && !strings.Contains(line, "name:") &&
				strings.TrimSpace(line) != "" && !strings.HasPrefix(strings.TrimSpace(line), "#") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					value := strings.TrimSpace(parts[1])
					value = strings.Trim(value, "\"")
					if value != "" {
						_, err := base64.StdEncoding.DecodeString(value)
						if err == nil && len(value) > 0 {
							hasBase64 = true
							break
						}
					}
				}
			}
		}
		if !hasBase64 {
			fmt.Println("\n⚠ Warning: Secret data should be base64-encoded")
			fmt.Println("Example: echo -n 'mypassword' | base64")
			fmt.Println("Result: bXlwYXNzd29yZA==")
		}
	}

	if !allPassed {
		fmt.Println("\nYour secret.yaml should include:")
		fmt.Println("  - kind: Secret")
		fmt.Println("  - type: Opaque")
		fmt.Println("  - data: with base64-encoded values")
		fmt.Println("\nExample:")
		fmt.Println("  data:")
		fmt.Println("    password: bXlwYXNzd29yZA==  # base64 encoded")
	}

	fmt.Println("✓ Found secret.yaml")

	// Check if kubectl is available
	kubectlPath, err := exec.LookPath("kubectl")
	if err != nil {
		fmt.Println("\n⚠ kubectl is not installed or not in PATH")
		fmt.Println("In a real environment, you would run: kubectl apply -f", yamlPath)
		fmt.Println("\nFor learning purposes, this exercise checks:")
		fmt.Println("  ✓ secret.yaml exists")
		fmt.Println("  ✓ secret.yaml is not empty")
		if allPassed {
			fmt.Println("  ✓ secret.yaml contains required fields")
		}
		fmt.Println("\nWhat you learned:")
		fmt.Println("  • Secrets store sensitive data (passwords, tokens, keys)")
		fmt.Println("  • Data is base64-encoded (not encrypted by default)")
		fmt.Println("  • Never commit secrets to git!")
		fmt.Println("  • Use Secrets for sensitive data, ConfigMaps for non-sensitive")
		fmt.Println("  • In production, use encryption at rest and RBAC")
		fmt.Println("\n💡 Security Tip: In real environments, use tools like:")
		fmt.Println("  - Sealed Secrets")
		fmt.Println("  - External Secrets Operator")
		fmt.Println("  - HashiCorp Vault")
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
		fmt.Println("\nFix the issues in your secret.yaml file")
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

	fmt.Println("\n✅ Success! Secret created in cluster")
	fmt.Println(string(stdout))
	fmt.Println("\nTry these commands to explore your Secret:")
	fmt.Println("  kubectl get secrets")
	fmt.Println("  kubectl describe secret/<secret-name>")
	fmt.Println("  kubectl get secret/<secret-name> -o yaml")
	fmt.Println("\n⚠️ Remember: Secrets in git are NOT secret! Use proper secret management.")
}
