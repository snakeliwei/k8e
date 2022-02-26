<img src="./docs/k8e-logo.png" alt="logo" style="zoom:30%;" /><br/>
K8e 🚀 (said 'kuber easy') - Simple Kubernetes Distribution
===============================================
[Kubernetes Easy (k8e)](https://getk8e.com) is a lightweight, Extensible, Enterprise Kubernetes distribution that allows users to uniformly manage, secure, and get out-of-the-box kubernetes cluster for enterprise environments.

The k8e 🚀 (said 'kuber easy') project builds on upstream project [K3s](https://github.com/rancher/k3s) as codebase, remove Edge/IoT features and extend enterprise features with best practices.

[![Go Report Card](https://goreportcard.com/badge/github.com/xiaods/k8e)](https://goreportcard.com/report/github.com/xiaods/k8e) [![Hex.pm](https://img.shields.io/hexpm/l/apa)](https://github.com/xiaods/k8e/blob/master/LICENSE)

Great for:
* CI
* Development
* Enterprise Deployment

Quick-Start - Building && Installing
--------------
1. Building `k8e`

The clone will be much faster on this repo if you do
```bash
git clone --depth 1 https://github.com/xiaods/k8e.git
```

This repo includes all of Kubernetes history so `--depth 1` will avoid most of that.

The k8e build process requires some autogenerated code and remote artifacts that are not checked in to version control.
To prepare these resources for your build environment, run:.
```bash
mkdir -p build/data && make download && make generate
```

To build the full release binary, you may now run `make`, which will create `./dist/artifacts/k8e`.

To build the binaries using without running linting (ie; if you have uncommitted changes):
```bash
SKIP_VALIDATE=true make
```
2. Run server.

```bash
sudo ./k8e check-config
sudo ./k8e server &
# Kubeconfig is written to /etc/k8e/k8e.yaml
export KUBECONFIG=/etc/k8e/k8e.yaml
sudo ./k8e kubectl get nodes

# On a different node run the below. NODE_TOKEN comes from
# /var/lib/k8e/server/node-token on your server
sudo ./k8e agent --server https://myserver:6443 --token ${NODE_TOKEN}
```

Acknowledgments
--------------
- Thanks [k3s](https://github.com/rancher/k3s) for the great open source project.