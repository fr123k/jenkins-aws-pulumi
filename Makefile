export PULUMI_CONFIG_PASSPHRASE=test
export ADMIN_PASSWORD ?= $(shell pwgen -s 16 1)

init:
	mkdir -p output

deploy:
	$(MAKE) -C pulumi-bucket deploy
	$(MAKE) -C jenkins-ec2 deploy

cleanup:
	$(MAKE) -C pulumi-bucket cleanup
	$(MAKE) -C jenkins-ec2 cleanup

local:
	$(MAKE) -C pulumi-bucket local
	$(MAKE) -C jenkins-ec2 local

test: ## Wait 300 seconds (5 minutes) but check every 10 seconds if resource is available and then check the build status of the Configure job to fail if status is not SUCCESS.
	timeout 300 bash -c 'while [[ "$$(curl -s -o /dev/null -w ''%{http_code}'' http://$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/login)" != "200" ]]; do sleep 10; done' || false
	sleep 90 # wait until jenkins job is finished
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Configure/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Jobs/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Setup/lastBuild/consoleText
	@curl -s http://admin:$(ADMIN_PASSWORD)@$(shell cat ./output/jenkins-ec2.json | jq -r ."publicDns"):80/job/Jenkins/job/Setup/lastBuild/api/json | jq -r .result | grep SUCCESS
