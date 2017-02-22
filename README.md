# openshift-golang-template

how to export template from OpenShift:
$ oc get templates -n openshift
$ oc edit template -n openshift nodejs-mongo-persistent

Steps to start (example):
$ oc login --insecure-skip-tls-verify -u <username> https://console.openshift.example.com:8443
$ oc new-project project1
$ oc create -f .\openshift-golang-template.yaml

Note: select to show advanced options and set the following:
1) set context dir '/example'
2) Set 'ROOT' environment variable (for build and run) as root package. Example:
   Git URL: https://github.com/amsokol/openshift-golang-template.git
   Context directory: /example
   So ROOT=github.com/amsokol/openshift-golang-template/example
