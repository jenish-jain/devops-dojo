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

// ResourceLimitsValidator validates that resource limits and requests are defined
type ResourceLimitsValidator struct{}

func (v *ResourceLimitsValidator) Validate(content []byte) error {
	contentStr := string(content)

	// Check for resources section
	if !strings.Contains(contentStr, "resources:") {
		fmt.Println("\n⚠ Warning: No 'resources:' field found")
		fmt.Println("Pods should define resource limits and requests")
		fmt.Println("Example:")
		fmt.Println("  resources:")
		fmt.Println("    requests:")
		fmt.Println("      cpu: 100m")
		fmt.Println("      memory: 128Mi")
		fmt.Println("    limits:")
		fmt.Println("      cpu: 200m")
		fmt.Println("      memory: 256Mi")
		return nil
	}

	// Check for limits
	if !strings.Contains(contentStr, "limits:") {
		fmt.Println("\n⚠ Warning: No 'limits:' field found")
		fmt.Println("Resource limits prevent pods from consuming excessive resources")
	}

	// Check for requests
	if !strings.Contains(contentStr, "requests:") {
		fmt.Println("\n⚠ Warning: No 'requests:' field found")
		fmt.Println("Resource requests help Kubernetes schedule pods efficiently")
	}

	// Check for CPU and memory
	if !strings.Contains(contentStr, "cpu:") {
		fmt.Println("\n⚠ Warning: No CPU resource specification found")
	}

	if !strings.Contains(contentStr, "memory:") {
		fmt.Println("\n⚠ Warning: No memory resource specification found")
	}

	return nil
}

// HealthProbesValidator validates that liveness and readiness probes are defined
type HealthProbesValidator struct{}

func (v *HealthProbesValidator) Validate(content []byte) error {
	contentStr := string(content)

	hasLiveness := strings.Contains(contentStr, "livenessProbe:")
	hasReadiness := strings.Contains(contentStr, "readinessProbe:")

	if !hasLiveness && !hasReadiness {
		fmt.Println("\n⚠ Warning: No health probes found")
		fmt.Println("Consider adding liveness and readiness probes for production")
		fmt.Println("Example:")
		fmt.Println("  livenessProbe:")
		fmt.Println("    httpGet:")
		fmt.Println("      path: /healthz")
		fmt.Println("      port: 8080")
		fmt.Println("    initialDelaySeconds: 3")
		fmt.Println("    periodSeconds: 3")
		return nil
	}

	if !hasLiveness {
		fmt.Println("\n⚠ Warning: No 'livenessProbe:' found")
		fmt.Println("Liveness probes restart unhealthy containers")
	}

	if !hasReadiness {
		fmt.Println("\n⚠ Warning: No 'readinessProbe:' found")
		fmt.Println("Readiness probes ensure traffic only goes to ready pods")
	}

	return nil
}

// PersistentVolumeClaimValidator validates PVC configuration
type PersistentVolumeClaimValidator struct{}

func (v *PersistentVolumeClaimValidator) Validate(content []byte) error {
	contentStr := string(content)

	if !strings.Contains(contentStr, "kind: PersistentVolumeClaim") {
		fmt.Println("\n⚠ Warning: Expected 'kind: PersistentVolumeClaim' in the YAML")
	}

	if !strings.Contains(contentStr, "resources:") || !strings.Contains(contentStr, "requests:") {
		fmt.Println("\n⚠ Warning: PVC should specify storage size in resources.requests.storage")
		fmt.Println("Example: storage: 1Gi")
	}

	if !strings.Contains(contentStr, "accessModes:") {
		fmt.Println("\n⚠ Warning: No 'accessModes' found")
		fmt.Println("Common modes: ReadWriteOnce, ReadOnlyMany, ReadWriteMany")
	}

	return nil
}

// IngressValidator validates Ingress configuration
type IngressValidator struct{}

func (v *IngressValidator) Validate(content []byte) error {
	contentStr := string(content)

	if !strings.Contains(contentStr, "kind: Ingress") {
		fmt.Println("\n⚠ Warning: Expected 'kind: Ingress' in the YAML")
	}

	if !strings.Contains(contentStr, "rules:") {
		fmt.Println("\n⚠ Warning: No 'rules' found")
		fmt.Println("Ingress needs routing rules to direct traffic")
	}

	if !strings.Contains(contentStr, "host:") {
		fmt.Println("\n⚠ Warning: Consider adding 'host:' for host-based routing")
	}

	return nil
}

