# DevOps Dojo - Exercise Roadmap

This document outlines suggested exercises for the DevOps Dojo learning tool. These exercises are designed to teach fundamental DevOps skills that interns encounter in real-world scenarios, all executable locally without paid cloud services.

## Exercise Categories

### 1. Kubernetes Fundamentals (Beginner)

#### 03_deploy_deployment
**Learning Objective:** Understand Kubernetes Deployments and ReplicaSets
- Create a deployment with 3 replicas
- Learn about desired state vs actual state
- Practice scaling applications
- **Real-world scenario:** Deploying a web application with multiple instances

#### 04_create_configmap
**Learning Objective:** Externalize configuration from application code
- Create a ConfigMap with application settings
- Mount ConfigMap as environment variables
- Update pod to use the ConfigMap
- **Real-world scenario:** Managing different configs for dev/staging/prod

#### 05_create_secret
**Learning Objective:** Handle sensitive data securely
- Create a Secret for database credentials
- Mount Secret as environment variables
- Understand base64 encoding
- **Real-world scenario:** Storing API keys and passwords

#### 06_pod_labels_selectors
**Learning Objective:** Organize and query resources using labels
- Add labels to pods
- Use selectors to filter resources
- Understand label-based grouping
- **Real-world scenario:** Organizing microservices by team, environment, version

#### 07_resource_limits
**Learning Objective:** Manage resource allocation
- Set CPU and memory requests
- Set CPU and memory limits
- Understand resource quotas
- **Real-world scenario:** Preventing resource starvation in shared clusters

### 2. Kubernetes Services & Networking (Beginner-Intermediate)

#### 08_nodeport_service
**Learning Objective:** Expose applications outside the cluster
- Create a NodePort service
- Access application from outside cluster
- Understand port mapping
- **Real-world scenario:** Making your app accessible for testing

#### 09_multi_container_pod
**Learning Objective:** Deploy sidecar containers
- Create pod with main app + logging sidecar
- Understand shared volumes between containers
- Learn sidecar pattern
- **Real-world scenario:** Log aggregation, monitoring agents

#### 10_health_checks
**Learning Objective:** Implement liveness and readiness probes
- Add liveness probe to detect crashes
- Add readiness probe for traffic management
- Understand probe types (HTTP, TCP, exec)
- **Real-world scenario:** Automatic recovery and zero-downtime deployments

#### 11_env_from_configmap
**Learning Objective:** Inject configuration as environment variables
- Use envFrom to load all ConfigMap keys
- Override specific values
- Understand configuration precedence
- **Real-world scenario:** 12-factor app configuration

### 3. Container Fundamentals (Beginner-Intermediate)

#### 12_basic_dockerfile
**Learning Objective:** Create a simple container image
- Write a Dockerfile for a Python/Node.js app
- Understand FROM, COPY, RUN, CMD
- Build and run the image
- **Real-world scenario:** Containerizing your application

#### 13_dockerfile_optimization
**Learning Objective:** Reduce image size and build time
- Use layer caching effectively
- Minimize layers
- Remove unnecessary files
- **Real-world scenario:** Faster CI/CD pipelines, lower storage costs

#### 14_multistage_build
**Learning Objective:** Separate build and runtime dependencies
- Create multi-stage Dockerfile
- Build stage vs runtime stage
- Drastically reduce final image size
- **Real-world scenario:** Production-ready containers without dev tools

#### 15_dockerignore
**Learning Objective:** Exclude unnecessary files from build context
- Create .dockerignore file
- Understand build context
- Speed up builds
- **Real-world scenario:** Keeping secrets out of images, faster builds

#### 16_container_debugging
**Learning Objective:** Debug running containers
- Use docker exec to inspect containers
- View logs
- Check running processes
- **Real-world scenario:** Troubleshooting production issues

### 4. Kubernetes Advanced Resources (Intermediate)

#### 17_persistent_volume
**Learning Objective:** Understand persistent storage
- Create a PersistentVolume
- Create a PersistentVolumeClaim
- Mount volume to a pod
- **Real-world scenario:** Database storage, file uploads

#### 18_namespace_management
**Learning Objective:** Isolate resources with namespaces
- Create namespaces for different environments
- Deploy resources to specific namespaces
- List resources by namespace
- **Real-world scenario:** Multi-tenant clusters, environment isolation

#### 19_rolling_update
**Learning Objective:** Update applications without downtime
- Perform a rolling update
- Monitor rollout status
- Understand update strategies
- **Real-world scenario:** Deploying new versions safely

