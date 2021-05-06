#!/usr/bin/env bash
export BUILD_NUMBER=2
export JIRA_ENVIRONMENT_ID=JFrogEnvironment_dev_${BUILD_NUMBER} 
export JIRA_ENVIRONMENT_NAME=JFrogEnvironment_dev_${BUILD_NUMBER} 
export JIRA_DEPLOYMENT_STATUS=successful
export JIRA_ENVIRONMENT_TYPE=development
echo "Deploying to Development Environment"
jfrog rt bag ${BUILD_NAME} ${BUILD_NUMBER} --config=$HOME/.jfrog/jira-cli.conf
jfrog rt bce ${BUILD_NAME} ${BUILD_NUMBER} 
jfrog rt bp ${BUILD_NAME} ${BUILD_NUMBER}
echo "Deploying to Test Environment"
export BUILD_NUMBER=3
export JIRA_ENVIRONMENT_ID=JFrogEnvironment_test_${BUILD_NUMBER}                            
export JIRA_ENVIRONMENT_NAME=JFrogEnvironment_test_${BUILD_NUMBER}                            
export JIRA_DEPLOYMENT_STATUS=successful
export JIRA_ENVIRONMENT_TYPE=testing
jfrog rt bag ${BUILD_NAME} ${BUILD_NUMBER} --config=$HOME/.jfrog/jira-cli.conf
jfrog rt bce ${BUILD_NAME} ${BUILD_NUMBER}
jfrog rt bp ${BUILD_NAME} ${BUILD_NUMBER}
echo "Deploying to Staging Environment"
export BUILD_NUMBER=4
export JIRA_ENVIRONMENT_ID=JFrogEnvironment_stage_${BUILD_NUMBER}
export JIRA_ENVIRONMENT_NAME=JFrogEnvironment_stage_${BUILD_NUMBER}
export JIRA_DEPLOYMENT_STATUS=successful
export JIRA_ENVIRONMENT_TYPE=staging
jfrog rt bag ${BUILD_NAME} ${BUILD_NUMBER} --config=$HOME/.jfrog/jira-cli.conf
jfrog rt bce ${BUILD_NAME} ${BUILD_NUMBER}
jfrog rt bp ${BUILD_NAME} ${BUILD_NUMBER}
echo "Deploying to Production Environment"
export BUILD_NUMBER=5
export JIRA_ENVIRONMENT_ID=JFrogEnvironment_prod_${BUILD_NUMBER}
export JIRA_ENVIRONMENT_NAME=JFrogEnvironment_prod_${BUILD_NUMBER}
export JIRA_DEPLOYMENT_STATUS=successful
export JIRA_ENVIRONMENT_TYPE=production
jfrog rt bag ${BUILD_NAME} ${BUILD_NUMBER} --config=$HOME/.jfrog/jira-cli.conf
jfrog rt bce ${BUILD_NAME} ${BUILD_NUMBER}
jfrog rt bp ${BUILD_NAME} ${BUILD_NUMBER}
