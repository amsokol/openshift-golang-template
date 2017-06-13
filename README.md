# Golang template for OpenShift v3 (version 1.4 and higher)
It supports:
- Golang [dep manager](https://github.com/golang/dep) (dep is a prototype dependency management tool) - [/example-golang-dep](https://github.com/amsokol/openshift-golang-template/tree/master/example-golang-dep)
- [glide](https://github.com/Masterminds/glide) manager - [/example-glide](https://github.com/amsokol/openshift-golang-template/tree/master/example-glide)
- [godep](https://github.com/tools/godep) manager - [/example-godep](https://github.com/amsokol/openshift-golang-template/tree/master/example-godep)
- [govendor](https://github.com/kardianos/govendor) manager - [/example-govendor](https://github.com/amsokol/openshift-golang-template/tree/master/example-govendor)
- Traditional `go get` - [/example](https://github.com/amsokol/openshift-golang-template/tree/master/example)

## Sample data to try golang template:
| Data                                 | Value                                                    |
|--------------------------------------|----------------------------------------------------------|
| Go                                   | Go v1.8                                                  |
| OpenShift                            | OpenShift Origin v1.5                                    |
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

14. [Optional] You can easy provide `go build` arguments. Just add `GO_BUILD_ARGS` environment variable to `Build Configuration` section. For example the following value tells `go build` to print the commands are run and packages are compiled:
```
GO_BUILD_ARGS=-x -v

Note: please DO NOT override output file name ("-o" argument)!
```

15. [Optional] You can easy provide command line arguments to your binary. Just add `GOPROJECT_CMD_ARGS` environment variable to `Deployment Configuration` section. For example the following value provides `--silent` and `--force` arguments the binary:
```
GOPROJECT_CMD_ARGS=--silent --force
```

16. [Optional] You can easy provide configuration files to your binary that are copied by the build script to `GOPATH\bin` (where your binary file is located). Just add `GOPROJECT_CMD_CONFIG`, `GOPROJECT_CMD_CONFIG1`, `GOPROJECT_CMD_CONFIG2`, `GOPROJECT_CMD_CONFIG3`, etc. environment variables to `Build Configuration` section. If your configuration files are located in the same project source control repository `GOPROJECT_CMD_CONFIGx` values should be `GOPROJECT_ROOT` related path. But you need to know that store configuration files into project repository is BAD PRACTICES (configuration files usually stores credentials, IPs, etc.). It's much more better to inject configuration file from outside (e.g. download from CI server). In this case set `GOPROJECT_CMD_CONFIGx` by URL value. Build scripts runs `curl` utility to download configuration files. It supports:
`DICT, FILE, FTP, FTPS, Gopher, HTTP, HTTPS, IMAP, IMAPS, LDAP, LDAPS, POP3, POP3S, RTMP, RTSP, SCP, SFTP, SMB, SMBS, SMTP, SMTPS, Telnet, TFTP, SSL certificates, proxies, HTTP/2, cookies, user+password authentication (Basic, Plain, Digest, CRAM-MD5, NTLM, Negotiate and Kerberos) and more.`
Here are some examples:
```
Provide configuration file is stored into the same source code repository (not recommended for production):
GOPROJECT_CMD_CONFIG=config/settings.yaml

Provide configuration files are stored into the same source code repository (not recommended for production):
GOPROJECT_CMD_CONFIG1=config/db/db.yaml
GOPROJECT_CMD_CONFIG2=config/messaging/broker.json

Download configuration file via HTTPS:
GOPROJECT_CMD_CONFIG=https://config.server.com/myapp/settings.yaml

Download configuration files via HTTPS:
GOPROJECT_CMD_CONFIG1=https://config.server.com/myapp/db/db.yaml
GOPROJECT_CMD_CONFIG2=smb://config.server.com/messaging/broker.json -u "domain\username:passwd"
```
Notes:
- If you have only one configuration file use `GOPROJECT_CMD_CONFIG` environment variable. If you have more that one configuration files use `GOPROJECT_CMD_CONFIG1`, `GOPROJECT_CMD_CONFIG2`, `GOPROJECT_CMD_CONFIG3`, etc. environment variables.
- As you see above you can provide not URL only but other `curl` arguments like credentials. It very important to pass arguments after URL:
```
Correct:
GOPROJECT_CMD_CONFIG=smb://config.server.com/messaging/broker.json -u "domain\username:passwd"

Wrong:
GOPROJECT_CMD_CONFIG=-u "domain\username:passwd" smb://config.server.com/messaging/broker.json
```

17. Leave other options with default values and click `Create` and wait while pod is created

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
s2i build https://github.com/amsokol/openshift-golang-template.git amsokol/golang-openshift:1.8.3-3 golang1 -e GOPROJECT_ROOT=github.com/amsokol/openshift-golang-template/example-golang-dep -e GOPROJECT_CMD=cmd/server --context-dir /example-golang-dep
```

## Helper #2 - how to export/edit template from OpenShift:
1. Login using oc as OpenShift administrator

2. Run:
```
# oc get templates -n openshift
# oc edit template -n openshift nodejs-mongo-persistent
```
