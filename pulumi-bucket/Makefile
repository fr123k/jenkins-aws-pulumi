export STACK_NAME ?= pulumi-bucket${STACK_SUFFIX}
export AWS_REGION ?= eu-west-1

go-init:
	go mod init main

build:
	go build -o $(shell basename $(shell pwd))

prepare: build
	pulumi plugin install resource aws 3.10.1
	pulumi plugin ls
	# pulumi login --local
	pulumi login --cloud-url s3://s3-pulumi-state-d12f2f1
	# pulumi login --cloud-url s3://$(shell cat ../output/pulumi-bucket.json | jq -r ."bucketName")
	# pulumi stack rm -f ${STACK_NAME}
	#pulumi stack select ${STACK_NAME}
	pulumi stack select -c ${STACK_NAME}
	pulumi config set aws:region eu-west-1

pulumi-init: build
	pulumi plugin install resource aws 3.10.1
	pulumi plugin ls
	pulumi login --local
	# pulumi stack rm -f ${STACK_NAME}
	# pulumi stack select ${STACK_NAME}
	pulumi stack init ${STACK_NAME}
	pulumi config set aws:region eu-west-1

apply:
	pulumi up --yes
	#verbose logging
	#pulumi up --yes --verbose 9 --logtostderr

cleanup:
	pulumi destroy --yes -s ${STACK_NAME}
	pulumi stack rm -f --yes ${STACK_NAME}

local-cleanup:
	pulumi destroy --yes -s ${STACK_NAME} || true
	pulumi stack rm -f --yes ${STACK_NAME} || true

deploy: build prepare apply output

local: local-cleanup deploy

output:
	mkdir -p ../output
	pulumi stack output --json > ../output/pulumi-bucket.json
