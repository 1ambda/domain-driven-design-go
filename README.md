# Domain Driven Design in Go

[![Build Status](https://travis-ci.com/1ambda/domain-driven-design-go.svg?branch=master)](https://travis-ci.com/1ambda/domain-driven-design-go) [![Go Report Card](https://goreportcard.com/badge/github.com/1ambda/domain-driven-design-go)](https://goreportcard.com/report/github.com/1ambda/domain-driven-design-go)

![](https://raw.githubusercontent.com/1ambda/domain-driven-design-go/master/screenshots/g-street.png)


## Features

- Reproducible Infra using Terraform on Google Cloud Platform
    * GKE (Kubernetes Engine)
    * Cloud SQL (HA-RR MySQL)
    * GCE Ingress (GLBC)
    * cert-manager on GKE 
    * Cloud Build
    * Spinnaker on GKE 
- Containerized Applications
    * E-commerce Domain (Simplified)
    * Swagger Codegen for Server (Golang) 
    * Swagger Codegen for Client (Vue, TypeScript)
    * Automated CI / CD Pipeline (GCB -> GCR -> Spinnaker -> GKE) 
    
    
## Running All Services 

| name | version | description |
|---|---|---|
| [go](https://github.com/golang/go) | 1.10+ | use [gvm](https://github.com/moovweb/gvm) |
| [nodejs](https://nodejs.org/) | 8.10.0+ | use [nvm](https://github.com/creationix/nvm) |

```bash
$ go get -g github.com/1ambda/domain-driven-design-go
$ cd $GOPATH/src/github.com/1ambda/domain-driven-design-go 

$ make compose.all

# visit http://localhost:8080
```
