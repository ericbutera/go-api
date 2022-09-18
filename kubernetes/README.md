# K8s Deployment

- uses <https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/>
- build
  - `kustomize build base`
- build production
  - `kustomize build overlays/prod`
- deploy production
  - `kustomize build overlays/prod | kubectl apply -f -`
