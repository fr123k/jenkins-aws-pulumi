export PULUMI_CONFIG_PASSPHRASE=test

go-init:
	go mod init main

build:
	go build

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

cleanup-aws-pulumi:
	pulumi destroy --yes
	pulumi stack rm -f --yes jenkins-aws-s3

deploy: build deploy-aws-pulumi cleanup-aws-pulumi
