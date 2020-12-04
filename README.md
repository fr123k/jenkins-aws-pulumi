[![Build Status](https://travis-ci.com/fr123k/jenkins-aws-pulumi.svg?branch=master)](https://travis-ci.com/fr123k/jenkins-aws-pulumi)

# jenkins-aws-pulumi

This repository is a tutorial how to setup pulumi golang with aws.

## Introduction

The goal of this repository is to have [pulumi](https://www.pulumi.com/) based
example that provide the following functionality.

* bootstrap an ec2 instance
* configure cloud-init to start the [jocker](https://hub.docker.com/r/fr123k/jocker) docker container as systemd service

## Prerequisites
```bash
brew install pulumi
```

## Configure pulumi

The config/secret passphrase is `test`.
```bash
export PULUMI_CONFIG_PASSPHRASE=test
pulumi plugin install resource aws 0.18.3
pulumi plugin ls
pulumi login --local
pulumi stack init jenkins-aws-s3
pulumi config set aws:region eu-west-1
export AWS_ACCESS_KEY_ID=<YOUR_ACCESS_KEY_ID>
export AWS_SECRET_ACCESS_KEY=<YOUR_SECRET_ACCESS_KEY>
```

## Build pulumi go

```bash
go mod init main
go mod vendor
```
## Run pulumi

```bash
pulumi up
```
## Multiple aws profiles

https://www.pulumi.com/docs/intro/cloud-providers/aws/setup/

## Todo

* bucket can'r be created because of the following error 
```
creating urn:pulumi:jenkins-aws-s3::jenkins-aws-s3::aws:s3/bucket:Bucket::s3-pulumi-state: AccessDenied: Access Denied
    	status code: 403, request id: 04F5B33564D44CD7, host id: HwKam+lDpRhLDyk/sxy0tnFX6wvMwoBdSIui30/RYCdVcbF1sKpw7FKZO7V4eWSITQMn0OKBh2c=
```
* add pulumi using AWS s3 bucket as state backed for full manage jenkins

## History

26.02.2020: bootstrap ec2 instance with jocker container by cloud-init
