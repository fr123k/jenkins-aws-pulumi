views = [
    [
        name: "Pulumi",
        description: "Pulumi Project",
        projects: [
            "jenkins-aws-pulumi",
        ],
    ],
]

for(view in views) {
    // for(project in view.projects) {
    //     folder("${project}") {
    //         description("${project} jobs")
    //     }
    // }

    listView("${view.name}") {
        description("${view.description}")
        filterBuildQueue()
        filterExecutors()
        jobs {
            for(project in view.projects) {
                name("${project}")
            }
        }
        columns {
            status()
            weather()
            name()
            lastSuccess()
            lastFailure()
            lastDuration()
            buildButton()
        }
    }
}