#### 20_rollback_deployment
**Learning Objective:** Recover from bad deployments
- Deploy a broken version
- Check rollout history
- Rollback to previous version
- **Real-world scenario:** Incident response and recovery

#### 21_horizontal_pod_autoscaler
**Learning Objective:** Auto-scale based on metrics
- Create HPA based on CPU usage
- Generate load to trigger scaling
- Observe scaling behavior
- **Real-world scenario:** Handling traffic spikes automatically

### 5. CI/CD Fundamentals (Intermediate)

#### 22_basic_build_script
**Learning Objective:** Automate build process
- Write a shell script to build application
- Run tests
- Create artifacts
- **Real-world scenario:** Local development workflow

#### 23_github_actions_basic
**Learning Objective:** Understand CI/CD pipelines (uses local .github/workflows)
- Create a simple GitHub Actions workflow YAML
- Understand jobs, steps, actions
- Run tests on commit
- **Real-world scenario:** Automated testing on every push

#### 24_docker_build_pipeline
**Learning Objective:** Build containers in CI
- Create pipeline to build Docker image
- Tag images properly
- Understand build caching
- **Real-world scenario:** Automated container builds

#### 25_kubectl_dry_run
**Learning Objective:** Test deployments safely
- Use kubectl apply --dry-run
- Validate manifests before applying
- Understand client-side vs server-side validation
- **Real-world scenario:** Preventing production mistakes

### 6. Configuration & Best Practices (Intermediate-Advanced)

#### 26_pod_security
**Learning Objective:** Run containers securely
- Set security context
- Run as non-root user
- Set read-only filesystem
- **Real-world scenario:** Security hardening for production

#### 27_network_policy
**Learning Objective:** Control pod-to-pod communication
- Create a NetworkPolicy
- Allow traffic from specific pods
- Deny all by default
- **Real-world scenario:** Microservice security, zero-trust networking

#### 28_ingress_controller
**Learning Objective:** HTTP/HTTPS routing to services
- Create an Ingress resource
- Configure path-based routing
- Understand ingress controllers
- **Real-world scenario:** External access to multiple services

#### 29_init_containers
**Learning Objective:** Run initialization tasks before main container
- Create init container for setup tasks
- Understand init container lifecycle
- Wait for dependencies
- **Real-world scenario:** Database migrations, waiting for services

#### 30_pod_disruption_budget
**Learning Objective:** Maintain availability during updates
- Create PodDisruptionBudget
- Ensure minimum available replicas
- Understand voluntary vs involuntary disruptions
- **Real-world scenario:** Maintenance windows without downtime

### 7. Troubleshooting & Debugging (Intermediate)

#### 31_debug_crashloop
**Learning Objective:** Fix a pod in CrashLoopBackOff
- Investigate pod that keeps crashing
- Use kubectl logs and describe
- Fix the underlying issue
- **Real-world scenario:** Common production issue

#### 32_debug_imagepull
**Learning Objective:** Resolve image pull errors
- Fix ImagePullBackOff error
- Check image name and tag
- Understand registry authentication
- **Real-world scenario:** Deployment failures

#### 33_debug_pending_pod
**Learning Objective:** Understand why pods won't schedule
- Investigate pod stuck in Pending
- Check resource constraints
- Fix node selectors or taints
- **Real-world scenario:** Cluster capacity planning

#### 34_service_not_reachable
**Learning Objective:** Debug networking issues
- Fix service that's not routing traffic
- Check selectors and endpoints
- Verify port configuration
- **Real-world scenario:** Service connectivity problems

### 8. Docker Compose & Local Development (Beginner-Intermediate)

#### 35_docker_compose_basic
**Learning Objective:** Multi-container apps with Docker Compose
- Create docker-compose.yml
- Define multiple services
- Link containers
- **Real-world scenario:** Local development environment

#### 36_docker_compose_volumes
**Learning Objective:** Persist data with volumes
- Define volumes in docker-compose
- Mount local directories
- Understand volume lifecycle
- **Real-world scenario:** Database persistence, hot-reload development

#### 37_docker_compose_networks
**Learning Objective:** Custom networking in Compose
- Create custom networks
- Isolate services
- Understand service discovery
- **Real-world scenario:** Microservices local testing

### 9. Monitoring & Observability (Intermediate)

#### 38_pod_metrics
**Learning Objective:** View resource usage
- Install metrics-server
- Use kubectl top
- Understand resource consumption
- **Real-world scenario:** Performance monitoring

#### 39_structured_logging
**Learning Objective:** Implement proper logging
- Output structured JSON logs
- Include correlation IDs
- Follow logging best practices
- **Real-world scenario:** Log aggregation and searching

