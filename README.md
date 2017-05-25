# Golang template for OpenShift v3 (version 1.4 and higher)
It supports:
- Golang [dep manager](https://github.com/golang/dep) (dep is a prototype dependency management tool) - [/example-golang-dep](https://github.com/amsokol/openshift-golang-template/tree/master/example-golang-dep)
- Traditional `go get` - [/example](https://github.com/amsokol/openshift-golang-template/tree/master/example)

## Sample data to try golang template:
| Data                                 | Value                                                    |
|--------------------------------------|----------------------------------------------------------|
| Go                                   | Go v1.8                                                  |
| OpenShift                            | OpenShift Origin v1.5.0                                  |
| Git repository                       | https://github.com/amsokol/openshift-golang-template.git |
| Context directory                    | /example-golang-dep                                      |
| Folder with main.go to build and run | /example-golang-dep/cmd/server                           |

## How to:
1. Login with Developer account to OpenShift:
```
# oc login --insecure-skip-tls-verify -u <username> https://openshift.example.com:8443
```

2. Create new project:
```
# oc new-project project1
```

3. Select `project1` project:
```
# oc project project1
```

4. Create ImageStream:
```
# oc create -f ./openshift-golang-template.yaml
```

5. Login to OpenShift console using browser (eg https://openshift.example.com:8443) with Developer account

6. Open `project1` project

7. Click `Add to Project` and select `Go`

8. Select `1.8.3 - latest` from drop down list and click `Select`

9. Set `Name` to `golang1`

10. Set `Git Repository URL` to `https://github.com/amsokol/openshift-golang-template.git` or click `Try It`

11. Click `advanced options` to add additional configuration params

12. Set `Context Dir` to `/example-golang-dep` (sample project with [dep manager](https://github.com/golang/dep) support)

13. Add the following environment variables to `Build Configuration` section:
- GOPROJECT_ROOT=github.com/amsokol/openshift-golang-template/example-golang-dep
- GOPROJECT_CMD=cmd/server

```
GOPROJECT_ROOT tells builder the root package of go project
GOPROJECT_CMD tells builder the where "main()" function of "main" package to build and run is located (relative to GOPROJECT_ROOT).
Note: ignore GOPROJECT_CMD if "main()" function of "main" package is located in GOPROJECT_ROOT folder. 

In example above "main()" function of "main" package is located in `github.com/amsokol/openshift-golang-template/example-golang-dep/cmd/server`.
GOPROJECT_ROOT is set to `github.com/amsokol/openshift-golang-template/example-golang-dep`.
So GOPROJECT_CMD is set to `cmd/server`
```

14. Leave other options with default values and click `Create` and wait while pod is created

## [Optional] How to add health (liveness and readiness) checks
1. Login to OpenShift console using browser (eg https://openshift.example.com:8443) with Developer account

2. Open `project1` project

3. Click `Applications -> Deployments`

4. Click `golang1` in the list

5. Select `Configuration` tab

6. Click `Add Health Checks`

7. Click `Add Readiness Probe`

8. Set `Path` to '/healthz/ready' and leave other options with default values

9. Click `Add Liveness Probe`

10. Set `Path` to '/healthz/live' and leave other options with default values

11. Click `Save` and wait while pod is created

## Helper #1 - you can try golang template using S2I:
```
s2i build https://github.com/amsokol/openshift-golang-template.git amsokol/golang-openshift:1.8.3-1 golang1 -e GOPROJECT_ROOT=github.com/amsokol/openshift-golang-template/example-golang-dep -e GOPROJECT_CMD=cmd/server --context-dir /example-golang-dep
```

## Helper #2 - how to export/edit template from OpenShift:
1. Login using oc as OpenShift administrator

2. Run:
```
# oc get templates -n openshift
# oc edit template -n openshift nodejs-mongo-persistent
```
