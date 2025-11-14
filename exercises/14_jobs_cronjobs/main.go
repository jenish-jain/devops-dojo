// I AM NOT DONE
package main

import (
	"fmt"
	"os"

	"github.com/jenish-jain/devops-dojo/internal/k8sexercise"
)

func main() {
	fmt.Println("Exercise: Create a Job or CronJob for batch workloads")

	runner := k8sexercise.NewRunner(
		"exercises/14_jobs_cronjobs/job.yaml",
		"Job",
		&k8sexercise.JobValidator{},
	)

	if err := runner.Run(); err != nil {
		os.Exit(1)
	}

	// Additional educational content
	fmt.Println("\nWhat you learned:")
	fmt.Println("  • Jobs run tasks to completion (one-time batch work)")
	fmt.Println("  • CronJobs run Jobs on a schedule (recurring tasks)")
	fmt.Println("  • Jobs create pods that run until successful completion")
	fmt.Println("  • Use restartPolicy: OnFailure or Never for Jobs")
	fmt.Println("  • CronJob schedule uses standard cron format")
	fmt.Println("  • Examples: backups, reports, data processing, cleanup")
	fmt.Println("\nTry these commands:")
	fmt.Println("  kubectl get jobs")
	fmt.Println("  kubectl get cronjobs")
	fmt.Println("  kubectl logs job/<job-name>")
	fmt.Println("  kubectl describe cronjob <cronjob-name>")
}
