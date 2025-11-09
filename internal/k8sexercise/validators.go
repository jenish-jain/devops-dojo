package k8sexercise

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// NoOpValidator is a validator that does nothing (for simple exercises)
type NoOpValidator struct{}

func (v *NoOpValidator) Validate(content []byte) error {
	return nil
}

// DeploymentValidator validates Deployment-specific requirements
type DeploymentValidator struct{}

func (v *DeploymentValidator) Validate(content []byte) error {
	contentStr := string(content)

	// Check for Deployment kind
	if !strings.Contains(contentStr, "kind: Deployment") {
		fmt.Println("\n⚠ Warning: Expected 'kind: Deployment' in the YAML")
		fmt.Println("Make sure you're creating a Deployment resource")
	}

	// Check for replicas
	if !strings.Contains(contentStr, "replicas:") {
		fmt.Println("\n⚠ Warning: No 'replicas' field found")
		fmt.Println("Hint: Your Deployment should specify replicas: 3")
	} else if !strings.Contains(contentStr, "replicas: 3") {
		fmt.Println("\n⚠ Warning: Expected 'replicas: 3'")
	}

	// Check for selector
	if !strings.Contains(contentStr, "selector:") {
		fmt.Println("\n⚠ Warning: No 'selector' field found")
		fmt.Println("Deployments require a selector to match pods")
	}

	return nil
}

// ConfigMapValidator validates ConfigMap-specific requirements
type ConfigMapValidator struct{}

func (v *ConfigMapValidator) Validate(content []byte) error {
	contentStr := string(content)

	// Check for ConfigMap kind
	if !strings.Contains(contentStr, "kind: ConfigMap") {
		fmt.Println("\n⚠ Warning: Expected 'kind: ConfigMap' in the YAML")
		fmt.Println("Make sure you're creating a ConfigMap resource")
	}

	// Check for data section
	if !strings.Contains(contentStr, "data:") {
		fmt.Println("\n⚠ Warning: No 'data:' field found")
		fmt.Println("ConfigMaps store configuration data in the 'data:' section")
		fmt.Println("Example:")
		fmt.Println("  data:")
		fmt.Println("    key: value")
	}

	return nil
}

// SecretValidator validates Secret-specific requirements
type SecretValidator struct{}

func (v *SecretValidator) Validate(content []byte) error {
	contentStr := string(content)

	// Check for Secret kind
	if !strings.Contains(contentStr, "kind: Secret") {
		fmt.Println("\n⚠ Warning: Expected 'kind: Secret' in the YAML")
		fmt.Println("Make sure you're creating a Secret resource")
	}

	// Check for type
	if !strings.Contains(contentStr, "type:") {
		fmt.Println("\n⚠ Warning: No 'type' field found")
		fmt.Println("Secrets should specify a type (e.g., 'type: Opaque')")
	}

	// Check for data section
	if !strings.Contains(contentStr, "data:") {
		fmt.Println("\n⚠ Warning: No 'data:' field found")
		fmt.Println("Secrets store sensitive data in the 'data:' section")
		return nil
	}

	// Check if data is base64 encoded
	lines := strings.Split(contentStr, "\n")
	hasBase64 := false
	inDataSection := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "data:") {
			inDataSection = true
			continue
		}

		// Exit data section if we hit another top-level key
		if inDataSection && !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "\t") {
			break
		}

		// Check for key-value pairs in data section
		if inDataSection && strings.Contains(trimmed, ":") && !strings.HasPrefix(trimmed, "#") {
			parts := strings.SplitN(trimmed, ":", 2)
			if len(parts) == 2 {
				value := strings.TrimSpace(parts[1])
				// Try to decode as base64
				if _, err := base64.StdEncoding.DecodeString(value); err == nil && len(value) > 0 {
					hasBase64 = true
					break
				}
			}
		}
	}

	if !hasBase64 {
		fmt.Println("\n⚠ Warning: Secret data should be base64-encoded")
		fmt.Println("Use: echo -n 'your-value' | base64")
		fmt.Println("Example:")
		fmt.Println("  data:")
		fmt.Println("    password: bXlwYXNzd29yZA==")
	}

	return nil
}
