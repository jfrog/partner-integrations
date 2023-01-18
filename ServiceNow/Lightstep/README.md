# Welcome to the Documentation for Lightstep and Xray

The integration between JFrog Xray and Lightstep provides organizations a way to manage incident response for security and license compliance issues. With this integration organizations can combine JFrog Xray application security code scanning with Lightstep’s intelligent incident response and management capabilities, to identify compliance and security issues earlier in their DevOps pipeline and engage the necessary teams for timely response and remediation.

[To see a video of this example, click here](https://youtu.be/uFbo8EQ9e08)

# Support
If you need help with this integration, please contact support@jfrog.com

# How It Works
JFrog Xray violations are sent via a webhook into Lightstep where they surface as incidents and alerts. Once in Lightstep, they can be managed through Lightstep’s incident lifecycle management tools. 

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/Lightstep/images/image2.png?raw=true" width="800">

Each incident/alert from Xray includes metadata about the vulnerability or license issue. From Lightstep, a response administrator can assign these incidents to dedicated SRE teams, set up automated actions, prioritize and acknowledge incoming issues and compose notes, or even collaborate with other teams to accelerate the remediation.

<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/Lightstep/images/image1.png?raw=true" width="800">

  # How to Get Started
  
The JFrog Xray integration is available from within the Lightstep console:
Select Integrations from the Lightstep console’s navigation pane.
Click the JFrog Xray integrations card
Fill out the details in the form to enable the integration.
  
<img src="https://github.com/jfrog/partner-integrations/blob/main/ServiceNow/Lightstep/images/image3.png?raw=true" width="800">

You’ll then need to [create a webhook endpoint in Lightstep](https://lightstep.com/incident-response/docs/jfrog-xray-create-webhook) for JFrog Xray to send automated real-time messages and information to Lightstep Incident Response. When you click Generate Webook, Lightstep will create a webhook URL for the secure endpoint. 
You can then use this webhook URL to [configure the webhook in JFrog Xray](https://lightstep.com/incident-response/docs/jfrog-xray-configure-webhook). 
Once Lightstep and Xray are connected in this way, you can create [security and license policies](https://www.jfrog.com/confluence/display/JFROG/Creating+Xray+Policies+and+Rules) in JFrog Xray. You can specify rules for Xray to look for the specific CVEs, severity level, or other criteria that you care most about. In your policies, you can specify an automatic action to trigger the Lightstep webhook through the JFrog Platform event service when Xray discovers that policy has been violated, and send a violation event message to Lightstep.

**Need More Information?**
    
To see how to create Xray policies and watches, [view this support video](https://www.youtube.com/watch?v=88hwwMJsS58).
Everything is now ready for you to manage and remediate your Xray policy violations through Lightstep!
    
# Requirements
* Your organization should have a running instance of the JFrog Platform with Xray installed. You can signup for a free instance at: https://jfrog.com/start-free/
* You must be a user with Admin permissions to create the initial webhooks for Artifactory and Xray.
* Your organization must already have setup policies and watches prior to getting Xray notifications. 
* Your organization must already have an active subscription with Lightstep.
