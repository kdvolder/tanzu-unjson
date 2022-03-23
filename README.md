tanzu-unjson
============

Various tanzu components / controllers / pods produce logs in a json format that is hard
to process as a human reader.

This simple commandline utility processes and prints out the log entries in a more
'human friendly' format.

## Installation

```
go install github.com/kdvolder/tanzu-unjson@latest
```

Also make sure to have your `${GOPATH}/bin` in your ${PATH}.

## Usage:

Example:

```
kubectl logs ${pod_name} | tanzu-unjson
```
