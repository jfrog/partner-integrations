# Slack

This integration is between JFrog Artifactory and Xray and Slack. We know that software development happens in a myriad of tools and collaboration environments. Today there are key events throughout the JFrog Platform that can be difficult for a user to interact with if they aren't logged in. When it comes to people across the organization knowing what’s going on, there aren’t great solutions. This will give each user situational awareness about occurrences in the JFrog Platform.  Additionally, where appropriate - they will have easy links and action buttons to go follow-up on the event.

# How it Works

* This integration allows you the ability to see Artifact, Artifact Properties, Docker, Release Bundle, and Build events through notifications and actionable cards inside the Microsoft Teams browser. 
* Additionally, you can get notification based on policies setup in JFrog Xray.

# Requirements

* You must be a user with Admin permissions to generate a JPD key.
* You must have the ability to setup policies and watches prior to getting Xray notifications.
* You should have a cloud instance of Artifactory. You can signup for a free cloud instance at: https://jfrog.com/start-free/

# Support

If you need help with this integration, please contact `partner_support@jfrog.com`

# Getting Started

The first thing to do is to download the Slack Application. 

Then, configure your JFrog instance.

Next, login to your JFrog account. 

Once logged in with admin privledges, you can start creating notifications. 

Hit **create notifications** to bring up the list of options. 

Select which type notification you would like to create from the drop-down menu. 

On the next screen, name the notification, select which events you would like to include in the notification, and which repos should be included. You can also setup include/exclude pattens and select a channel to send the notifications to.

Once you have setup notifications, you should start seeing the notification cards in the channel within 20 minutes.

# List of Commands

## Artifactory

**/jfrog login**
Connects to your JFrog Instance and asks for JFrog Username and Password.

**/jfrog rt notify list**
Provides a list of current notifications subscribed to by the personal or channel

**/jfrog rt notify stop <#>**
Ends subscription to the specified notification.  

**/jfrog rt notify artifact <reponame> <pattern>**
Reponame should have any local, remote or virtual repository as an option. This should send notifications on the ‘artifact deployed’ event for the specified repository and path.  

**/jfrog rt notify docker <reponame> <pattern>**
This should send notifications on the ‘docker push’ and ‘docker promote’ events for the specified repository and path.

**/jfrog rt notify <packagetype> <reponame> <package coordinate>**
This should send notifications on the ‘artifact deployed’ event for the specified repository with a path calculated based on the package coordinate. 

**/jfrog rt notify build <buildname>**
Active card should provide build name, number, build agent, build principal, timestamp, a link back, and promotion status, as well as triggering event (uploaded to, promoted to)

**/jfrog rt artifact list <wildcard or regex>**
Yields a list of artifacts matching the repo filter

**/jfrog rt repo list <wildcard or regex>**
Yields a list of all repos matching the filter
  
## Xray 

**/jfrog xr notify list**
Provides a list of current xray notifications subscribed to by the personal or channel bot context.

**/jfrog xr ignore watch watch_name (date) <#>**
This ignores the specific watch. 
Users must specify length of time. 

**/jfrog xr ignore artifact artifact_name (date) <#>**
This ignores the specific watch. 
Users must specify length of time. 
Based on specific watch for any any component, any artifact.

**/jfrog xr ignore vuln vulnerability_id (date) <#>**
Allows user to specify an ignore rule based on a specific vulnerability and specify the date it stops ignoring.
Users must specify length of time. 
Based on any watch, any component, any artifact.

**/jfrog xr watch <watchname> **
Directs a watch to a given channel. 
Active card should provide a link back. 

**/jfrog xr watch list <regex wildcard>**
Provides a list of current watches (that user can has read access to) with a micro action to subscribe the bot to the notification.

**/jfrog xr policy list <regex or wildcard>**
Provides a list of current policies with a micro action to open JFrog Platform.

