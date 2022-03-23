# Welcome to the JFrog Slack App

This integration is between JFrog Artifactory and Xray and Slack. 

We know that software development happens in a myriad of collaboration environments. Today, there are key events throughout the JFrog Platform that can be difficult for a user to interact with if they aren't logged into the platform. When it comes to people across the organization knowing what’s going on, this Slack integration provides users real-time information about Artifactory and Xray events. This will give each user situational awareness about occurrences in the JFrog Platform. Additionally, where appropriate - they will have links and action buttons to follow-up on the event.

# How it Works

* This integration allows you the ability to see Artifact, Artifact Properties, Docker Tag, and Build events through notifications and actionable UI cards inside Slack.
* You can  send these notifications to multiple channels.
* Additionally, you can get vulnerability and license compliance notification based on policies setup in JFrog Xray and take actions. Here is an example of a security violation and adding an ignore rule to snooze the notification (happens in Slack and in the JFrog Platform):

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image4.png?raw=true" width="800">

# Support

If you need help with this integration, please contact `partner-support@jfrog.com`

# Getting Started

If you do not already have Slack, **download it now** ([for Windows](https://slack.com/help/articles/209038037-Download-Slack-for-Windows) or [for Mac](https://slack.com/help/articles/207677868-Download-Slack-for-Mac)).

Next, **[Install the JFrog Slack Application using this link](https://slack-connector.jfrog.io/webhooks/slack/oauth/install-url/?utm_source=slack&utm_medium=wiki)**.

Make sure your organization has the latest cloud instance of Artifactory. If not, please upgrade or you can signup for a free cloud instance at: [https://jfrog.com/start-free/](https://jfrog.com/start-free/).

You can scroll down to see the Requirements for using this integration.

# Next: Connect Your JFrog Platform Deployment (JPD)

Once you’ve installed the JFrog Slack Application from the Slack Store, you must connect the application to your JFrog Platform Deployment (or JPD) instance.

First, login to the JFrog platform. You must be an admin to access your credentials.

A) If you are an JFrog Platform paid user, go to **General** tab and look for **Applications**. From there, you can add a new application.

Then click on **Next, Generate your ID and Secret.**.

B) If you are using the free version of the JFrog Platform, go to the **Administration** section and then click on **Security**. Next, scroll down and click on **Integrations**.

Then click on **+ New Client Integration**. Provide it a name and then from the dropdown under Application, select Slack.

Then click on **Next, Generate your ID and Secret.**.

When you have your **Integration ID and Secret**, copy and paste these items into the Slack modal window.

Last, copy and paste your JFrog Platform URL (found at the top of your browser window) into the Slack modal window where it says **JFrog URL**.

Hit Save and look for the success message! Great, you have now connected your organization’s JFrog account to the Slack app.

# Login

Next, login to your JFrog account on the Slack app.

Once you see the confirmation message, you can create new notifications and add them to Slack channels.

We expect you to have Slack channels already setup - how you want to organize notifications to different Slack channels is up to your organization.

# Creating Notifications

Once logged in with Admin privileges, you can start creating notifications.

Hit **Create Notification** to bring up the list of options.

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/imagec.png?raw=true" width="1000">

Select which type notification you would like to create from the drop-down menu.

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image2.png?raw=true" width="600">

On the next screen, name the notification and select which events you would like to include in the notification, and which repos should be included. 

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image8n.png?raw=true" width="600">

You select a channel to send the notifications to. Hit **Next**.

On the next screen, you may see options for your notification. For example, for build notifications, you can select any build or find an existing build (in your JFrog Platform Deployment) by name or pattern.

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image9n.png?raw=true" width="600">

Once you hit **Next**, you should see a success message.

Once you have setup notifications, you should start seeing the notification cards in the channel within about 20 minutes. If you do not see notifications working, first type **/jfrog rt list** to bring up the list of active notifications. 

If you see nothing on the notification list, please try again or contact support.

# Types of Supported Notifications

All notifications are based on webhook events in the JFrog Platform. The currently supported notifications include:

### Artifactory

Type | Events
------------ | -------------
Artifact | *deployed, deleted, moved, copied*
Artifact Properties | *added, deleted*
Docker Tag | *pushed, deleted, promoted*
Builds | *uploaded, deleted, promoted*
Release Bundles (Enterprise+) | *created, signed, deleted*
Distribution (Enterprise+) | *stared, completed, aborted, failed, deletion started, deletion failed, deletion completed*

### Xray

Type | Description
------------ | -------------
Security Violations by CVE | *This sends individual notification cards for each CVE or issue*
Security Violations by Component (Summary view) | *This provides a summary of all CVEs and severities by component*
License Compliance | *This sends individual notification cards for each license compliance issue*

# Xray Notifications

JFrog Xray notifications are special in that only repositories that are being actively watched in Xray and have a policy setup will generate notification events. To learn more about how Xray policies and watch work, [click here](https://www.youtube.com/watch?v=88hwwMJsS58). 

If you already have policies and watches setup in Xray, you can create notifications in the Slack app. 

Hit Create Notification.

Give the Notification a name, which policy it is coming from, which channel to send the notification to, and whether you would like the security violation to send you information by individual CVE or send a grouped notification by Component (Summary).

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image6.png?raw=true" width="600">

**Example - by CVE:**

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image1.png?raw=true" width="800">

**Example - by Component:**

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image10n.png?raw=true" width="800">

## Additionally:

All notifications can be paused, which removes them from being active in Slack channels, but not does not delete the underlying webhook so they can be added again.

<img src="https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/images/image3.png?raw=true" width="600">

The delete notification button deletes the entire notification from Slack as well as the underlying webhook in the JFrog Platform. 

# List of Shortcuts

Right under the text input area in Slack, you will see a lightning bolt symbol. Clicking on it brings up a list of 5 shortcuts that you can also use to create notifications and view lists. These shortcuts are:

**Create Notifications**

**List XR Notifications**

**List RT Notifications**

**List Policies**

**List Watches**

# List of Commands

Outside the UI elements, you can also interact with our application using commands in the Slack chat area. The commands we currently support are:

### General

*/jfrog help* - Show help content

*/jfrog configure* - Connects to your JFrog Instance and asks for JFrog URL, Integration ID, and Integration secret

*/jfrog logout* - Log out from the JFrog Platform with this Slack app

### Artifactory

*/jfrog rt notify list* - Provides a list of current notifications subscribed to by the personal or channel

*/jfrog rt notify stop {notification name}* - Pauses subscription to the specified notification

*/jfrog rt notify resume {notification name}* - Resumes subscription to the specified notification

### Xray

*/jfrog xr notify list* - Provides a list of current xray notifications subscribed to by the personal or channel bot context

*/jfrog xr watch list* - Provides a list of current watches (that user can has read access to) with a micro action to subscribe the bot to the notification

*/jfrog xr policy list* - Provides a list of current policies with a micro action to open JFrog Platform.

# Requirements 

* Your organization has a cloud instance of Artifactory. You can signup for a free cloud instance at: [https://jfrog.com/start-free/](https://jfrog.com/start-free/)

* You must be a user with Admin permissions to authenticate your organization’s Slack app with your JFrog Platform Deployment (JPD).

* You must be a user with Admin permissions to create the initial notifications for Artifactory and Xray. Once created, any team member can add existing notifications to new Slack channels.

* Your organization must already have setup policies and watches prior to getting Xray notifications in Slack. [Learn how to setup watches and policies in Xray](https://www.youtube.com/watch?v=88hwwMJsS58).

You can also click here for a list of [FAQ's](https://github.com/jfrog/partner-integrations/blob/main/Slack/Platform/faq.md).
