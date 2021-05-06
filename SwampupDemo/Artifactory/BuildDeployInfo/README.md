
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

## Build Info Setup

The JFrog CLI must be configured properly with an issue tracker CLI configuration file for this demo to work.

See [this blog post](https://jfrog.com/blog/using-jfrog-cli-to-see-your-builds-up-close/) with steps to guide you through the JFrog CLI setup required.

## Git Setup

Before we build `example-go` we need to ensure we have a `git commit` with our issue tracker ticket number.

Example git commit:

`git commit -m 'JFROG-1 - Commit with issue tracker id`

Once we have a `git commit` with our issue tracker id this information will be available in the Artifactory build info json.

## Builds & Build Info

Now that we have Artifactory and our JFrog CLI configured we can build the `example-go` application in this project.

To build `example-go` we can use the `Makefile` provided and run the following command:

```bash
export BUILD_NUMBER=1
export BUILD_NAME=my-example-go-build-name
make build
```

This will compiled the `example-go` module and produce the Go binary.

To publish `example-go` build info

```bash
make publish
```

## Jira Deploy Info

Deploy information can also be added to the Artifactory build info when the build info is published.

Deployments rely upon the following Build Info environment variables being set to track where the build was deployed into.

### Deploy Info Environment Variables:

| Environment Variable  | Description                                         |
|-----------------------|-----------------------------------------------------|
| JIRA_ENVIRONMENT_ID   | The identifier of this environment                  |
| JIRA_ENVIRONMENT_NAME | The name of the environment to present to the user. |
| JIRA_ENVIRONMENT_TYPE | The type of environment this was deployed into.     |
| JIRA_DEPLOYMENT_STATUS| The state of the deployment                         |
                         
                      
###### JIRA_ENVIRONMENT_TYPE                
```text
Valid values: unmapped, development, testing, staging, production
```                  
                         
###### JIRA_DEPLOYMENT_STATUS                
```text
Valid values: unknown, pending, in_progress, cancelled, failed, rolled_back, successful
```
                          

To deploy `example-go` into environment you can run:

```bash
export BUILD_NUMBER=2                                       # Create new deployment build number
export BUILD_NAME=my-example-go-build-name                  # Use existing build info name from build above
export JIRA_ENVIRONMENT_ID=MyDevelopmentEnvironmentID       # Example development environment id
export JIRA_ENVIRONMENT_NAME=MyDevelopmentEnvironmentName   # Example development environment name
export JIRA_DEPLOYMENT_STATUS=successful                    # Successful deployment status
export JIRA_ENVIRONMENT_TYPE=development                    # Deployment environment type
make publish                                                # Collect VCS, Envs, Publish Build Info
```


## Docker

To build the `example-go` app into a docker image the following environment variables must be supplied:

| Environment Variable    | Description                                         |
|-------------------------|-----------------------------------------------------|
| ARTIFACTORY_URL         | The URL of your Artifactory instance.               |
| ARTIFACTORY_DOCKER_REPO | The name of the Docker repo to push image into.     |
| VERSION                 | The version/tag to apply to the docker image.       |

Then to build the docker image run:

`make image`

To publish the docker image to Artifactory run:

`make push`

## Datadog APM and Traces

Datadog APM and tracing has been added to this example-go application.

For more information on how to setup Datadog APM for Golang [visit here.](https://docs.datadoghq.com/tracing/)