#### 40_application_health_endpoint
**Learning Objective:** Create health check endpoints
- Implement /health endpoint
- Implement /readiness endpoint
- Return proper HTTP status codes
- **Real-world scenario:** Integration with orchestration platforms

### 10. Advanced Scenarios (Advanced)

#### 41_statefulset_basics
**Learning Objective:** Deploy stateful applications
- Create a StatefulSet
- Understand stable network identities
- Ordered deployment and scaling
- **Real-world scenario:** Databases, distributed systems

#### 42_job_and_cronjob
**Learning Objective:** Run batch workloads
- Create a Job for one-time tasks
- Create a CronJob for scheduled tasks
- Understand job completion and parallelism
- **Real-world scenario:** Data processing, backups, cleanup tasks

#### 43_configmap_hot_reload
**Learning Objective:** Update config without restart
- Mount ConfigMap as volume
- Trigger config reload in application
- Understand volume mount behavior
- **Real-world scenario:** Dynamic configuration updates

#### 44_blue_green_deployment
**Learning Objective:** Zero-downtime deployment strategy
- Maintain two identical environments
- Switch traffic using service selectors
- Instant rollback capability
- **Real-world scenario:** Risk-free production deployments

#### 45_canary_deployment
**Learning Objective:** Gradual rollout strategy
- Deploy new version to small subset
- Monitor metrics
- Gradually increase traffic
- **Real-world scenario:** Testing new features with limited blast radius

## Exercise Progression Path

### Path 1: Kubernetes Fundamentals (For Complete Beginners)
01 → 02 → 03 → 04 → 05 → 06 → 07 → 08 → 10

### Path 2: Container & Docker Track
12 → 13 → 14 → 15 → 16 → 35 → 36 → 37

### Path 3: Kubernetes Production Ready
03 → 07 → 10 → 17 → 18 → 19 → 20 → 26 → 30

### Path 4: Troubleshooting Specialist
31 → 32 → 33 → 34 → 38

### Path 5: Full DevOps Journey
01 → 02 → 03 → 12 → 14 → 04 → 05 → 10 → 22 → 23 → 24 → 19 → 20

## Implementation Priority

### Phase 1 (Foundation) - Next 5 Exercises
1. 03_deploy_deployment - Critical Kubernetes concept
2. 04_create_configmap - Essential for configuration management
3. 12_basic_dockerfile - Fundamental container skill
4. 10_health_checks - Production-ready deployments
5. 07_resource_limits - Cluster stability

### Phase 2 (Practical Skills) - Next 5 Exercises
6. 19_rolling_update - Common deployment pattern
7. 20_rollback_deployment - Incident response
8. 14_multistage_build - Optimization skill
9. 31_debug_crashloop - Common troubleshooting
10. 32_debug_imagepull - Another common issue

### Phase 3 (Advanced Concepts) - Next 5 Exercises
11. 17_persistent_volume - Stateful applications
12. 28_ingress_controller - External access
13. 41_statefulset_basics - Advanced workloads
14. 42_job_and_cronjob - Batch processing
15. 21_horizontal_pod_autoscaler - Auto-scaling

## Exercise Difficulty Legend

- **Beginner**: Basic concepts, guided exercises, clear solutions
- **Intermediate**: Multiple steps, some research required, real-world complexity
- **Advanced**: Complex scenarios, multiple concepts, requires deep understanding

## Notes for Exercise Creation

1. **Local-First**: All exercises should work on minikube/kind
2. **No Paid Services**: Avoid exercises requiring cloud providers (unless optional)
3. **Progressive Complexity**: Each exercise builds on previous knowledge
4. **Real Errors**: Use actual errors students will see in production
5. **Quick Feedback**: Exercises should complete in < 5 minutes
6. **Clear Objectives**: Each exercise teaches one primary concept
7. **Practical Focus**: Prefer real scenarios over theoretical exercises

## Community Contributions

We encourage community members to:
- Suggest new exercises via GitHub issues
- Improve existing exercises
- Add exercises for new DevOps tools (Helm, ArgoCD, etc.)
- Translate exercises to other languages
- Create video walkthroughs

## Future Exercise Areas

- Helm charts and templating
- GitOps with ArgoCD/Flux
- Service mesh basics (Istio/Linkerd)
- Monitoring with Prometheus/Grafana
- Log aggregation with ELK/Loki
- Infrastructure as Code (Terraform basics)
- Container security scanning
- Policy enforcement (OPA/Kyverno)
