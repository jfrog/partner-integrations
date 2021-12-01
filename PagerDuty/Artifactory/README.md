# PagerDuty + JFrog Artifactory Integration Benefits
* **Incident control** during an incident, the user will be able to view recently completed jobs or builds that are relevant for the environment / area of stack. What versions are built, where they are being delivered. More information on Success / Failures
* **Monitoring** performance, think number of artifacts served per second, latency in artifactory response time, artifactory up time
* **Error Notifications** Alerting such that whenever there are errors
* **Service Status** Inform on-call team if an API or a web application has issues, goes unavailable etc

# How it Works
* Users will install the JFrog Artifactory integration application from within the PagerDuty Service Directory and receive a URL to setup the PagerDuty webhook within JFrog Artifactory
* JFrog Artifactory will utilize the PagerDuty webhook trigger to send notifications to PagerDuty.  Users will be able to configure their webhook events for repositories, builds and release bundles within JFrog Artifactory and associate these events with the PagerDuty webhook trigger. 
* Once the webhook is configured within JFrog Artifactory, an incident will be sent to the PagerDuty service whenever an event occurs.
* Artifactory performs a recheck of all selected repositories, builds and release bundles and any events generated will be sent as an event to a service in PagerDuty. Events from Artifactory will trigger a new incident on the corresponding PagerDuty service, or group as alerts into an existing incident.

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
## In PagerDuty
Add JFrog Artifactory integration to an existing PagerDuty service or create a new service by following the steps below

### Integrating With a PagerDuty Service
1. From the **Configuration** menu, select **Services**.
2. There are two ways to add an integration to a service:
   * **If you are adding your integration to an existing service**: Click the **name** of the service you want to add the integration to. Then, select the **Integrations** tab and click the **New Integration** button.
   * **If you are creating a new service for your integration**: Please read our documentation in section [Configuring Services and Integrations](https://support.pagerduty.com/docs/services-and-integrations#section-configuring-services-and-integrations) and follow the steps outlined in the [Create a New Service](https://support.pagerduty.com/docs/services-and-integrations#section-create-a-new-service) section, selecting ***JFrog Artifactory Notifications*** as the **Integration Type** in step 4. Continue with the In  ***JFrog Artifactory Setup***  section (below) once you have finished these steps.
3. Enter an **Integration Name** in the format `JFrog` and select  ***JFrog Artifactory Notifications***  from the Integration Type menu.
4. Click the **Add Integration** button to save your new integration. You will be redirected to the Integrations tab for your service.
5. An **Integration Key** will be generated on this screen. Click on the integration name next to the integration key and copy the **integration URL** in a safe place as it will be used when you configure the integration with  ***JFrog Artifactory Notifications***  in the next section.
![integrations](https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/images/step1.png)

## In JFrog Artifactory
1. Navigate to **General** in JFrog Platform's **Administration** section and click on it

![step1](https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/images/step1.png)

2. Click on **Webhooks**
3. Create a **New webhook** for pagerduty

![new webhook](https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/images/new_webhook.png)

4. Enter **Webhook Name**
5. Enter **URL** (URL is the Events API endpoint from pagerduty) _ex: https://events.pagerduty.com/integration/<integration_id>/enqueue_
6. Select event from the list of Events

![event_selection](https://raw.githubusercontent.com/jfrog/partner-integrations/main/PagerDuty/Artifactory/images/event_selection.png)

7. Select specific Repositories or Builds or Release Bundles to get notified about
8. Click **Create**

# How to Uninstall
1. Find the integration to delete by navigating to **Services** and selecting **Service Directory**
2. Select the service with your integration and navigate to **Integrations** tab
3. Navigate to the integration by clicking on the integration name (**JFrog Artifactory Notifications**)
4. Click on the **Delete Integration** button on the top right corner
