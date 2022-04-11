
# Microsoft Teams 

Today, there are key events throughout the JFrog Platform that can be difficult for a user to interact with if they aren't logged into the platform. When it comes to people across different departments in an organization knowing what’s going on, it’s imperative that events that affect artifacts and repositories be made more visible. 

This integration is between the JFrog Platform and Microsoft Teams. The main features of this integration include notifications coming from JFrog Xray, Artifactory, and Distribution. Build Information is also available.

# How it Works

* This integration allows you the ability to see Artifact, Artifact Properties, Docker Tag, and Build events through notifications and cards where you can take actions that interact back with the JFrog Platform.
* Additionally, you can get vulnerability and license compliance notifications based on policies setup in JFrog Xray. 
* If you are an Enterprise user, you can also get notifications around release bundle and distribution events. 

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

### Xray

Type | Description
------------ | -------------
Security Violations by CVE | *This sends individual notification cards for each CVE or issue*
Security Violations by Component (Summary view) | *This provides a summary of all CVEs and severities by component*
License Compliance | *This sends individual notification cards for each license compliance issue*
