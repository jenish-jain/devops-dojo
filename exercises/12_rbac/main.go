// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create RBAC resources (ServiceAccount, Role, RoleBinding)")

	runner := k8sexercise.NewRunner(
		"exercises/12_rbac/rbac.yaml",
		"RBAC",
		&k8sexercise.RBACValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • RBAC (Role-Based Access Control) manages permissions in Kubernetes")
	fmt.Println("  • ServiceAccount: Identity for processes running in pods")
	fmt.Println("  • Role: Defines permissions (verbs + resources)")
	fmt.Println("  • RoleBinding: Grants permissions to ServiceAccounts/Users")
	fmt.Println("  • ClusterRole/ClusterRoleBinding: Cluster-wide permissions")
	fmt.Println("  • Principle of least privilege: Grant minimal necessary permissions")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get serviceaccounts")
	fmt.Println("  kubectl get roles")
	fmt.Println("  kubectl get rolebindings")
	fmt.Println("  kubectl describe role <role-name>")
}
