# Xray Summary view sample code

## Installation

Install NPM packages

```sh
$ npm install
```

## Run script

```sh
$ node index.js <path to event payload JSON file>
```

Example:

```sh
$ node index.js payloads/security-multiple.json
```

You can use the optional argument to control the format type. Default to 'summary' if omitted.

```sh
$ node index.js payloads/security-multiple.json issue
```

# Visualize Slack Message

The output of the script is Slack message in JSON using Slack's BlockKit.

You can visualize the message using Slack's [BlockKit Builder](https://app.slack.com/block-kit-builder)
