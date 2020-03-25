node('master') {
    stage('Checkout') {
        cleanWs()
        checkout scm
    }

    stage('Configure') {
        // run SetupWizard from fr123k/jenkins-shared-library
        def setupScript = load('jenkins/config/groovy/setup.groovy')
        def setupWizard = setupScript()

        stage('Credentials') {
            setupWizard.getCredsUtil()
                .AddUsernameAndPassword("aws-key-credentials", env.AWS_KEY_ID, env.AWS_KEY_SECRET)
        }

        stage('Agents') {
            setupWizard.setup()
                //the pipeline script in the jobDSL/pulumi.groovy
                .getScriptApproval().approveScript('6d2ccc5267db0f3b500aa96a1ec53264613a1209')
        }
    }

    stage('Seed') {
        // https://issues.jenkins-ci.org/browse/JENKINS-44142
        // --> Note: when using multiple Job DSL build steps in a single job, set this to "Delete" only for the last Job DSL build step. 
        // Otherwise views may be deleted and re-created. See JENKINS-44142 for details.
        jobDsl(targets: 'jenkins/jobDSL/*.groovy', sandbox: false, removedJobAction: 'DELETE')
    }
}
