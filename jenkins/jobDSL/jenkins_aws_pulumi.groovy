pipelineJob("jenkins-aws-pulumi") {
    definition {
        cps {
            script("""
node ("master") {
    sh("echo jenkins-aws-pulumi")
}""")
        }
    }
}
