# goal
Make sure aws x-ray works across multiple environments like k8s and aws lambdas.

# spoiler
![xray-complete-trace](xray-complete-trace.png?raw=true "xray-complete-trace")

# requirements
- aws account
- minikube
- virtualbox
- docker
- go
- serverless

# usage

deploy lambda

```bash
# export aws credentials
export AWS_SECRET_ACCESS_KEY=...
export AWS_ACCESS_KEY_ID=...
# build and deploy
$ cd lambda
$ make build
$ sls deploy
```

run local k8s cluster and deploy app and x-ray daemon

```bash
$ minikube start --driver=virtualbox
# add aws credentials
$ kubectl create secret generic aws \
--from-literal=AWS_ACCESS_KEY_ID=... \
--from-literal=AWS_SECRET_ACCESS_KEY=... \
--from-literal=AWS_LAMBDA_URL=...
# build and push app image
$ k8s/app/build.sh
# deploy app
$ kubectl apply -f k8s/app/manifest.yaml
# deploy xray daemon
$ kubectl apply -f k8s/xrayd/manifest.yaml
```

make a remote call that will be traced

```bash
# grab the app url
$ minikube service xray-k8s-app --url
# make a call
$ curl <url>/ip -sik
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
- [x] deploy app-02 to a lambda
- [x] instrument app-02
- [x] expect complete tracing data to be present in x-ray gui

# refs
- aws lambda and x-ray https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html
- x-ray on k8s https://aws.amazon.com/blogs/compute/application-tracing-on-kubernetes-with-aws-x-ray/
- x-ray dev guide https://docs.aws.amazon.com/xray/latest/devguide/aws-xray.html
- sls permissions https://www.serverless.com/blog/abcs-of-iam-permissions
- sls and x-ray https://www.serverless.com/framework/docs/providers/aws/guide/functions/#aws-x-ray-tracing
