export PULUMI_CONFIG_PASSPHRASE=test
export ADMIN_PASSWORD ?= $(shell pwgen -s 16 1)

export SEED_BRANCH_JOBS ?= $(shell [ -z "${TRAVIS_PULL_REQUEST_BRANCH}" ] && echo "${TRAVIS_BRANCH}"|| echo "${TRAVIS_PULL_REQUEST_BRANCH}")
export SEED_BRANCH_JOBS_URL=$(shell echo "${SEED_BRANCH_JOBS}" | sed s/\\//\\%2F/g)

init:
	mkdir -p output

deploy:
	$(MAKE) -C pulumi-bucket deploy
	$(MAKE) -C jenkins-ec2 deploy

cleanup:
	$(MAKE) -C jenkins-ec2 cleanup
	$(MAKE) -C pulumi-bucket cleanup

local:
	$(MAKE) -C jenkins-ec2 local-cleanup
	$(MAKE) -C pulumi-bucket local
	$(MAKE) -C jenkins-ec2 local

shell:
	$(MAKE) -C jenkins-ec2 shell

test: ## Wait 300 seconds (5 minutes) but check every 10 seconds if resource is available and then check the build status of the Configure job to fail if status is not SUCCESS.
	timeout 300 bash -c 'while [[ "$$(curl -s -o /dev/null -w ''%{http_code}'' http://$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/login)" != "200" ]]; do sleep 10; done' || false
	sleep 90 # wait until jenkins job is finished
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Configure/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Jobs/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Setup/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Setup/lastBuild/api/json | jq -r .result | grep SUCCESS

jenkins-aws:
	./scripts/jenkins-cli.sh $(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns") Pulumi/job/${SEED_BRANCH_JOBS_URL} $(ADMIN_PASSWORD) 90
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Pulumi/indexing/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Pulumi/job/${SEED_BRANCH_JOBS_URL}/lastBuild/consoleText

	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Pulumi/job/${SEED_BRANCH_JOBS_URL}/lastBuild/api/json | jq -r .result | grep SUCCESS
