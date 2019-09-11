# Kubernetes and Cloud Native Meetup @ Etermax HQ, Buenos Aires

### How to run:

#### Prerequisites
You need to have access to a Kubernetes cluster. The rest of the guide assumes you are using either **Minikube** or **Docker for Mac**

#### Add an entry to your /etc/hosts

If using **Docker for Mac**:
```
127.0.0.1 api.example.com traefik.example.com
```

If using **Minikube**:
```
<your minikube ip> api.example.com traefik.example.com
```

*This is optional, but you'll have to change some things if you ignore this step*

#### Deploy the kubernetes manifests
With the proper Kubernetes configuration in place, just run

```
kubectl apply -R -f k8s
```

#### Test that everything is working
Wait a few seconds until everything is running, fire up your browser and go to https://traefik.example.com

You should be able to see the `Traefik Web UI`

#### Editing the bundle
Commiting to this repo will automatically update the bundle. However, you will need to set some variables in travis:

```
GITHUB_USER=<your GitHub username>
GITHUB_REPO=<the name of the repo (k8s-meetup-20191112)>
GITHUB_TOKEN=<a valid GitHub access token>
```

Then, you'll need to modify *k8s/traefik/config-map.yml* to point to the bundle generated for your project