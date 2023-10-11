# Contributing

These guidelines will help you get started with Kubernetes Security Manager.

## Table of Contents

- [Set up your Development Environment](#set-up-your-development-environment)
- [Build Container Image](#build-container-image)
- [Run Tests](#run-tests)
  - [Run Unit Tests](#run-unit-tests)
  - [Run Integration Tests](#run-integration-tests)
  - [Cove Coverage](#code-coverage)
- [Custom Resource Definitions](#custom-resource-definitions)
  - [Generate Code](#generate-code)
- [Test Starboard Operator](#test-starboard-operator)
  - [In Cluster](#in-cluster)
  - [Out of Cluster](#out-of-cluster)

## Set up your Development Environment

1. Install Go

   The project requires [Go 1.19][go-download] or later. We also assume that
   you are familiar with Go's [GOPATH workspace][go-code] convention, and have
   the appropriate environment variables set.
2. Get the source code:
   ```
   git clone --recurse-submodules https://github.com/kube-security-manager/kube-security-manager.git
   cd kube-security-manager
   ```
3. Access to a Kubernetes cluster. We assume that you're using a [KIND][kind]
   cluster. To create a single-node KIND cluster, run:
   ```
   kind create cluster \
     --image="kindest/node:v1.21.14@sha256:220cfafdf6e3915fbce50e13d1655425558cb98872c53f802605aa2fb2d569cf"
   ```

## Build Container Image

| Binary                  | Image                                             | Description        |
|-------------------------|---------------------------------------------------|--------------------|
| `kube-security-manager` | `docker.io/danielpacak/kube-security-manager:dev` | Starboard Operator |

To build the `docker.io/danielpacak/kube-security-manager:dev` container image, run:

```
make docker-image
```


To load the Docker image into your KIND cluster, run:

```
make kind-load-image
```

## Run Tests

We generally require tests to be added for all, but the most trivial of changes. However, unit tests alone don't
provide guarantees about the behaviour of Starboard. To verify that each Go module correctly interacts with its
collaborators, more coarse grained integration tests might be required.

### Run Unit Tests

To run all unit tests with code coverage enabled, run:

```
make unit-tests
```

To open the test coverage report in your web browser, run:

```
go tool cover -html=coverage.txt
```

### Run Integration Tests

The integration tests assumes that you have a working kubernetes cluster (e.g KIND cluster) and `KUBECONFIG` environment
variable is pointing to that cluster configuration file. For example:

```
export KUBECONFIG=~/.kube/config
```

There are separate integration tests for Starboard CLI and for Starboard Operator. The tests may leave the cluster in a
dirty state, so running one test after the other may cause spurious failures.

To run the integration tests for Starboard CLI with code coverage enabled, run:

```
make itests-starboard
```

To open the test coverage report in your web browser, run:

```
go tool cover -html=itest/starboard/coverage.txt
```

To run the integration tests for Starboard Operator and view the coverage report, first do the
[prerequisite steps](#prerequisites), and then run:

```
OPERATOR_NAMESPACE=starboard-system \
  OPERATOR_TARGET_NAMESPACES=default \
  OPERATOR_LOG_DEV_MODE=true \
  make itests-starboard-operator
go tool cover -html=itest/starboard-operator/coverage.txt
```

### Code Coverage

In the CI workflow, after running all tests, we do upload code coverage reports to [Codecov][codecov]. Codecov will
merge the reports automatically while maintaining the original upload context as explained
[here][codecov-merging-reports].

## Custom Resource Definitions

### Generate Code

Code generators are used a lot in the implementation of native Kubernetes resources, and we're using the very same
generators here for custom security resources. This project follows the patterns of
[k8s.io/sample-controller][k8s-sample-controller], which is a blueprint for many controllers built in Kubernetes itself.

The code generation starts with:

```
go mod vendor
export GOPATH="$(go env GOPATH)"
./hack/update-codegen.sh
```

In addition, there is a second script called `./hack/verify-codegen.sh`. This script calls the
`./hack/update-codegen.sh` script and checks whether anything changed, and then it terminates with a nonzero return
code if any of the generated files is not up-to-date. We're running it as a step in the CI workflow.

## Test Starboard Operator

You can deploy Kubernetes Security Manager in the `starboard-system` namespace
and configure it to watch the `default` namespace. In OLM terms such install
mode is called *SingleNamespace*. The *SingleNamespace* mode is good to get
started with a basic development workflow. For other install modes see
[Operator Multitenancy with OperatorGroups][olm-operator-groups].

### In cluster

1. Build the Docker container image and load it from your host into KIND
   cluster nodes:
   ```
   make kind-load-image
   ```
2. Create the `starboard-operator` Deployment in the `starboard-system`
   namespace to run the operator's container:
   ```
   kubectl create -k deploy/static
   ```

You can uninstall Kubernetes Security Manager with:

```
kubectl delete -k deploy/static
```

### Out of cluster

1. Deploy the operator in cluster:
   ```
   kubectl apply -f deploy/static/starboard.yaml
   ```
2. Scale the operator down to zero replicas:
   ```
   kubectl scale deployment starboard-operator \
     -n starboard-system \
     --replicas 0
   ```
3. Delete pending scan jobs with:
   ```
   kubectl delete jobs -n starboard-system --all
   ```
4. Run the main method of the operator program:
   ```
   OPERATOR_NAMESPACE=starboard-system \
     OPERATOR_TARGET_NAMESPACES=default \
     OPERATOR_LOG_DEV_MODE=true \
     OPERATOR_CIS_KUBERNETES_BENCHMARK_ENABLED=true \
     OPERATOR_VULNERABILITY_SCANNER_ENABLED=true \
     OPERATOR_VULNERABILITY_SCANNER_SCAN_ONLY_CURRENT_REVISIONS=false \
     OPERATOR_CONFIG_AUDIT_SCANNER_ENABLED=false \
     OPERATOR_CONFIG_AUDIT_SCANNER_SCAN_ONLY_CURRENT_REVISIONS=false \
     OPERATOR_CONFIG_AUDIT_SCANNER_BUILTIN=true \
     OPERATOR_VULNERABILITY_SCANNER_REPORT_TTL="" \
     OPERATOR_BATCH_DELETE_LIMIT=3 \
     OPERATOR_BATCH_DELETE_DELAY="30s" \
     go run cmd/starboard-operator/main.go
   ```

You can uninstall the operator with:

```
kubectl delete -f deploy/static/starboard.yaml
```

## Resources

* https://github.com/eunomia-bpf/bpf-developer-tutorial/

[go-download]: https://golang.org/dl/
[go-code]: https://golang.org/doc/code.html
[kind]: https://github.com/kubernetes-sigs/kind
[codecov]: https://codecov.io/
[codecov-merging-reports]: https://docs.codecov.io/docs/merging-reports/
[k8s-sample-controller]: https://github.com/kubernetes/sample-controller
