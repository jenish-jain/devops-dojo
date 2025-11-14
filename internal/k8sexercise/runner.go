package k8sexercise

import (
	"fmt"
	"os"
	"os/exec"
)

// ContentValidator is an optional interface for custom YAML content validation
type ContentValidator interface {
	Validate(content []byte) error
}

// Runner handles the common kubectl exercise workflow
type Runner struct {
	YAMLPath      string
	ResourceType  string // e.g., "Pod", "Service", "Deployment"
	Validator     ContentValidator
}

// NewRunner creates a new K8s exercise runner
func NewRunner(yamlPath, resourceType string, validator ContentValidator) *Runner {
	return &Runner{
		YAMLPath:     yamlPath,
		ResourceType: resourceType,
		Validator:    validator,
	}
}

// Run executes the complete exercise workflow
func (r *Runner) Run() error {
	// Step 1: Validate YAML file exists
	if err := r.validateFileExists(); err != nil {
		return err
	}

	// Step 2: Validate YAML file is not empty
	content, err := r.validateFileNotEmpty()
	if err != nil {
		return err
	}

	// Step 3: Run custom content validation if provided
	if r.Validator != nil {
		if err := r.Validator.Validate(content); err != nil {
			return err
		}
	}

	// Step 4: Check kubectl availability
	kubectlPath, available := r.checkKubectlAvailable()
	if !available {
		r.printDryRunSuccess()
		return nil
	}

	fmt.Printf("✓ Found kubectl at: %s\n", kubectlPath)

	// Step 5: Run kubectl dry-run validation
	if err := r.runDryRun(); err != nil {
		return err
	}

	// Step 6: Apply to cluster
	r.applyToCluster()

	return nil
}

// validateFileExists checks if the YAML file exists
func (r *Runner) validateFileExists() error {
	if _, err := os.Stat(r.YAMLPath); os.IsNotExist(err) {
		fmt.Printf("❌ Error: %s does not exist\n", r.YAMLPath)
		fmt.Printf("Create the %s.yaml file to define your %s\n", r.ResourceType, r.ResourceType)
		return fmt.Errorf("file does not exist: %s", r.YAMLPath)
	}
	return nil
}

// validateFileNotEmpty checks if the YAML file has content
func (r *Runner) validateFileNotEmpty() ([]byte, error) {
	content, err := os.ReadFile(r.YAMLPath)
	if err != nil {
		fmt.Printf("❌ Error reading %s: %v\n", r.YAMLPath, err)
		return nil, err
	}

	if len(content) == 0 {
		fmt.Printf("❌ Error: %s is empty\n", r.YAMLPath)
		fmt.Printf("Add a valid %s specification to the file\n", r.ResourceType)
		return nil, fmt.Errorf("file is empty: %s", r.YAMLPath)
	}

	return content, nil
}

// checkKubectlAvailable checks if kubectl is installed
func (r *Runner) checkKubectlAvailable() (string, bool) {
	kubectlPath, err := exec.LookPath("kubectl")
	return kubectlPath, err == nil
}

// printDryRunSuccess prints success message when kubectl is not available
func (r *Runner) printDryRunSuccess() {
	fmt.Println("\n⚠ kubectl is not installed or not in PATH")
	fmt.Printf("In a real environment, you would run: kubectl apply -f %s\n", r.YAMLPath)
	fmt.Println("\nFor learning purposes, this exercise checks:")
	fmt.Printf("  ✓ %s.yaml exists\n", r.ResourceType)
	fmt.Printf("  ✓ %s.yaml is not empty\n", r.ResourceType)
	if r.Validator != nil {
		fmt.Println("  ✓ YAML content validation passed")
	}
	fmt.Println("\nTo complete this exercise in a real cluster:")
	fmt.Println("  1. Install kubectl: https://kubernetes.io/docs/tasks/tools/")
	fmt.Println("  2. Setup local cluster: minikube start OR kind create cluster")
	fmt.Println("  3. Run this exercise again")
	fmt.Println("\n✅ Exercise validated successfully (kubectl not available - dry run mode)")
}

// runDryRun validates YAML with kubectl dry-run
func (r *Runner) runDryRun() error {
	cmd := exec.Command("kubectl", "apply", "-f", r.YAMLPath, "--dry-run=client")
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("\n❌ Error validating YAML with kubectl:")
		fmt.Println(string(stdout))
		fmt.Printf("\nFix the issues in your %s.yaml file\n", r.ResourceType)
		return fmt.Errorf("kubectl dry-run failed")
	}

	fmt.Println("\n✓ YAML validation passed")
	fmt.Println(string(stdout))
	return nil
}

// applyToCluster attempts to apply the manifest to the cluster
func (r *Runner) applyToCluster() {
	fmt.Println("\nAttempting to apply to cluster...")
	cmd := exec.Command("kubectl", "apply", "-f", r.YAMLPath)
	stdout, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("\n⚠ Could not apply to cluster (cluster might not be running):")
		fmt.Println(string(stdout))
		fmt.Println("\nYour YAML is valid, but ensure your cluster is running:")
		fmt.Println("  minikube status  OR  kubectl cluster-info")
		fmt.Println("\n✅ Exercise validated successfully (dry-run mode)")
		return
	}

	fmt.Printf("\n✅ Success! %s deployed to cluster\n", r.ResourceType)
	fmt.Println(string(stdout))
}
