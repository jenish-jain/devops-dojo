// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create a NetworkPolicy for pod-level network security")

	runner := k8sexercise.NewRunner(
		"exercises/11_network_policies/networkpolicy.yaml",
		"NetworkPolicy",
		&k8sexercise.NetworkPolicyValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • NetworkPolicies control traffic between pods")
	fmt.Println("  • Act as a firewall at the pod level")
	fmt.Println("  • Define ingress (incoming) and egress (outgoing) rules")
	fmt.Println("  • Use label selectors to target specific pods")
	fmt.Println("  • Require a CNI plugin that supports NetworkPolicy (Calico, Cilium, etc.)")
	fmt.Println("  • Default deny all unless explicitly allowed")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get networkpolicies")
	fmt.Println("  kubectl describe networkpolicy <policy-name>")
}
