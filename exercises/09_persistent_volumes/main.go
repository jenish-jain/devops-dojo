// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create a PersistentVolumeClaim for persistent storage")

	runner := k8sexercise.NewRunner(
		"exercises/09_persistent_volumes/pvc.yaml",
		"PersistentVolumeClaim",
		&k8sexercise.PersistentVolumeClaimValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • PersistentVolumes (PV) provide storage in the cluster")
	fmt.Println("  • PersistentVolumeClaims (PVC) request storage from PVs")
	fmt.Println("  • PVCs decouple storage from pod lifecycle")
	fmt.Println("  • Data persists even if pods are deleted")
	fmt.Println("  • Access modes: ReadWriteOnce (RWO), ReadOnlyMany (ROX), ReadWriteMany (RWX)")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get pvc")
	fmt.Println("  kubectl get pv")
	fmt.Println("  kubectl describe pvc <pvc-name>")
}
