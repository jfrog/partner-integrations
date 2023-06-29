# JIRA Cloud + JFrog Artifactory Integration

This integration creates an automated connection that delivers your JFrog Artifactory repository build information directly to your JIRA issues. JFrog Artifactory supports build integrations whether you are running builds on one of the common CI servers in use today, on cloud-based CI servers, or standalone without a CI server.

Integration of Artifactory into your build ecosystem provides important information that supports fully reproducible builds through visibility of artifacts deployed, dependencies and information on the build environment.

With this feature, your team gets immediate feedback for continuous improvement that helps fulfill CI/CD needs and helps support a "shift left" strategy that allows developers to respond to problems as soon as they are known.

## How it Works

This integration allows you the ability to see what JIRA issues went into a build (and eventually a release) to enhance the traceability information provided inside Artifactory.  

With this, your development team gets more visibility to see what build is associated with a specific feature, and exposes in JIRA where a specific issue is inside the release process by providing the promotion status of the build.

All of this together provides critical traceability of an artifact to the build that generated that artifact, and the information needed to reproduce that build.

## On-premise JFrog Installation

To use the integration with an on-premise JFrog installation, the network port `8082` will need to be exposed to the external network. See [JFrog System Requirements](https://www.jfrog.com/confluence/display/JFROG/System+Requirements#SystemRequirements-RequirementsMatrix) for more information.

Then when initializing the integration, include the network port in the Artifactory Cloud URL, e.g. `https://example.com:8082`

## Requirements

* You must be a user with Admin permissions to generate an API key or scoped token
* Your organization has an instance of Artifactory. You can signup for a free instance at: [https://jfrog.com/start-free/](https://jfrog.com/start-free/)

## Limitation

JIRA issue accepts only **whole number** as build number (e.g. 1, 2, 999, etc.). Therefore, common build/release versioning schemes such as SemVer (`1.2.3`), date-based (`2023.06.28.999`), or even empty text will **not** be accepted by JIRA.

If your build numbering scheme falls into these categories, please be aware that this integration will **not** function as intended.

### Workarounds

One possible workaround will be to use the build/release version as build name and the least significant part of the version as build number.

For example for SemVer:
* build name = `1.2.3`
* build number = `3`

For date-based version:
* build name = `2023.6.28.999`
* build number = `999`

## Support

If you need help with this integration, please contact [support@jfrog.com](mailto:support@jfrog.com)

## How to connect your Artifactory cloud instance to JIRA Cloud

You can [watch the video](https://drive.google.com/file/d/1l_oBI5m66uPvv2woZtAZpS2l9CRhLXAQ/view) or **follow the steps below**:

1. First, visit the [Atlassian Marketplace](https://marketplace.atlassian.com/search?query=JFrog) and search for **JFrog**.

![Atlassian Marketplace Search](https://github.com/jfrog/partner-integrations/blob/main/img/img1.png?raw=true)

2. You should see the **JFrog Platform: Artifactory App for Jira**. Click on the icon and on the top-right, hit the **Get it now** button.

![Integration Screenshot](https://github.com/jfrog/partner-integrations/blob/main/jira/images/image3.png?raw=true)

3. This will install the JFrog app into your JIRA instance. Once installed, a pop-up should appear on the top-right of your screen. Click on **Get started**.

![Integration Screenshot](https://github.com/jfrog/partner-integrations/blob/main/jira/images/image2.png?raw=true)

4. You should see a screen asking for your **Artifactory Cloud URL** and **API Key**.

![Integration Screenshot](https://github.com/jfrog/partner-integrations/blob/main/jira/images/image5.png?raw=true)

5. At this point, you'll need to make sure you are logged into your Artifactory cloud instance. Copy the first part of the URL for your Artifactory cloud instance and paste it into the JIRA Initialization screen for **Artifactory Cloud URL**.

6. Next, to create a scoped token, you'll need to have Admin permissions.

7. As an admin, log into your Artifactory SaaS instance, click on gear icon near top left, open “User Management”, click on “Access Tokens”. Click on “Generate Token" at top right, click “Generate”.

![Integration Screenshot](https://github.com/jfrog/partner-integrations/blob/main/jira/images/image1.png?raw=true)

8. Now, copy the **token** into the JIRA Initialization screen and click the **Submit** button.

![Integration Screenshot](https://github.com/jfrog/partner-integrations/blob/main/jira/images/image6.png?raw=true)

Once you get the notification that the initialization was successful, you can go back into your JIRA Cloud instance.

## Uploading a Build to Artifactory

There are two ways to upload a build information to Artifactory:
1. Upload a `buildinfo.json` file
2. Send a API request to your Artifactory instance

### Upload buildinfo.json file

To add build information to Artifactory, manually uploading a `buildinfo.json` to your Artiactory Cloud instance.

In this example, the important thing to note is the  `issues` and then `affectedIssues` keys in your json file. This is how this information will be associated to your issue in JIRA.

Next, upload the `buildinfo.json` to Artifactory and inside your JFrog Platform instance, you can navigate to **Builds** and you should see the build information you just uploaded.

Now, navigate to the **Issues** tab and you will see a direct link to your **JIRA Cloud issue**.

You'll now notice that issue has been updated.

### Make API request

You can also add the build information to Artifactory by making a request to the [build upload API](https://jfrog.com/help/r/jfrog-rest-apis/build-upload)

Example:
```sh
curl -L -X PUT 'https://<your artifactory hostname>/artifactory/api/build' -H 'Content-Type: application/json' -H 'Authorization: Bearer <your auth token>' -d '{
    "version": "1.0",
    "name": "build-test",
    "number": "1",
    "started": "2023-06-15T16:24:42.116+0300",
    "modules": [
        {
            "id": "org.jfrog.test:multi:2.36-2-SNAPSHOT",
            "artifacts": [
                {
                    "type": "pom",
                    "sha1": "a1ab9222c219ec885ce4374b77cc4bd93485c6c8",
                    "md5": "495a5e916f0e308a8ed4fc170e4d9136",
                    "name": "multi-2.36-2-SNAPSHOT.pom"
                }
            ]
        }
    ],
    "issues": {
        "tracker": {
            "name": "JIRA",
            "version": "6.0.1"
        },
        "aggregateBuildIssues": false,
        "aggregationBuildStatus": "Released",
        "affectedIssues": [
            {
                "key": "JT-5",
                "url": "https://jiracloud-connector-test.atlassian.net/browse/JT-5",
                "summary": "Dev testing",
                "aggregated": false
            }
        ]
    }
}'
```

## How to Uninstall

To uninstall, use the **Uninstall** button in Manage apps in JIRA Cloud:

![Uninstall app](https://github.com/jfrog/partner-integrations/blob/main/img/jira_image.png?raw=true)
