# openshift-golang-template

## Steps to start (example):
```
# oc login --insecure-skip-tls-verify -u <username> https://console.openshift.example.com:8443
# oc new-project project1
# oc create -f .\openshift-golang-template.yaml
```
Context directory is "/example/cmd/example-http-server"

### Note: select advanced options and set the following:
1. Set context dir "/example"
2. Set GOPROJECT_ROOT and GOPROJECT_CMD environment variables (for build and run phases):
    * if, for example, git URL is "https://github.com/amsokol/openshift-golang-template.git"
    * and context directory is "/example"
    * and "main" package is in "/example/cmd/example-http-server"
    * than set GOPROJECT_ROOT environment variable to "github.com/amsokol/openshift-golang-template/example"
    * at last set GOPROJECT_CMD environment variable to "cmd/example-http-server"

## Helper: how to export template from OpenShift to be example:
```
# oc get templates -n openshift
# oc edit template -n openshift nodejs-mongo-persistent
```
