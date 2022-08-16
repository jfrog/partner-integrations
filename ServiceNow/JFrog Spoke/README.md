# Welcome to the JFrog Spoke Documentation

The JFrog Xray Spoke allows your organization to build automated workflows that interact directly with the JFrog Platform.

Focused on Xray-related actions: from generation violations reports to creating ignore rules, re-scanning builds and artifacts, assigning custom item properties, and assigning permissions to users and groups, the JFrog Xray Spoke provides out-of-the-box actions that your company can mix-and-match with other ServiceNow workflows to automate your overall IT-operations.

[To see a video of this example, click here.](https://youtu.be/pPcUVXn1Ds0)

# Support

If you need help with this integration, please contact `partner-support@jfrog.com`

# How It Works

ServiceNow Spokes are applications with predefined actions that customers can use to build workflows on the NOW Platform. The JFrog Spoke has actions that can be combined with other native Spokes from ServiceNow and other third-party tools to create enterprise-grade workflows for vulnerability and change management for all your ITSM needs.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/im-1.png?raw=true" width="800">

## Major Features Include:

* Track Xray Vulnerabilities
* Track License Violations
* Create Ignore Rules
* Create Users
* Update Permissions
* Generate and Send Xray Violations Reports
* Trigger Scans of new Artifacts and Builds
* Manage Custom Item Properties

# Getting Started

## Prerequisite: Required ServiceNow Dependencies

Before the spoke is installed, you will need to install the following plugins for the workflow to completely work.
**The required plugins are:**
* ServiceNow IntegrationHub Starter Pack Installer
* REST API Builder Backend

## JFrog Platform Steps:

Before you begin building a workflow using the JFrog Xray Spoke in ServiceNow, you will need to do some initial setup. To start, you can [watch this installation video here](https://youtu.be/yVUG4MqmGDg) and then follow along with the documentation below.

**Setup JFrog Xray Spoke with Connection Alias**

Most of our actions use the JFrog REST API. In order to use it, users must add their credential to the appropriate connection alias: “JFrog Platform Admin” or ID “x_jfro_xray_spoke.JFrog_Platform_Admin”.

The authentication method uses a bearer token generated from the JFrog platform. To generate a token, go to your profile on the JFrog Platform in the top right hand corner and select “edit profile.” You should see this screen:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m1.png?raw=true" width="800">

Click that “Generate an Identity Token” button after unlocking this page by entering your password. We will use this identity token to create an API credential and connect it to the aforementioned connection alias. (Do not use the Generate API Key option)

## Next: Go to ServiceNow

To connect these credentials in ServiceNow navigate to the Credentials Tab.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m2.png?raw=true" width="800">

Create a new API Key Credentials

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m3.png?raw=true" width="800">

Enter Identity Token from JFrog Platform as the API Key here:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m4.png?raw=true" width="800">

Next you go to Connections and Credentials Aliases, and enter the JFrog Platform Admin Connection and Credential Alias.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m5.png?raw=true" width="800">

In the connections tab near the bottom area, hit the New button

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m6.png?raw=true" width="800">

Then you provide the name of the alias and in the Credential input (hit the search button) and it will allow you to select the newly created credential

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m7.png?raw=true" width="800">

Enter your JPD base url in the Connection URL box:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m8.png?raw=true" width="800">

To save the credential alias, please hit Submit at the bottom.

Go to Flow Designer and connect Xray.

Next, you need to setup your webhook from ServiceNow.

Go to your ServiceNow instance. For example: https://jfrogshare.service-now.com/

Login using your credentials.

Next, Go to Flow Designer:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m9.png?raw=true" width="800">
Select New to create a new flow and name it.

Once you are in your flow, select a Trigger

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m10.png?raw=true" width="800">

Under Trigger, select REST API for Flows in ServiceNow

Most Xray flows will be triggered using a REST API post step. Creating this step will automatically generate a webhook that needs to be added to Xray as mentioned above in the section. Example:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m11.png?raw=true" width="800">

The Path can be used when you create the webhook in the JFrog Platform.

In the JFrog Platform, Xray webhooks can be found under Administration > Xray > Settings

Click on New Webhook in the top right hand corner

The ServiceNow + Path is the full webhook URL that needs to be added to Xray.

You add the URL from the Spoke Path to the URL area in the webhook:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/s1.png?raw=true" width="800">

Then you want to make sure to connect your webhook to the policy in JFrog Xray by going to Administration > Xray > Watches & Policies and click on the policy you want to add this webhook to.

Go to the rules for the policy and select trigger webhook and choose the webhook with the name you provided.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/s22.png?raw=true" width="800">

Once that is done, Xray should be setup to send violation data to ServiceNow.

This will allow ServiceNow to read the incoming application json.

## Watch the Spoke Video

To see an example of building a workflow just [watch this video here](https://youtu.be/q7MOf_r0s9Y) and then you can follow the steps below.

## Example 1: Build Your First Xray Flow

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/img3.png?raw=true" width="400">

First, to start use the JFrog Xray Spoke, you must be an Admin of your JFrog Platform Instance so you can create the initial webhook needed for Xray.

If you do not already have an Xray Watch and Policy created, this will be required to setup violations. [Learn how to setup watches and policies in Xray](https://www.youtube.com/watch?v=88hwwMJsS58).

Next, in your organization's ServiceNow instance, go to **Integration Hub** > **Flow Designer** > **New** > **Flow**

Create a New Flow.

Next, for the Trigger > select **REST API - Asynchronous**

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/i1.png?raw=true" width="800">

The ServiceNow + Path is the full webhook URL that needs to be added to Xray.

Once that is done, Xray should be setup to send violation data to ServiceNow. For ServiceNow to be able to read these violation records we need to define the complex object of the response body as follows:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/i2.png?raw=true" width="800">

This will allow ServiceNow to read the incoming application json.

This will then bring in the Xray violations payload data into ServiceNow so we can parse it and use the information to build out our flows.

Next, select **Actions**

Then, search for **JFrog Create Violation Record**. This will load up the Xray violation data (via webhook in the JFrog Platform) and make it available to all the Spoke actions.

Next, you can use that action to create a flow such as this example, **“For Each”** Violation Record where the severity is high, take an action:

**Action**: JFrog Xray Generate Violations Report

Next, you can use that action to create a flow such as this example, **“For Each”** Violation Record where the severity is medium, take an action:

**Action** JFrog Artifactory Set Item Properties

Next, you can use that action to create a flow such as this example, **“For Each”** Violation Record where the severity is low, take an action:

**Action** JFrog Xray Create Ignore Rules

Now you can save your flow! You will need to activate the flow before the Xray violations will work in the flow.

# List of Supported Actions

Name |
------------ |
Create Violation Record
Set Item Properties
Set Update Properties
Set Delete Properties
Create User
Update User
Delete User
Get Users
Get User Details
Scan Artifact
Artifact Summary
Create Ignore Rules
Delete Ignore Rules
Generate Violations Report
Export Violations Report
Create or Replace Permission Target
Delete Permission Target
Get Permission Target Details
Get Permission Targets
Trigger Scan Build
Scan Build Results

# On-premise JFrog Installation

To use the spoke with an on-premise JFrog installation, the network port `8082` will need to be exposed to the external network. See [JFrog System Requirements](https://www.jfrog.com/confluence/display/JFROG/System+Requirements#SystemRequirements-RequirementsMatrix) for more information.

Then in the Connection URL for the Credential Alias, include the network port in the JPD base url, e.g. `https://example.com:8082`

# Requirements:

* Your organization should have an instance of the JFrog Platform version 7.27 or above. You can signup for a free instance at: https://jfrog.com/start-free/

* You must be a user with Admin permissions to create the initial webhooks for Artifactory and Xray.

* Your organization must already have setup policies and watches prior to getting Xray notifications in Slack. [Learn how to setup watches and policies in Xray](https://www.youtube.com/watch?v=88hwwMJsS58).

* Your organization must already be have an active subscription to ServiceNow with access to IntegrationHub Spokes
