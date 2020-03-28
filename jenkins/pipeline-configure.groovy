node('master') {
    stage('Checkout') {
        cleanWs()
        checkout scm
    }

    stage('Configure') {
        // run SetupWizard from fr123k/jenkins-shared-library
        def setupScript = load('jenkins/config/groovy/setup.groovy')
        def setupWizard = setupScript()

        stage('Agents') {
            setupWizard.setup()
                //the pipeline script in the jobDSL/pulumi.groovy
                .getScriptApproval().approveScript('6d2ccc5267db0f3b500aa96a1ec53264613a1209')
        }

        stage('Credentials') {
            setupWizard.getCredsUtil()
                .AddUsernameAndPassword("aws-key-credentials", env.AWS_KEY_ID, env.AWS_KEY_SECRET)
        }
    }

    stage('Seed') {
        // echo "create Jobs based repository: ${repository}, revision: ${jobDSLRevision}, jobDSLPath: ${jobDSLPath}, removedJobAction ${removedJobAction}"
        jobDsl(
            targets: "jenkins/jobDSL/*.groovy",
            sandbox: false,
            removedJobAction: 'IGNORE',
            lookupStrategy: 'JENKINS_ROOT',
        )
    }
}
