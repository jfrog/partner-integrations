# Example Go

Example go app to demo JFrog partner integrations with DataDog, Jira cloud, PagerDuty, and Slack.

## Development

For available `make` commands run:

```bash
make usage
```

To `go build` module locally run:

```bash
export BUILD_NAME=go-example-build-info-dev
export BUILD_NUMBER=1
export GO_VIRTUAL_REPO=go-example-virtual
make build
```

To `go publish` module run:

```bash
export BUILD_NAME=go-example-build-info-dev
export BUILD_NUMBER=1
export GO_VIRTUAL_REPO=go-example-virtual
make publish
```

To `docker build` images run:

```bash
make image
```

This will generate the example-go image that is used by the helm chart.

## Datadog Agent

To deploy the Datadog agent helm chart run the following command:

`helm install datadog-agent -f values.yaml  --set datadog.apiKey=<DATADOG_API_KEY> datadog/datadog --set targetSystem=<TARGET_SYSTEM>`

Replace `<DATADOG_API_KEY>` with your DataDog API key.

Replace `<TARGET_SYSTEM>` with either `linux` or `windows`.

## Deployment

Helm chart is available under `helm/example-go`.

To deploy the example-go helm chart run the following command:

`helm install example-go helm/example-go`

This will deploy the Go application into Kubernetes.

## Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository.

## Contact

To contact JFrog about this reference material please create a Github issue.