// NetworkPolicyValidator validates NetworkPolicy configuration
type NetworkPolicyValidator struct{}

func (v *NetworkPolicyValidator) Validate(content []byte) error {
	contentStr := string(content)

	if !strings.Contains(contentStr, "kind: NetworkPolicy") {
		fmt.Println("\n⚠ Warning: Expected 'kind: NetworkPolicy' in the YAML")
	}

	if !strings.Contains(contentStr, "podSelector:") {
		fmt.Println("\n⚠ Warning: NetworkPolicy requires 'podSelector' to select target pods")
	}

	hasIngress := strings.Contains(contentStr, "ingress:")
	hasEgress := strings.Contains(contentStr, "egress:")

	if !hasIngress && !hasEgress {
		fmt.Println("\n⚠ Warning: NetworkPolicy should define ingress or egress rules")
	}

	return nil
}

// RBACValidator validates RBAC resources (ServiceAccount, Role, RoleBinding)
type RBACValidator struct{}

func (v *RBACValidator) Validate(content []byte) error {
	contentStr := string(content)

	hasServiceAccount := strings.Contains(contentStr, "kind: ServiceAccount")
	hasRole := strings.Contains(contentStr, "kind: Role") || strings.Contains(contentStr, "kind: ClusterRole")
	hasRoleBinding := strings.Contains(contentStr, "kind: RoleBinding") || strings.Contains(contentStr, "kind: ClusterRoleBinding")

	if !hasServiceAccount && !hasRole && !hasRoleBinding {
		fmt.Println("\n⚠ Warning: Expected RBAC resources (ServiceAccount, Role, or RoleBinding)")
	}

	if hasRole && !strings.Contains(contentStr, "rules:") {
		fmt.Println("\n⚠ Warning: Role should define 'rules' with permissions")
	}

	if hasRoleBinding && !strings.Contains(contentStr, "roleRef:") {
		fmt.Println("\n⚠ Warning: RoleBinding requires 'roleRef' to reference a Role")
	}

	return nil
}

// StatefulSetValidator validates StatefulSet configuration
type StatefulSetValidator struct{}

func (v *StatefulSetValidator) Validate(content []byte) error {
	contentStr := string(content)

	if !strings.Contains(contentStr, "kind: StatefulSet") {
		fmt.Println("\n⚠ Warning: Expected 'kind: StatefulSet' in the YAML")
	}

	if !strings.Contains(contentStr, "serviceName:") {
		fmt.Println("\n⚠ Warning: StatefulSet requires 'serviceName' for stable network identity")
	}

	if !strings.Contains(contentStr, "volumeClaimTemplates:") {
		fmt.Println("\n⚠ Warning: Consider adding 'volumeClaimTemplates' for persistent storage")
	}

	return nil
}

// JobValidator validates Job and CronJob configuration
type JobValidator struct{}

func (v *JobValidator) Validate(content []byte) error {
	contentStr := string(content)

	hasJob := strings.Contains(contentStr, "kind: Job")
	hasCronJob := strings.Contains(contentStr, "kind: CronJob")

	if !hasJob && !hasCronJob {
		fmt.Println("\n⚠ Warning: Expected 'kind: Job' or 'kind: CronJob' in the YAML")
	}

	if hasCronJob && !strings.Contains(contentStr, "schedule:") {
		fmt.Println("\n⚠ Warning: CronJob requires 'schedule' field (cron format)")
		fmt.Println("Example: schedule: '*/5 * * * *'  # Every 5 minutes")
	}

	if hasJob && !strings.Contains(contentStr, "restartPolicy:") {
		fmt.Println("\n⚠ Warning: Jobs should specify 'restartPolicy' (OnFailure or Never)")
	}

	return nil
}

// HPAValidator validates HorizontalPodAutoscaler configuration
type HPAValidator struct{}

func (v *HPAValidator) Validate(content []byte) error {
	contentStr := string(content)

	if !strings.Contains(contentStr, "kind: HorizontalPodAutoscaler") {
		fmt.Println("\n⚠ Warning: Expected 'kind: HorizontalPodAutoscaler' in the YAML")
	}

	if !strings.Contains(contentStr, "scaleTargetRef:") {
		fmt.Println("\n⚠ Warning: HPA requires 'scaleTargetRef' to reference the target resource")
	}

	if !strings.Contains(contentStr, "minReplicas:") {
		fmt.Println("\n⚠ Warning: Consider setting 'minReplicas' for minimum pod count")
	}

	if !strings.Contains(contentStr, "maxReplicas:") {
		fmt.Println("\n⚠ Warning: HPA requires 'maxReplicas' for maximum pod count")
	}

	hasMetrics := strings.Contains(contentStr, "metrics:") || strings.Contains(contentStr, "targetCPUUtilizationPercentage:")
	if !hasMetrics {
		fmt.Println("\n⚠ Warning: HPA should define metrics for autoscaling decisions")
	}

	return nil
}

