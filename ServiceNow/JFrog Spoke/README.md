# Welcome to the JFrog Spoke Documentation

The JFrog Xray Spoke allows your organization to build automated workflows that interact directly with the JFrog Platform. 

Focused on Xray-related actions: from generation violations reports to creating ignore rules, re-scanning builds and artifacts, assigning custom item properties, and assigning permissions to users and groups, the JFrog Xray Spoke provides out-of-the-box actions that your company can mix-and-match with other ServiceNow workflows to automate your overall IT-operations. 

# How It Works

ServiceNow Spokes are applications with predefined actions that customers can use to build workflows on the NOW Platform. The JFrog Spoke has actions that can be combined with other native Spokes from ServiceNow and other third-party tools to create enterprise-grade workflows for vulnerability and change management for all your ITSM needs.

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/ServiceNow/Jfrog Spoke/images/ms-2.png" width="800">

## Major Features Include:

* Track Xray Vulnerabilities
* Track License Violations
* Create Ignore Rules
* Create Users
* Update Permissions
* Generate and Send Xray Violations Reports
* Trigger Scans of new Artifacts and Builds
* Manage Custom Item Properties

# Support

If you need help with this integration, please contact `partner-support@jfrog.com`

# Getting Started

## Example 1: Build Your First Xray Flow

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/ServiceNow/Jfrog Spoke/images/img3.png" width="400">

First, to start use the JFrog Xray Spoke, you must be an Admin of your JFrog Platform Instance so you can create the initial webhook needed for Xray. 

If you do not already have an Xray Watch and Policy created, this will be required to setup violations. [Learn how to setup watches and policies in Xray](https://www.youtube.com/watch?v=88hwwMJsS58).

Next, in your organization's ServiceNow instance, go to **Integration Hub** > **Flow Designer** > **New** > **Flow**

Create a New Flow. 

Next, for the Trigger > select **REST API - Asynchronous**

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/ServiceNow/Jfrog Spoke/images/i1.png" width="800">

The ServiceNow + Path is the full webhook URL that needs to be added to Xray.

Once that is done, Xray should be setup to send violation data to ServiceNow. For ServiceNow to be able to read these violation records we need to define the complex object of the response body as follows: 

<img src="https://raw.githubusercontent.com/jfrog/partner-integrations/ServiceNow/Jfrog Spoke/images/i2.png" width="800">

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

## Next: Watch the Video

[To watch a video of these examples and learn more, click here.](https://youtu.be/pPcUVXn1Ds0)

## List of Supported Actions

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
Create or Replace Permission Target | 
Delete Permission Target | 
Get Permission Target Details | 
Get Permission Targets
Trigger Scan Build
Scan Build Results

## Requirements:

* Your organization should have a cloud instance of the JFrog Platform version 7.27 or above. You can signup for a free cloud instance at: https://jfrog.com/start-free/

* You must be a user with Admin permissions to create the initial webhooks for Artifactory and Xray.

* Your organization must already have setup policies and watches prior to getting Xray notifications in Slack. [Learn how to setup watches and policies in Xray](https://www.youtube.com/watch?v=88hwwMJsS58).

* Your organization must already be have an active subscription to ServiceNow with access to IntegrationHub Spokes
