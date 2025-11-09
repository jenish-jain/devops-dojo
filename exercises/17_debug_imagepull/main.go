// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Debug ImagePullBackOff Error")
	fmt.Println("\n🔍 Debugging Scenario:")
	fmt.Println("A pod can't start because it can't pull the container image. Your task is to:")
	fmt.Println("  1. Identify what's wrong with the image specification")
	fmt.Println("  2. Fix the image name in the YAML file")
	fmt.Println("")

	runner := k8sexercise.NewRunner(
		"exercises/17_debug_imagepull/imagepull-pod.yaml",
		"Pod",
		&k8sexercise.ImagePullValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\n📚 What you learned:")
	fmt.Println("  • ImagePullBackOff means Kubernetes can't pull the container image")
	fmt.Println("  • Common causes: typo in image name, wrong tag, missing registry auth")
	fmt.Println("  • Always specify explicit image tags (e.g., nginx:1.21, not nginx:latest)")
	fmt.Println("  • Check image exists: docker pull <image-name>")
	fmt.Println("  • For private registries, create imagePullSecrets")
	fmt.Println("\n🔧 Debugging commands:")
	fmt.Println("  kubectl get pods  # Check pod status")
	fmt.Println("  kubectl describe pod <pod-name>  # See detailed error message")
	fmt.Println("  kubectl get events  # See recent cluster events")
	fmt.Println("  docker search <image-name>  # Verify image exists")
}
