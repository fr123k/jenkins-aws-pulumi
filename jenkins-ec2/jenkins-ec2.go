package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

const size = "t2.large"

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		config := config.New(ctx, "")

		awsKeyID := config.Require("key")
		awsKeySecret := config.Require("secret")
		return createJenkinsVM(ctx, awsKeyID, awsKeySecret)
	})
}

func getCloudInitYaml(fileName string, awsKeyID string, awsKeySecret string) (*string, error) {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return nil, err
	}
	yaml := parseCloudInitYaml(string(b), awsKeyID, awsKeySecret)
	return &yaml, nil
}

func parseCloudInitYaml(content string, awsKeyID string, awsKeySecret string) string {
	adminPassword, ok := os.LookupEnv("ADMIN_PASSWORD")
	var result string
	if ok == true {
		result = strings.ReplaceAll(content, "{{ ADMIN_PASSWORD }}", "ADMIN_PASSWORD="+adminPassword)
	} else {
		result = strings.ReplaceAll(content, "{{ ADMIN_PASSWORD }}", "")
	}
	seedBranchJobs, ok2 := os.LookupEnv("SEED_BRANCH_JOBS")
	if ok2 == true {
		result = strings.ReplaceAll(result, "{{ SEED_BRANCH_JOBS }}", "SEED_BRANCH_JOBS="+seedBranchJobs)
	} else {
		result = strings.ReplaceAll(result, "{{ SEED_BRANCH_JOBS }}", "SEED_BRANCH_JOBS=origin/master")
	}
	result = strings.ReplaceAll(result, "{{ AWS_KEY_ID }}", "AWS_KEY_ID="+awsKeyID)
	result = strings.ReplaceAll(result, "{{ AWS_KEY_SECRET }}", "AWS_KEY_SECRET="+awsKeySecret)

	return result
}

func createJenkinsVM(ctx *pulumi.Context, awsKeyID string, awsKeySecret string) error {
	group, err := ec2.NewSecurityGroup(ctx, "jenkins-security-group", &ec2.SecurityGroupArgs{
		Description: pulumi.String("Enable HTTP access"),
		Ingress: ec2.SecurityGroupIngressArray{
			ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			ec2.SecurityGroupIngressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("95.90.244.46/32")},
			},
		},
		Egress: ec2.SecurityGroupEgressArray{
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(80),
				ToPort:     pulumi.Int(80),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(443),
				ToPort:     pulumi.Int(443),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("192.30.252.0/22")},
			},

			//github ssh
			//TODO allow ssh to any host until can retrieve all ssh ip address of github.com
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("0.0.0.0/0")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("140.82.118.0/24")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("140.82.121.4/32")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("204.232.175.90/32")},
			},
			ec2.SecurityGroupEgressArgs{
				Protocol:   pulumi.String("tcp"),
				FromPort:   pulumi.Int(22),
				ToPort:     pulumi.Int(22),
				CidrBlocks: pulumi.StringArray{pulumi.String("207.97.227.239/32")},
			},
		},
	})
	if err != nil {
		return err
	}

	mostRecent := true
	//TODO check if jenkins master jocker ami exists use it otherwise use this one.
	//make this behaviour configurable always use the following ami except following cases
	// 1) jenkins jocker ami exists 2) 1) && env var JENKINS_AMI=ami
	ami, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
		Filters: []aws.GetAmiFilter{
			{
				Name:   "name",
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-20200129"},
			},
		},
		Owners:     []string{"099720109477"},
		MostRecent: &mostRecent,
	})

	if err != nil {
		return err
	}

	//TODO cloud-init use only if jenkins ami doesn't exists.
	yaml, err := getCloudInitYaml("cloud-init/cloud-init.yaml", awsKeyID, awsKeySecret)

	ctx.Export("cloud-init", pulumi.String(*yaml))

	if err != nil {
		return err
	}

	server, err := ec2.NewInstance(ctx, "jenkins-master", &ec2.InstanceArgs{
		Tags:         pulumi.StringMap{
			"Name": pulumi.String("jenkins-master"),
			"JobUrl": pulumi.String(os.Getenv("TRAVIS_JOB_WEB_URL")),
		},
		InstanceType: pulumi.String(size),
		SecurityGroups: pulumi.StringArray{
			group.Name,
		},
		KeyName:  pulumi.String("development"), //create the keypair with pulumi
		Ami:      pulumi.String(ami.Id),
		UserData: pulumi.String(*yaml),
	})

	ctx.Export("publicIp", server.PublicIp)
	ctx.Export("publicDns", server.PublicDns)

	return err
}
