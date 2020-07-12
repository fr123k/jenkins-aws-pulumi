@Library('shared-lib')

import org.jocker.setup.SetupWizard

def call(args) {
    println "${args.workspace}/jenkins/config/casc-config/"
    def setupWizard = new SetupWizard("${args.workspace}/jenkins/config/casc-config/")

    setupWizard
        .setup()
        .getCredsUtil()
            .AddUsernameAndPassword("aws-key-credentials", env.AWS_KEY_ID, env.AWS_KEY_SECRET)
}

return this
