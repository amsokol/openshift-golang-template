# openshift-golang-template

## Steps to start (example):
```
# oc login --insecure-skip-tls-verify -u <username> https://console.openshift.example.com:8443
# oc new-project project1
# oc create -f .\openshift-golang-template.yaml
```
Context directory is "/example/cmd/example-http-server"

### Note: select advanced options and set the following:
1. Set context dir '/example'
2. Set 'ROOT' environment variable (for build and run) as root package. Example:
    * if git URL is "https://github.com/amsokol/openshift-golang-template.git"
    * and context directory is "/example/cmd/example-http-server"
    * than set ROOT environment variable to "github.com/amsokol/openshift-golang-template/example"

## Helper: how to export template from OpenShift to be example:
```
# oc get templates -n openshift
# oc edit template -n openshift nodejs-mongo-persistent
```
