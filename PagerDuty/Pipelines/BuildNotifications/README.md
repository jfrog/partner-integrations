## PagerDuty + JFrog Pipelines Integration Benefits
* Monitor your software release pipeline in real-time.  Receive up-to-date notifications within PagerDuty about the progress of your builds throughout the various stages in your continuous integration and continuous deployment pipeline.   
* Respond quickly to build failures.  Receive actionable, granular information about a particular pipeline step, including its status, name, type, and other details to respond quickly to builds that are failing or taking too long to complete

## How does it Work?
* JFrog Pipeline will send notifications by calling PagerDuty’s Events API.
* PagerDuty users will be able to register the JFrog Pipeline integration application within their developer platform and set up the event integration.  Once the setup is complete, JFrog Pipelines will use the Events API to push notifications containing build status information.

## Requirements
* PagerDuty integration requires an admin base role for account authorization. If you do not have this role, please reach out to an Admin or Account Owner within your organization to configure the integration.
* JFrog Platform with Pipelines installed with a PagerDuty Pipelines Integration configured.

## Support
If you need help with this integration, please contact `partner_support@jfrog.com`

## Integration Walkthrough
### In PagerDuty
Add JFrog Pipelines integration to an existing PagerDuty service or create a new service by following the steps below

#### Integrating With a PagerDuty Service
1. From the **Configuration** menu, select **Services**.
2. There are two ways to add an integration to a service:
   * **If you are adding your integration to an existing service**: Click the **name** of the service you want to add the integration to. Then, select the **Integrations** tab and click the **New Integration** button.
   * **If you are creating a new service for your integration**: Please read our documentation in section [Configuring Services and Integrations](https://support.pagerduty.com/docs/services-and-integrations#section-configuring-services-and-integrations) and follow the steps outlined in the [Create a New Service](https://support.pagerduty.com/docs/services-and-integrations#section-create-a-new-service) section, selecting ***JFrog Pipelines Changes*** as the **Integration Type** in step 4. Continue with the In  ***JFrog Pipelines Setup***  section (below) once you have finished these steps.
3. Enter an **Integration Name** in the format `JFrog` and select  ***JFrog Pipelines + PagerDuty Notifications***  from the Integration Type menu.
4. Click the **Add Integration** button to save your new integration. You will be redirected to the Integrations tab for your service.
5. An **Integration Key** will be on this screen. This is the key that is called `service integration/routing key` when configuring the UI in pipelines. 
This confusion is inherited from the inconsistencies in the [PagerDuty api docs](https://developer.pagerduty.com/docs/ZG9jOjExMDI5NTgy-send-a-change-event) 
![IntegrationKey](images/IntegrationKey.png)
![IntegrationKey](images/IntegrationKey2.png)



### In JFrog Pipelines

#### Creating a Pipelines Integration for PagerDuty
1. Open your JFrog Platform with Pipelines and click on the `Adminstration gears` in the left menu. Click on the `Pipelines` menu item and then click on the sub-tab `Integrations`.
2. Click the `+ Add an Integration` link on the top right.
3. Select `PagerDuty Events` for the integration type
![AddNewIntegration](images/AddNewIntegration.png)
4. Enter a name for this integration which will be used in the pipeline yaml. ex: `MyPagerDutyIntegration`
5. Enter the Pagerduty `Integration Key` created in prior steps.
6. Click Create

#### Creating a Pipelines Build YAML
1. Create a new git repository or use an existing repository to save a new file `pipeline.yaml`
2. Create a new Pipelines integration for your SCM provider. [Visit the official documentation.](https://www.jfrog.com/confluence/display/JFROG/Managing+Pipelines+Integrations)
3. Create a new pipeline YAML definition that uses the new PagerDuty Events integration above to send build notifications to PagerDuty.
4. Replace `<MyPagerDutyIntegration>` with the name of the Pipelines integration for Pagerduty
````text
resources:
  - name: pagerDutyGitRepo
    type: GitRepo
    configuration:
      gitProvider: MyGithubIntegration
      path: jfrog/jfrog-pipelines-example
pipelines:
  - name: pipeline_pagerduty_example
    steps:
      - name: pagerduty_example_step
        type: Bash
        configuration:
         integrations:
           - name: <MyPagerDutyIntegration>
          inputResources:
            - name: pagerDutyGitRepo
        execution:
          onExecute:
           - send_notification <MyPagerDutyIntegration> --text "<Your message>"
````
Commit the pipeline yaml to your SCM provider and then [follow the official steps on adding a pipeline](https://www.jfrog.com/confluence/display/JFROG/Managing+Pipeline+Sources#ManagingPipelineSources-AddingaPipelineSource). 

#### Running a Pipelines Build
1. Click the JFrog Pipeline tab `My Pipelines`.
2. Select the Pipeline you have synced .
3. Click the Pipeline tile and click the `Trigger this step` button
![PipelinesTrigger](images/PipelineTrigger.png).

Or follow the official steps on all the ways of how to [trigger a JFrog pipelines.](https://www.jfrog.com/confluence/display/JFROG/Triggering+Pipelines+and+Steps)

## How to Uninstall
1. Find the integration to delete by navigating to **Services** and selecting **Service Directory**.
2. Select the service with your integration and navigate to **Integrations** tab.
3. Navigate to the integration by clicking on the integration name (**JFrog Pipelines + PagerDuty Notifications**).
4. Click the **Delete Integration** button on the top right corner.
5. Remove any PagerDuty Pipelines in the `Pipelines Sources` tab of JFrog Pipelines by clicking the `Delete` button shown below.
![PipelinesDelete](images/PipelinesDelete.png).
