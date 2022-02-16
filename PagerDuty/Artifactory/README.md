# Welcome to the PagerDuty + JFrog Artifactory Integration. 

Usually, when making software updates to your binaries and applications, your development teams don’t always have direct access to the metadata surrounding events (example: during the compilation of that code). Getting this information is especially important when you have different teams in charge of different parts of your software delivery pipeline. Our integration with PagerDuty allows this monitoring to take place so your on-call teams can monitor events in real-time.

# How it Works
* Users will install the JFrog Artifactory integration application from within the PagerDuty Service Directory and receive a URL to setup the PagerDuty webhook within JFrog Artifactory
* JFrog Artifactory will utilize the PagerDuty webhook trigger to send notifications to PagerDuty.  Users will be able to configure their webhook events for repositories, builds and release bundles within JFrog Artifactory and associate these events with the PagerDuty webhook trigger. 
* Once the webhook is configured within JFrog Artifactory, an change event will be sent to the PagerDuty service whenever an event occurs.
* Artifactory performs a recheck of all selected repositories, builds and release bundles and any events generated will be sent as an event to a service in PagerDuty. Events from Artifactory will trigger a new change event on the corresponding PagerDuty service, or group as alerts into an existing recent change event.

# Requirements
* PagerDuty integration requires an admin base role for account authorization. If you do not have this role, please reach out to an Admin or Account Owner within your organization to configure the integration.
* JFrog Artifactory requires an Admin based role with administrative privileges to configure the integration.  If you do not have this role, please reach out to an Admin or Account Owner within your organization to configure the integration.

# Support
If you need help with this integration, please contact `partner_support@jfrog.com`

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

# Integration Walkthrough

## Pre-Step: PagerDuty Setup

First, if you don’t already have one, you’ll need to register an account with PagerDuty. Then, once logged into PagerDuty, you’ll hover over the People tab and do the following:

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image2.png" width="900">

Next, you'll need to take the following steps:
1. Add Users
2. Create an On-Call Schedule
3. Create an Escalation Policy 
4. Create a Service 

You can refer to the [PagerDuty Quick Start Guide](https://support.pagerduty.com/docs/quick-start-guide) for more information regarding each of these pre-steps to setup your on-call schedule and team within PagerDuty. 

## Add the JFrog Integration as New Service:

Add JFrog Artifactory integration to an existing PagerDuty service or create a new service by following the steps below:

First, you’ll need to create a new service. Navigate to **Services > Service Directory** and click **+New Service**.

On the next screen, give your new service a name:
<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image9.png" width="900">

Next, assign an escalation policy, or create a new one.

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image5.png" width="900">

Next, select if you’d like intelligent, content-based, or time-based grouping.

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image10.png" width="900">

Next, search for “jfrog” and select JFrog Artifactory Notifications. You’ll then click on the Create Service button.

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image12.png" width="900">

You’ll now see the following information. Please copy the **Integration URL** as you’ll be using it to create your webhooks in the JFrog Platform:

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image7.png" width="900">

Once, you have created a service inside PagerDuty, you’ll need to login to your JFrog instance to create webhooks. These webhooks will send event data to PagerDuty an serve as the basis for your alerts.

## Next, Setup Webhooks in JFrog Artifactory.

First, please login to your JFrog Platform instance. 

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image6.png" width="900">

Under the administration panel, click on **General > Webhooks**. Next, select **+ New Webhook** from the top-right corner.

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image4.png" width="900">

Next, provide a name to your new webhook. After providing the name, please copy the Integration URL from PagerDuty and paste it where it says URL here:

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image14.png" width="900">

Then, select the types of events to include. In this example, we’re creating a Build webhook notification:

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image8.png" width="900">

Now, select the repositories to include in this webhook notification:

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image13.png" width="900">

Hit **Save**.

At the bottom of the main screen, you can Test the notification, by selecting the **Test** button. 

This will send your webhook event to your PagerDuty integration. After you have tested your webhook, hit **Create** and you’re done!

Once you have tested and finish creating your webhooks, you can see them as recent changes in PagerDuty:

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/img/image15.png" width="900">

# How to Uninstall
1. Find the integration to delete by navigating to **Services** and selecting **Service Directory**
2. Select the service with your integration and navigate to **Integrations** tab
3. Navigate to the integration by clicking on the integration name (**JFrog Artifactory Notifications**)
4. Click on the **Delete Integration** button on the top right corner
