
# JFrog Platform Build Deploy Info Demo

## Purpose
Utilize JFrog Artifactory and Partner integrations (Slack, Pagerduty, Datadog, Jiracloud) to create a robust end to end Security information event management build pipeline to enrich your Jira ticket with build and deployment information from Artifactory.

## Artifactory Setup
Artifactory is required and we recommend following the below best practices highlighted in this demo:

```text
1. Separation of dev/test/prod artifact repos
2. Usage of build info and build promotion in Artifactory
3. Artifact property usage for labels/tags on artifacts
```

Artifact repos by type must be properly separated by dev/test/prod to enable our demo flow to monitor impacted environment types for an infected component.



## Putting it all together

