# K8s, ARGO CD and ARGO Rollout and Docker

# Installation:

## MiniKube

Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster on your personal computer so you can try out Kubernetes, or for daily development work.

Installation steps are pretty straightford and is available at [https://minikube.sigs.k8s.io/docs/start/](ghttps://minikube.sigs.k8s.io/docs/start/) for Linux, MacOS and Windows.

However since Argo Rollouts is not natively available in Windows (something I found out too late), working with WSL2 is possible instead of dealing with a docker image. Using the instructions for linux will lead to certificate and plethora of other errors. Instead use the instructions [here](https://gist.github.com/wholroyd/748e09ca0b78897750791172b2abb051) and it lead you to a smooth install.

## Kubectl

If you followed the gist to install MiniKube in WSL2, you can skip this section and move to ArgoCD.

Installation of Kubectl is pretty straightford and the instructions can be found [here.](https://kubernetes.io/docs/tasks/tools/)

## Argo CD

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

This will create a new namespace, `argocd`, where Argo CD services and application resources will live.

Download the latest Argo CD version from [https://github.com/argoproj/argo-cd/releases/latest](https://github.com/argoproj/argo-cd/releases/latest). More detailed installation instructions can be found via the [CLI installation documentation](cli_installation.md).

Also available in Mac, Linux and WSL Homebrew:

```bash
brew install argocd
```

## Argo Rollouts

We can install Argo Rollouts by running:

```shell
kubectl create namespace argo-rollouts
kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml

```

or by following the [official documentation that you can find here](https://argoproj.github.io/argo-rollouts/installation/#controller-installation).

You also need to install the [Argo Rollouts `kubectl` plugin](https://argoproj.github.io/argo-rollouts/installation/#kubectl-plugin-installation)

Once you have the plugin, you can start a local version of the Argo Rollouts Dashboard by running in a new terminal the following command:

```shell
kubectl argo rollouts dashboard
```

Then you can access the dashboard by pointing your browser to [http://localhost:3100/rollouts](http://localhost:3100/rollouts)

# Procedure

Clone the repository to any directory

```
git clone https://github.com/athifirshad/k8s-demo && cd k8s-demo
```

We start by deploying the simple api web app that has been dockerized with

```
 kubectl apply -f k8s/api.yml
```

Check the kubernetes dashboard by executing the command below which will return a link:

```
 minikube dashboard
```

Your dashboard should look similiar to this -

![Dashboard](./docs/deployment.png)

### Access The Argo CD API Server

By default, the Argo CD API server is not exposed with an external IP. To access the API server,

```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```

The API server can then be accessed using https://localhost:8080

### Login Using The CLI

The initial password for the `admin` account is auto-generated and stored as
clear text in the field `password` in a secret named `argocd-initial-admin-secret`
in your Argo CD installation namespace. You can simply retrieve this password
using the `argocd` CLI:

```bash
argocd admin initial-password -n argocd
```

We now need to create an app for the api deployment and service:

```bash
 kubectl apply -f k8s/argo.yml
