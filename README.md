# DevOps Dojo

An interactive CLI-based learning tool for mastering DevOps skills, inspired by Rustlings and Golings. Perfect for interns and beginners learning Kubernetes, containerization, and CI/CD concepts through hands-on exercises.

## What is DevOps Dojo?

DevOps Dojo is a command-line tool that helps you learn DevOps concepts by fixing exercises. Each exercise is broken, and your task is to fix it! The exercises cover fundamental DevOps topics that you'll encounter in real-world scenarios.

## Features

- Interactive CLI interface with progress tracking
- Hands-on exercises for Kubernetes, Docker, and CI/CD
- Learn by fixing real problems
- No paid cloud services required - everything runs locally
- Instant feedback on your solutions
- Hints system when you're stuck

## Prerequisites

Before starting, ensure you have the following installed:

- **Go** (1.21 or higher) - [Installation Guide](https://golang.org/doc/install)
- **kubectl** - [Installation Guide](https://kubernetes.io/docs/tasks/tools/)
- **Docker** - [Installation Guide](https://docs.docker.com/get-docker/)
- **Minikube** or **kind** - For local Kubernetes cluster
  - [Minikube Installation](https://minikube.sigs.k8s.io/docs/start/)
  - [kind Installation](https://kind.sigs.k8s.io/docs/user/quick-start/)

## Installation

### Option 1: Install from source

```bash
git clone https://github.com/jenish-jain/devops-dojo.git
cd devops-dojo
make build
sudo mv devops-dojo /usr/local/bin/
```

### Option 2: Build locally

```bash
go build -o devops-dojo
./devops-dojo
```

## Getting Started

### 1. Start your local Kubernetes cluster

Using Minikube:
```bash
minikube start
```

Or using kind:
```bash
kind create cluster
```

### 2. Verify your setup

```bash
kubectl cluster-info
kubectl get nodes
```

### 3. List all exercises

```bash
devops-dojo list
```

### 4. Run an exercise

Run the next pending exercise:
```bash
devops-dojo run next
```

Or run a specific exercise:
```bash
devops-dojo run deploy_a_pod
```

### 5. Fix the exercise

Each exercise file starts with `// I AM NOT DONE`. This marks the exercise as incomplete.

1. Open the exercise file mentioned in the error output
2. Read the code and understand what's wrong
3. Fix the issue
4. Remove the `// I AM NOT DONE` comment when you're satisfied with your solution
5. Run the exercise again to verify

### 6. Get help

If you're stuck, use the hint system:
```bash
devops-dojo hint deploy_a_pod
```

### 7. Watch mode

Auto-run exercises as you save changes:
```bash
devops-dojo watch
```

### 8. Verify all exercises

Check that all exercises pass:
```bash
devops-dojo verify
```

## Available Commands

- `devops-dojo list` - List all exercises with their status
- `devops-dojo run <exercise>` - Run a specific exercise
- `devops-dojo run next` - Run the next pending exercise
- `devops-dojo watch` - Watch mode: auto-run on file changes
- `devops-dojo hint <exercise>` - Get a hint for an exercise
- `devops-dojo verify` - Verify all exercises pass
- `devops-dojo --help` - Show help message

## Exercise Structure

Exercises are organized in the `exercises/` directory:

```
exercises/
├── 01_create_pod/
│   ├── main.go          # Exercise code
│   └── pod.yaml         # Kubernetes manifest
├── 02_create_service/
│   ├── main.go
│   └── service.yaml
└── ...
```

## Current Exercises

1. **deploy_a_pod** - Learn to deploy a basic Kubernetes pod
2. **deploy_a_service** - Learn to create a Kubernetes service

More exercises coming soon!

## Learning Path

The exercises are designed to progressively build your DevOps skills:

1. **Kubernetes Basics** - Pods, Services, Deployments
2. **Configuration Management** - ConfigMaps, Secrets
3. **Storage** - PersistentVolumes, PersistentVolumeClaims
4. **Networking** - Ingress, NetworkPolicies
5. **Container Basics** - Dockerfile, multi-stage builds
6. **CI/CD** - Basic pipelines, testing, deployment

## Tips for Success

1. **Read error messages carefully** - They often contain the solution
2. **Use kubectl** to inspect resources: `kubectl get pods`, `kubectl describe pod <name>`
3. **Check logs** when things fail: `kubectl logs <pod-name>`
4. **Don't skip exercises** - They build on each other
5. **Use hints** when stuck, but try to solve it yourself first
6. **Experiment** - Breaking things is part of learning!

## Troubleshooting

### "kubectl not found"
Install kubectl following the [official guide](https://kubernetes.io/docs/tasks/tools/).

### "No cluster available"
Start your local Kubernetes cluster with `minikube start` or `kind create cluster`.

### Exercise won't run
Make sure you've removed the `// I AM NOT DONE` comment after fixing the exercise.

### Permission errors
Ensure your kubeconfig is properly set up: `kubectl config view`

## Contributing

We welcome contributions! See [CONTRIBUTING.md](./CONTRIBUTING.md) for guidelines on:
- Adding new exercises
- Improving existing exercises
- Reporting bugs
- Suggesting features

## Project Structure

```
devops-dojo/
├── cmd/                 # CLI commands
│   ├── list.go
│   ├── run.go
│   ├── watch.go
│   ├── hint.go
│   └── verify.go
├── exercises/           # All exercises
├── internal/           # Internal packages
│   ├── exercises/      # Exercise logic
│   └── ui/            # UI components
├── info.toml          # Exercise metadata
├── main.go            # Entry point
└── README.md          # This file
```

## License

MIT License - feel free to use this for learning and teaching!

## Acknowledgments

Inspired by:
- [Rustlings](https://github.com/rust-lang/rustlings)
- [Golings](https://github.com/mauricioabreu/golings)

## Support

- Open an issue for bugs or feature requests
- Star the repo if you find it helpful!
- Share with others learning DevOps

Happy Learning! 🚀