// CrashLoopValidator validates that the pod won't crash
type CrashLoopValidator struct{}

func (v *CrashLoopValidator) Validate(content []byte) error {
	contentStr := string(content)

	// Check for common crash causes
	if strings.Contains(contentStr, "command:") {
		if strings.Contains(contentStr, "exit 1") || strings.Contains(contentStr, "false") {
			fmt.Println("\n⚠ Warning: Command contains 'exit 1' or 'false' which will cause the pod to crash")
			fmt.Println("Fix: Change to a command that runs successfully")
		}
	}

	if !strings.Contains(contentStr, "restartPolicy:") {
		fmt.Println("\n💡 Tip: Consider setting 'restartPolicy: Always' (default) for regular pods")
	}

	return nil
}

// ImagePullValidator validates image configuration
type ImagePullValidator struct{}

func (v *ImagePullValidator) Validate(content []byte) error {
	contentStr := string(content)

	// Check for common image issues
	if strings.Contains(contentStr, "image:") {
		// Look for obviously wrong image names
		if strings.Contains(contentStr, "ngimx") || strings.Contains(contentStr, "ngin:") {
			fmt.Println("\n⚠ Warning: Image name appears to have a typo")
			fmt.Println("Common images: nginx, busybox, redis, postgres")
		}

		// Check for latest tag
		lines := strings.Split(contentStr, "\n")
		for _, line := range lines {
			if strings.Contains(line, "image:") && !strings.Contains(line, ":") {
				fmt.Println("\n💡 Tip: Always specify image tags explicitly (e.g., nginx:1.21) instead of using implicit :latest")
				break
			}
		}
	}

	return nil
}

// PendingPodValidator validates resource requests
type PendingPodValidator struct{}

func (v *PendingPodValidator) Validate(content []byte) error {
	contentStr := string(content)

	// Check for unreasonable resource requests
	if strings.Contains(contentStr, "cpu:") {
		if strings.Contains(contentStr, "1000") || strings.Contains(contentStr, "100\"") {
			fmt.Println("\n⚠ Warning: CPU request seems very high (e.g., 1000 cores)")
			fmt.Println("Typical values: 100m (0.1 core), 500m (0.5 core), 1 (1 core)")
		}
	}

	if strings.Contains(contentStr, "memory:") {
		if strings.Contains(contentStr, "1000Gi") || strings.Contains(contentStr, "100Ti") {
			fmt.Println("\n⚠ Warning: Memory request seems unreasonably high")
			fmt.Println("Typical values: 128Mi, 256Mi, 512Mi, 1Gi")
		}
	}

	// Check for node selectors
	if strings.Contains(contentStr, "nodeSelector:") {
		fmt.Println("\n💡 Tip: Ensure nodeSelector labels match available nodes")
		fmt.Println("Use: kubectl get nodes --show-labels")
	}

	return nil
}

// ServiceDebugValidator validates service configuration
type ServiceDebugValidator struct{}

func (v *ServiceDebugValidator) Validate(content []byte) error {
	contentStr := string(content)

	if !strings.Contains(contentStr, "kind: Service") {
		return nil
	}

	// Check for selector
	if !strings.Contains(contentStr, "selector:") {
		fmt.Println("\n⚠ Warning: Service has no selector!")
		fmt.Println("Service selector must match pod labels")
		return nil
	}

	// Check for common selector mismatches
	if strings.Contains(contentStr, "selector:") {
		fmt.Println("\n💡 Debugging tip: Verify selector matches pod labels")
		fmt.Println("  1. Check service selector: kubectl describe svc <service-name>")
		fmt.Println("  2. Check pod labels: kubectl get pods --show-labels")
		fmt.Println("  3. Check endpoints: kubectl get endpoints <service-name>")
	}

	// Check for port configuration
	if !strings.Contains(contentStr, "port:") {
		fmt.Println("\n⚠ Warning: Service should define 'port'")
	}

	if !strings.Contains(contentStr, "targetPort:") {
		fmt.Println("\n⚠ Warning: Consider specifying 'targetPort' explicitly")
	}

	return nil
}
