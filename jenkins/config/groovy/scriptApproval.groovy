#! groovy
import static org.jenkinsci.plugins.scriptsecurity.scripts.ScriptApproval.get

// approve the jenkins-aws-pulumi pipeline script
get().approveScript('a9581ecb9b30bcd45b80e42c56d793abf51360ec')
