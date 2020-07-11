export STACK_NAME ?= jenkins-ec2

go-init:
	go mod init main

build:
	go build -o $(shell basename $(shell pwd))

prepare: build
	pulumi plugin install resource aws 0.18.3
	pulumi plugin ls
	pulumi login --local
	# pulumi stack rm -f ${STACK_NAME}
	# pulumi stack select ${STACK_NAME}
	pulumi stack init ${STACK_NAME}
	pulumi config set aws:region eu-west-1
	@pulumi config set --secret key $(shell cat ../output/pulumi-bucket.json | jq -r ."AccessKeys")
	@pulumi config set --secret secret $(shell cat ../output/pulumi-bucket.json | jq -r ."AccessKeysSecret")

apply:
	pulumi up --yes
	#verbose logging
	#pulumi up --yes --verbose 9 --logtostderr

cleanup:
	pulumi destroy --yes -s ${STACK_NAME}
	pulumi stack rm -f --yes ${STACK_NAME}

local-cleanup:
	echo "ADMIN_PASSWORD = ${ADMIN_PASSWORD}"
	pulumi destroy --yes -s ${STACK_NAME} || true
	pulumi stack rm -f --yes ${STACK_NAME} || true

deploy: build prepare apply output

local: local-cleanup deploy

shell:
	pulumi stack output publicDns
	ssh -i "~/.ssh/development.pem" -vvvv ubuntu@$(shell pulumi stack output publicDns)

browse:
	pulumi stack output publicDns
	open http://$(shell pulumi stack output publicDns)

output:
	mkdir -p ../output
	pulumi stack output --json > ../output/jenkins-ec2.json