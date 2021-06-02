# JIRA Cloud + JFrog Artifactory Integration

* This integration creates an automated connection that delivers your JFrog Artifactory repository build information directly to your JIRA issues. JFrog Artifactory supports build integrations whether you are running builds on one of the common CI servers in use today, on cloud-based CI servers, or standalone without a CI server.

* Integration of Artifactory into your build ecosystem provides important information that supports fully reproducible builds through visibility of artifacts deployed, dependencies and information on the build environment. 

* With this feature, your team gets immediate feedback for continuous improvement that helps fulfill CI/CD needs and helps support a "shift left" strategy that allows developers to respond to problems as soon as they are known.

# How it Works

* This integration allows you the ability to see what JIRA issues went into a build (and eventually a release) to enhance the traceability information provided inside Artifactory.  

* With this, your development team gets more visibility to see what build is associated with a specific feature, and exposes in JIRA where a specific issue is inside the release process by providing the promotion status of the build.

* All of this together provides critical traceability of an artifact to the build that generated that artifact, and the information needed to reproduce that build.

# Requirements

* You must be a user with Admin permissions to generate an API Key
* You should have a Jira Cloud instance. You can signup for a free instance at: https://www.atlassian.com/software/jira/pricing

# Support
If you need help with this integration, please contact `partner_support@jfrog.com`

# How to connect your Artifactory cloud instance to JIRA Cloud

You can [watch the video](https://drive.google.com/file/d/1l_oBI5m66uPvv2woZtAZpS2l9CRhLXAQ/view) or **follow the steps below**:

1. First, visit the [Atlassian Marketplace](https://marketplace.atlassian.com/search?query=JFrog) and search for **JFrog**.

![Atlassian Marketplace Search](https://github.com/Dattax/partner-integrations/blob/main/img/img1.png?raw=true)

2. You should see the **JFrog Platform: Artifactory App for Jira**. Click on the icon and on the top-right, hit the **Get it now** button.

![Atlassian Marketplace Search](https://github.com/Dattax/partner-integrations/blob/main/img/img1.png?raw=true)

This will install the JFrog app into your JIRA instance. Once installed, a pop-up should appear on the top-right of your screen. Click on **Get started**.

You should see a screen asking for your **Artifactory Cloud URL** and **API Key**. 

At this point, you'll need to make sure you are logged into your Artifactory cloud instance.

Copy the first part of the URL for your Artifactory cloud instance and paste it into the JIRA Initilization screen for **Artifactory Cloud URL**.

Next, to find the API Key, you'll need to have Admin permissions.

Assuming you are an Admin, click on the top right where your name is, and then **edit profile**, and there will be an API Key section. You may need to generate a new API Key if you have never generated one. 

At this point, copy the **API Key** into the JIRA Initilization screen and click the **Submit** button. Once you get the notification that the initializiation was successful, you can go back into your JIRA Cloud instance.

# Uploading a Build to Artifactory

One way to add build information to Artifactory is manually uploading a **buildinfo.json** to your Artiactory Cloud instance. 

In this example, the important thing to note is the  **issues** and then **affectedIssues** keys in your json file. This is how this information will be associated to your issue in JIRA.

Next, upload the **buildinfo.json** to Artifactory and inside your JFrog Platform instance, you can navigate to **Builds** and you should see the build information you just uploaded. 

Now, navigate to the **Issues** tab and you will see a direct link to your **JIRA Cloud issue**. 

You'll now notice that issue has been updated.

# How to Uninstall

To uninstall, use the **Uninstall** button in Manage apps in JIRA Cloud:

![Uninstall app](https://github.com/Dattax/partner-integrations/blob/main/img/jira_image.png?raw=true)




