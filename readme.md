# goal
Make sure aws x-ray works across multiple environments like k8s and aws lambdas.

# requirements
- minikube
- virtualbox
- docker
- go

# usage
```bash
# run local cluster
$ minikube start --driver=virtualbox
# create secret
$ kubectl create secret generic aws \
--from-literal=AWS_ACCESS_KEY_ID=... \
--from-literal=AWS_SECRET_ACCESS_KEY=...
# build and push app image
$ k8s/app/build.sh
# deploy app
$ kubectl apply -f k8s/app/manifest.yaml
# deploy xray daemon
$ kubectl apply -f k8s/xrayd/manifest.yaml
```

# x-ray

- sampling

The default rule traces the first request each second, and five percent of any additional requests across all services sending traces to X-Ray. Update rules https://docs.aws.amazon.com/xray/latest/devguide/xray-console-sampling.html

# todos
- [x] install minikube
- [x] install docker
- [x] install virtualbox
- [x] deploy app-01 to k8s
- [x] instrument app-01
- [x] expect traces in x-ray gui
- [ ] instrument and deploy app-02 to a lambda
- [ ] expect complete tracing data to be present in x-ray gui

# refs
- aws lambda and x-ray https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html
- x-ray on k8s https://aws.amazon.com/blogs/compute/application-tracing-on-kubernetes-with-aws-x-ray/
- x-ray sdk go https://docs.aws.amazon.com/xray/latest/devguide/xray-sdk-go.html
