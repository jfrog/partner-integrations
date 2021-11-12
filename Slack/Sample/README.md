# Xray Summary view sample code

This is the sample code for the [Implementing the JFrog Xray “Summary View” in Slack](http://dev.to) post on DEV.to. This provides a simple way to execute the sample code and transform the Xray webhook payload into Slack message, and see how Xray payload is summarized automatically when too many issues are encountered.

## Installation

(Install Node/NPM first if you don't already have it)

Install NPM package dependencies

```sh
$ npm install
```

## Code Structure

```
index.js
formatter\
  index.js
  issue\
    common.js
    index.js
    license.js
    security.js
  summary\
    common.js
    index.js
    license.js
    security.js
```

### index.js

The [`index.js`](index.js) is the entry point to the `formatter` module. It passes the payload data to the formatter, along with the format type.

### Formatter Module

The [`formatter` module](./formatter/index.js) normalizes the various type of issues in the payload into a single structure. Then looks up the appropriate format mapper based on the format type and violation type (security or license).

### Mappers

Each format type has mappers for each issue type (`license` and `security`). The mapper transforms the normalized issues into Slack message format.

## Run script

```sh
$ node index.js <path to xray webhook payload JSON file>
```

We have provide a few examples of Xray webhook payload in the `payloads` directory. You can use the JSON captured from your own Xray webhook too.

For example, using [payloads/security-multiple.json](./payloads/security-multiple.json):

```sh
$ node index.js payloads/security-multiple.json
```

To control which format type the code to transform the payload, you can use the optional argument (`issue` or `summary`). Script will default to `summary` if omitted.

This will process the `security-multiple.json` payload using the `issue` format.

```sh
$ node index.js payloads/security-multiple.json issue
```

# Visualize Slack Message

The output of the script is a Slack message in JSON using Slack's BlockKit.

You can visualize the message using Slack's [BlockKit Builder](https://app.slack.com/block-kit-builder)
