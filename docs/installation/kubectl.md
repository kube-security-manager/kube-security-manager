# kubectl

You can use static YAML manifests to install the operator in the `starboard-system` namespace and configure it to select
all namespaces, except `kube-system` and `starboard-system`.

```
kubectl apply -f https://raw.githubusercontent.com/aquasecurity/starboard/{{ git.tag }}/deploy/static/starboard.yaml
```

To confirm that the operator is running, check that the `starboard-operator` Deployment in the `starboard-system`
namespace is available and all its containers are ready:

```console
$ kubectl get deployment -n starboard-system
NAME                 READY   UP-TO-DATE   AVAILABLE   AGE
starboard-operator   1/1     1            1           11m
```

If for some reason it's not ready yet, check the logs of the `starboard-operator` Deployment for errors:

```
kubectl logs deployment/starboard-operator -n starboard-system
```

Starboard ensures the default [Settings] stored in ConfigMaps and Secrets created in the `starboard-system` namespace.
You can always change these settings by editing configuration objects. For example, you can use Trivy in [ClientServer]
mode, which is more efficient that the [Standalone] mode.

You can further adjust the [Configuration](../configuration.md) of the operator with environment variables. For
example, to change the target namespace from all namespaces to the `default` namespace edit the `starboard-operator`
Deployment and change the value of the `OPERATOR_TARGET_NAMESPACES` environment variable from the blank string
(`""`) to the `default` value.

Static YAML manifests with fixed values have shortcomings. For example, if you want to change the container image or
modify default configuration settings, you have to edit existing manifests or customize them with tools such as
[Kustomize]. Thus, we also provide [Helm] chart as an alternative installation option.

## Uninstall

!!! danger
    Uninstalling the operator and deleting custom resource definitions will also delete all generated security reports.

You can uninstall the operator with the following command:

```
kubectl delete -f https://raw.githubusercontent.com/aquasecurity/starboard/{{ git.tag }}/deploy/static/starboard.yaml
```

[Settings]: ../settings.md
[Standalone]: ../vulnerability-scanning/trivy.md#standalone
[ClientServer]: ../vulnerability-scanning/trivy.md#clientserver
[Kustomize]: https://kustomize.io
[Helm]: helm.md
