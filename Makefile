export PULUMI_CONFIG_PASSPHRASE=test
ADMIN_PASSWORD ?= $(shell pwgen -s 16 1)

go-init:
	go mod init main

build:
	go build -o $(shell basename $(shell pwd))

prepare-aws-pulumi: build
	pulumi plugin install resource aws 0.18.3
	pulumi plugin ls
	pulumi login --local
	# pulumi stack rm -f jenkins-aws-s3
	# pulumi stack select jenkins-aws-s3
	pulumi stack init jenkins-aws-s3
	pulumi config set aws:region eu-west-1

deploy-aws-pulumi:
	pulumi up --yes
	#verbose logging
	#pulumi up --yes --verbose 9 --logtostderr

cleanup-aws-pulumi:
	pulumi destroy --yes
	pulumi stack rm -f --yes jenkins-aws-s3

local-cleanup-aws-pulumi:
	pulumi destroy --yes || true
	pulumi stack rm -f --yes jenkins-aws-s3 || true

deploy: build prepare-aws-pulumi deploy-aws-pulumi cleanup-aws-pulumi

local: local-cleanup-aws-pulumi build prepare-aws-pulumi deploy-aws-pulumi

shell:
	pulumi stack output publicDns
	ssh -i "~/Downloads/test.pem" -v ubuntu@$(shell pulumi stack output publicDns)

test: ## Wait 300 seconds (5 minutes) but check every 10 seconds if resource is available and then check the build status of the Configure job to fail if status is not SUCCESS.
	timeout 300 bash -c 'while [[ "$$(curl -s -o /dev/null -w ''%{http_code}'' http://$(shell pulumi stack output publicDns):80/login)" != "200" ]]; do sleep 10; done' || false
	sleep 90 # wait until jenkins job is finished
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell pulumi stack output publicDns):80/job/Jenkins/job/Configure/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell pulumi stack output publicDns):80/job/Jenkins/job/Jobs/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell pulumi stack output publicDns):80/job/Jenkins/job/Setup/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell pulumi stack output publicDns):80/job/Jenkins/job/Setup/lastBuild/api/json | jq -r .result | grep SUCCESS
