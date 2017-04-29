# openshift-golang-template

## Steps to start (example):
```
# oc login --insecure-skip-tls-verify -u <username> https://openshift.example.com:8443
# oc new-project project1
# oc create -f ./openshift-golang-template.yaml
```
Context directory is "/example-golang-dep/cmd/server"

### Note: select advanced options and set the following:
1. Set context dir "/example-golang-dep"
2. Set GOPROJECT_ROOT and GOPROJECT_CMD environment variables by corresponding values in "Build Configuration" section, "Environment Variables (Build and Runtime)" list:
    * if, for example, git URL is "https://github.com/amsokol/openshift-golang-template.git"
    * and context directory is "/example-golang-dep"
    * and "main" package is in "/example-golang-dep/cmd/server"
    * than set GOPROJECT_ROOT environment variable to "github.com/amsokol/openshift-golang-template/example-golang-dep"
    * at last set GOPROJECT_CMD environment variable to "cmd/server"

## Helper: how to export template from OpenShift to be example:
```
# oc get templates -n openshift
# oc edit template -n openshift nodejs-mongo-persistent
```
