# Contributing to DevOps Dojo

Thank you for your interest in contributing to DevOps Dojo! This document provides guidelines for adding new exercises and improving the project.

## Table of Contents

- [Getting Started](#getting-started)
- [Adding New Exercises](#adding-new-exercises)
- [Exercise Guidelines](#exercise-guidelines)
- [Testing Your Exercise](#testing-your-exercise)
- [Submitting Your Contribution](#submitting-your-contribution)

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/devops-dojo.git
   cd devops-dojo
   ```
3. Create a new branch:
   ```bash
   git checkout -b feature/new-exercise-name
   ```

## Adding New Exercises

### Step 1: Create Exercise Directory

Create a new directory in `exercises/` with a descriptive name:

```bash
mkdir -p exercises/03_new_exercise
cd exercises/03_new_exercise
```

### Step 2: Create Exercise Files

Each exercise needs:
- `main.go` - The main exercise code
- Any supporting files (YAML files, Dockerfiles, etc.)

Example `main.go`:

```go
// I AM NOT DONE
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    fmt.Println("Exercise: Deploy a ConfigMap")

    // Check if kubectl is available
    if _, err := exec.LookPath("kubectl"); err != nil {
        fmt.Println("Error: kubectl is not installed or not in PATH")
        fmt.Println("Please install kubectl to run this exercise")
        fmt.Println("Visit: https://kubernetes.io/docs/tasks/tools/")
        os.Exit(1)
    }

    // Your exercise logic here
    yamlPath := "exercises/03_new_exercise/configmap.yaml"

    // Add validation
    if _, err := os.Stat(yamlPath); os.IsNotExist(err) {
        fmt.Printf("Error: %s does not exist\n", yamlPath)
        os.Exit(1)
    }

    // Execute kubectl command
    cmd := exec.Command("kubectl", "apply", "-f", yamlPath)
    stdout, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Println("Error executing kubectl:")
        fmt.Println(string(stdout))
        os.Exit(1)
    }

    fmt.Println("Success!")
    fmt.Println(string(stdout))
}
```

### Step 3: Add Exercise to info.toml

Edit `info.toml` and add your exercise:

```toml
[[exercises]]
name = "new_exercise_name"
path = "exercises/03_new_exercise/main.go"
mode = "compile"
hint = """
Your helpful hint here.
Explain what the user should look for or where to start.
"""
```

### Step 4: Create Supporting Files

If your exercise uses Kubernetes manifests, create them with intentional issues for users to fix:

```yaml
# configmap.yaml - This might have an error students need to fix
apiVersion: v1
kind: ConfigMap
metadata:
  name: my-config
data:
  key: value
```

## Exercise Guidelines

### Good Exercise Characteristics

1. **One Concept Per Exercise**
   - Focus on a single learning objective
   - Don't combine too many concepts

2. **Progressive Difficulty**
   - Build on previous exercises
   - Introduce new concepts gradually

3. **Real-World Relevance**
   - Use scenarios from actual DevOps work
   - Teach practical, applicable skills

4. **Clear Learning Objective**
   - State what the user will learn
   - Make the goal obvious

5. **Helpful Error Messages**
   - Provide clear error messages
   - Guide users toward the solution without giving it away

### Exercise Types

#### Type 1: Kubernetes Deployment Exercises
Focus on deploying and managing Kubernetes resources.

**Topics:**
- Pods, Services, Deployments
- ConfigMaps, Secrets
- PersistentVolumes, StatefulSets
- Ingress, NetworkPolicies

#### Type 2: Container Exercises
Focus on building and optimizing containers.

**Topics:**
- Writing Dockerfiles
- Multi-stage builds
- Image optimization
- Security best practices

#### Type 3: CI/CD Pipeline Exercises
Focus on automation and pipelines.

**Topics:**
- Basic pipeline structure
- Testing automation
- Build and deployment stages
- Pipeline debugging

#### Type 4: Configuration & Infrastructure
Focus on configuration management and IaC.

**Topics:**
- Environment variables
- Configuration files
- Basic scripting
- Health checks

### Exercise Template

```go
// I AM NOT DONE
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    fmt.Println("Exercise: [Clear description]")

    // 1. Validate prerequisites
    if _, err := exec.LookPath("required-tool"); err != nil {
        fmt.Println("Error: required-tool is not installed")
        fmt.Println("Install from: [link]")
        os.Exit(1)
    }

    // 2. Check required files
    requiredFile := "path/to/file"
    if _, err := os.Stat(requiredFile); os.IsNotExist(err) {
        fmt.Printf("Error: %s not found\n", requiredFile)
        os.Exit(1)
    }

    // 3. Execute the exercise logic
    cmd := exec.Command("tool", "args...")
    output, err := cmd.CombinedOutput()

    // 4. Handle errors gracefully
    if err != nil {
        fmt.Println("Error:")
        fmt.Println(string(output))
        os.Exit(1)
    }

    // 5. Show success message
    fmt.Println("Success!")
    fmt.Println(string(output))
}
```

### Best Practices

1. **Always use `// I AM NOT DONE`** at the top of exercise files
2. **Validate dependencies** before running commands
3. **Use `os.Exit(1)` instead of `panic()`** for errors
4. **Provide helpful error messages** with links to documentation
5. **Use `CombinedOutput()`** to capture both stdout and stderr
6. **Test on a clean environment** before submitting

### Writing Good Hints

Hints should:
- Guide without giving away the solution
- Point to relevant documentation
- Suggest where to look or what to check
- Be encouraging

Example:
```toml
hint = """
Check the pod.yaml file - the image name might be incorrect.
Hint: nginx images are named 'nginx', not 'nginx-server'.
Try running: kubectl describe pod <name> to see the exact error.
"""
```

## Testing Your Exercise

### Manual Testing

1. **Build the project:**
   ```bash
   go build -o devops-dojo
   ```

2. **List exercises:**
   ```bash
   ./devops-dojo list
   ```

3. **Run your exercise:**
   ```bash
   ./devops-dojo run your_exercise_name
   ```

4. **Test with the `// I AM NOT DONE` comment:**
   - Should fail and show helpful error

5. **Test without the comment:**
   - Should pass when fixed correctly

6. **Test the hint:**
   ```bash
   ./devops-dojo hint your_exercise_name
   ```

### Automated Testing

Run the verification command:
```bash
./devops-dojo verify
```

This ensures all exercises pass when correctly solved.

### Checklist Before Submitting

- [ ] Exercise runs without errors when fixed
- [ ] Exercise fails with `// I AM NOT DONE` present
- [ ] Hint is helpful but doesn't give away the solution
- [ ] All required files are included
- [ ] Error messages are clear and actionable
- [ ] Code follows Go best practices
- [ ] Exercise is added to `info.toml`
- [ ] Exercise teaches a clear concept
- [ ] Documentation is updated if needed

## Submitting Your Contribution

1. **Commit your changes:**
   ```bash
   git add .
   git commit -m "Add exercise: brief description"
   ```

2. **Push to your fork:**
   ```bash
   git push origin feature/new-exercise-name
   ```

3. **Create a Pull Request:**
   - Go to the original repository
   - Click "New Pull Request"
   - Select your branch
   - Fill in the PR template with:
     - Exercise name and number
     - What it teaches
     - Prerequisites
     - Testing done

4. **Respond to feedback:**
   - Address review comments
   - Update your PR as needed

## Exercise Ideas

Looking for ideas? Here are some suggested exercises:

### Beginner
- Deploy a deployment with replicas
- Create and use a ConfigMap
- Create and use a Secret
- Create a simple Dockerfile
- Build and run a Docker container

### Intermediate
- Create a service with proper selectors
- Deploy an application with health checks
- Create a multi-container pod
- Use environment variables from ConfigMaps
- Implement resource limits and requests

### Advanced
- Create a StatefulSet
- Implement an Ingress resource
- Set up a NetworkPolicy
- Multi-stage Docker build
- Rolling update and rollback

## Code Style

- Follow standard Go formatting (`go fmt`)
- Use meaningful variable names
- Add comments for complex logic
- Keep functions focused and small
- Handle errors explicitly

## Questions?

- Open an issue for questions
- Tag it as `question`
- We're here to help!

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

Thank you for making DevOps Dojo better! 🎉
