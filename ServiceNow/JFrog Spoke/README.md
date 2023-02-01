# Welcome to the JFrog Spoke Documentation

The JFrog Spoke allows your organization to build automated workflows that interact directly with the JFrog Platform.

Focused on Xray, Artifact, Repository, and JFrog Admin actions. From generating violations reports to creating ignore rules, re-scanning builds and artifacts, assigning custom item properties, and assigning permissions to users and groups, the JFrog Spoke provides out-of-the-box actions that your company can mix-and-match with other ServiceNow workflows to automate your overall IT-operations.

[To see a video of this example, click here.](https://youtu.be/pPcUVXn1Ds0)

# Support

If you need help with this integration, please contact `support@jfrog.com`

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
* Move, Copy and Delete Artifacts
* Create, Update, and Delete Watches
* Create, Update and Delete Policies
* Manage Group Permissions
* Ingest Xray data into ServiceNow Tables

# Getting Started

## Prerequisite: Required ServiceNow Dependencies

Before the spoke is installed, you will need to install the following plugins to your ServiceNow instance for the spoke to completely work.
**The required plugins are:**
* ServiceNow IntegrationHub Starter Pack Installer
* REST API Builder Backend

## JFrog Platform Steps:

Before you begin building a workflow using the JFrog Spoke in ServiceNow, you will need to do some initial setup. To start, you can [watch this installation video here](https://youtu.be/yVUG4MqmGDg) and then follow along with the documentation below.

**Setup JFrog Spoke with Connection Alias**

Most of our actions use the JFrog REST API. In order to use it, users must add their credential to the appropriate connection alias: “JFrog Platform Admin” or ID “x_jfro_xray_spoke.JFrog_Platform_Admin”.

The authentication method uses a bearer token generated from the JFrog platform. To generate a token, go to your profile on the JFrog Platform in the top right hand corner and select “edit profile.” You should see this screen:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m1.png?raw=true" width="800">

Click that “Generate an Identity Token” button after unlocking this page by entering your password. We will use this identity token to create an API credential and connect it to the aforementioned connection alias. (Do not use the Generate API Key option)

## Next: Go to ServiceNow

To connect these credentials in ServiceNow navigate to the Credentials Tab.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m2.png?raw=true" width="800">

Create a new API Key Credentials

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m3.png?raw=true" width="800">

Enter the Identity Token from the JFrog Platform as the API Key here:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m4.png?raw=true" width="800">

Next you go to Connections and Credentials Aliases, and enter the JFrog Platform Admin Connection and Credential Alias.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m5.png?raw=true" width="800">

In the connections tab near the bottom area, hit the New button

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m6.png?raw=true" width="800">

The connection alias should come prefilled; if it isn't make sure to add an alias, so it matches what is seen below.
Next, in the Credential input (hit the search button), and it will allow you to select the newly created credential

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m7.png?raw=true" width="800">

Enter your JPD base url in the Connection URL box:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/m8.png?raw=true" width="800">

To save the credential alias, please hit Submit at the bottom.

### Send Xray Data to ServiceNow

- Our next step is to set up Xray to successfully send data to ServiceNow.
We are going to create a webhook in Xray that points to an endpoint that is a part of the ServiceNow Spoke.
The url/api endpoint you will need to use is:

- **https://{service-now-url}/api/x_jfro_xray_spoke/xray**

- Replace `service-now-url` with your servicenow base url. It should look something like `jfrog.service-now.com`.

- In the JFrog Platform, Xray webhooks can be found under Administration > Xray > Settings

- Click on New Webhook in the top right-hand corner

- You add the URL provided above to the URL area in the webhook:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/s1.png?raw=true" width="800">

This won't work until we authenticate with ServiceNow.
In order to do this you will need to create a service account that has "Web service access only" enabled in the ServiceNow platform:

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/ServiceAccount.png?raw=true" width="800">

- Enter these service account's credentials in the basic auth section on the webhook creation page in Xray.

- No data will be sent from Xray unless this webhook is attached to a policy.
Therefore, make sure you connect your webhook to a policy in JFrog Xray by going to Administration > Xray > Watches & Policies and clicking on the policy you want to add this webhook to.

- Go to the rules for the policy and select trigger webhook and choose the webhook with the name you provided.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/JFrog%20Spoke/images/s22.png?raw=true" width="800">

- Finally make sure that policy is associated with a watch on your relevant repositories and builds.
  For more information on how watches and policies work in Xray, please see: [Configuring Xray Watches.](https://www.jfrog.com/confluence/display/JFROG/Configuring+Xray+Watches)

Once that is done, Xray should be setup to send violation data to ServiceNow.

This will allow ServiceNow to read the incoming application json. It will store this data in four tables in ServiceNow that can be accessed from the spoke called:
* Violations
* Issues
* Impacted Artifacts
* Infected Files

A description of each table is located in [ServiceNow Tables for Xray](./TABLES.md)

## Example 1: Build Your First Xray Flow

<img src="./images/HelloWorldFlow.png?raw=true" width="800">

- First, to start using the JFrog Spoke, you must be an Admin of your JFrog Platform instance in order to create the initial webhook needed for Xray.

- If you do not already have a Xray Watch and Policy created, this will be required to set up violations. [Learn how to setup watches and policies in Xray](https://www.youtube.com/watch?v=88hwwMJsS58).

- Next, in your organization's ServiceNow instance, go to **Integration Hub** > **Flow Designer** > **New** > **Flow**

- Create a New Flow. Make sure the flow is run as a System User. (Some actions require System User privileges)

- Next, for the Trigger > select **Created**

- Select the table **Impacted Artifacts**:

<img src="./images/Trigger.png?raw=true" width="800">

- Once that is done, the flow should be able to interact with data that is being sent to ServiceNow through the URL mentioned above.

- Now we can do some cool things with this data using the actions in the JFrog Spoke.
In this flow we will do some actions based on the severity of the impacted artifact.

- To do this we will select **Flow Logic** to create an if statement that will evaluate true depending on the severity.
Then we will take actions based on the severity.

**Flow Logic**: If severity is High

**Action**: JFrog Artifactory Move Item

**Flow Logic**: Else If severity is medium

**Action**: JFrog Artifactory Set Item Properties

**Flow Logic**: Else If severity is low

**Action**: JFrog Xray Create Ignore Rules

All these actions require configuration that is up to you! To see what these actions do, please see the list below.

Now you can save your flow! You will need to activate the flow before the flow does anything. The button to do this is in the top right corner.
The flow will start working as expected after a violation is sent to ServiceNow after the flow is activated.

# List of Supported Actions
These actions all implement APIs from Artifactory, Xray, and the JFrog Platform. Please see these documents to gain a deeper understanding of the functionality available:
1. [Artifactory](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API)
2. [Xray](https://www.jfrog.com/confluence/display/JFROG/Xray+REST+API)
3. [JFrog Platform](https://www.jfrog.com/confluence/display/JFROG/JFrog+Platform+REST+API)

|Name | Description |
|------------ | ------------ |
|JFrog Create User | Creates an user |
|JFrog Update User | Updates an user |
|JFrog Delete User | Deletes an user |
|JFrog Get User | Retrieve a single user |
|JFrog Get Users | Retrieve a list of users |
|JFrog Create or Replace Permission Target | Creates a new permission target or replaces an existing one. |
|JFrog Delete Permission Target | Deletes a permission target |
|JFrog Get Permission Targets | Retrieves a list of permission targets |
|JFrog Get Permission Target Details | Retrieve the details of a specific permission target |
|JFrog Artifactory Update Group | Updates an existing group in artifactory |
|JFrog Artifactory Get Groups | Retrieves a list of groups |
|JFrog Artifactory Get Group Details | Get details of a single group |
|JFrog Artifactory Create or Replace Group | Creates a new group or replaces an existing one |
|JFrog Artifactory Set Item Properties | Set a new item property (or properties) |
|JFrog Artifactory Update Item Properties | Update a single item property (by key) |
|JFrog Artifactory Delete Item Properties | Delete an item property |
|JFrog Artifactory Move Item | Move an artifact or other item/file in artifactory |
|JFrog Artifactory Copy Item | Copy an artifact or other item/file in artifactory |
|JFrog Artifactory Delete Item | Delete an artifact or other item/file in artifactory |
|JFrog Artifactory Get Repos | Get a list of repositories |
|JFrog Artifactory Delete Repository | Delete a repository |
|JFrog Artifactory Create Repository | Create a repository |
|JFrog Xray Generate Violations Report | Generate a violations report |
|JFrog Xray Export Violations Report | Export a violations report. Stores it in the attachments table in ServiceNow |
|JFrog Xray Generate Violations Report (Advanced) | Generate a violations report. Only accepts a json payload input. Please check API spec to see details. |
|JFrog Xray Scan Artifact | Triggers an Xray scan of an artifact |
|Jfrog Xray Trigger Scan Build | Triggers an Xray scan of a build |
|Jfrog Xray Scan Build Results | Retrieves the results of a previous build scan |
|JFrog Xray Artifact Summary | Retrieves a security summary of an artifact |
|JFrog Xray Create Ignore Rules | Creates ignore rules |
|JFrog Xray Delete Ignore Rule | Deletes ignore rules |
|JFrog Xray Delete Watch | Deletes a Xray watch |
|JFrog Xray Delete Policy | Deletes a Xray policy |
|JFrog Xray Update Policy | Updates a Xray policy |
|JFrog Xray Create Security (min severity) Policy | Creates a policy based on severity |
|JFrog Xray Create Security (CVSS range) Policy | Creates a policy based on CVSS range |
|JFrog Xray Create License (allowed) Policy | Creates a policy based around which licenses are allowed |
|JFrog Xray Create License (banned) Policy | Creates a policy based around which licenses are banned |
|JFrog Xray Create Operational Risk (min risk) Policy | Creates a policy based around a minimum operational risk |
|JFrog Xray Create Operational Risk (custom criteria) Policy | Creates a policy based around custom parameters for operational risk |
|JFrog Xray Create Watch | Creates a Xray Watch |
|JFrog Xray Update Watch | Updates a Xray Watch |
|JFrog Xray Get Watch | Retrieves a Xray Watch |
|JFrog Xray Get Policy | Retrieves any Xray Policy |

# On-premise JFrog Installation

To use the spoke with an on-premise JFrog installation, the network port `8082` will need to be exposed to the external network. See [JFrog System Requirements](https://www.jfrog.com/confluence/display/JFROG/System+Requirements#SystemRequirements-RequirementsMatrix) for more information.

Then in the Connection URL for the Credential Alias, include the network port in the JPD base url, e.g. `https://example.com:8082`

# Requirements:

* Your organization should have an instance of the JFrog Platform version 7.27 or above. You can signup for a free instance at: https://jfrog.com/start-free/

* You must be a user with Admin permissions to create the initial webhooks for Artifactory and Xray.

* Your organization must already have setup policies and watches prior to getting Xray notifications in Slack. [Learn how to setup watches and policies in Xray](https://www.youtube.com/watch?v=88hwwMJsS58).

* Your organization must already be have an active subscription to ServiceNow with access to IntegrationHub Spokes
