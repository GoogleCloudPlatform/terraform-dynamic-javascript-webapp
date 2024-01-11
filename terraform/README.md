# ⚙️  Provision with Terraform

## Provision a CI/CD pipeline <a name="provision-a-cicd-pipeline"></a>

This includes an example CI/CD pipeline that uses [Cloud Build] and [Cloud Deploy] to continuously intergrate 
changes to the application code into new builds of your application and automatically deploy those builds to to Cloud 
Run.  The following steps are required to set this pipeline up:

1. Deploy application infrastructure to a Google Cloud project
2. Fork this repository 
3. Connect the fork to your Google Cloud project
4. Run the Terraform code in `terraform/environments/main`

### 1. Deploy the application infrastructure

Using a Google Cloud Project with billing enabled, deploy the application infrastructure via the Terraform made 
available in the 
[terraform-dynamic-javascript-webapp repo](https://github.com/GoogleCloudPlatform/terraform-dynamic-javascript-webapp/tree/main/infra) 

If you'd like to provision manually, see the [manual setup section](#manually-provision-run-firestore) before continuing.

### 2. Connect to your Google Cloud Project

Cloud Build triggers can be created to respond to GitHub repository actions, such as pull requests, and merges to the
main branch. To be able to create such triggers, the Google Cloud project must have the repository added to it in the 
region of the trigger.

Connect your GitHub repo to your Google Cloud project via the 
[Cloud Build triggers page](https://console.cloud.google.com/cloud-build/triggers/connect), making sure to specify the 
'global' region.


### 3. Run the Terraform code

The `terraform/environments/main` directory is a root Terraform module that creates the CICD pipeline resources in the
specified project. To apply the Terraform to your project, do the following from the `main` directory:

1. Supply the values for the variables

    Rename `terraform.tfvars.example` to terraform.tfvars. This file is where you can provide values for the required
    variables.  

2. Run `terraform init`

   To initiate the root module, run [`terraform init`](https://developer.hashicorp.com/terraform/cli/commands/init). 
   This initializes the working directory containing Terraform configuration files.

3. Run `terraform plan`

   Running [`terraform plan`](https://developer.hashicorp.com/terraform/cli/commands/plan) will output an execution 
   plan that will servce as a preview of any changes that will be made to your Google Cloud project. 

4. Run `terraform apply`
   
   To apply the changes to your Google Cloud project, run 
   [`terraform apply`](https://developer.hashicorp.com/terraform/cli/commands/apply). The output will include the plan
   and a prompt requesting that you approve the plan.  Enter `yes` to deploy the Google Cloud resources to your 
   project.  

___

<!-- doc links -->
[Artifact Registry]:
https://cloud.google.com/artifact-registry

[Artifact Registry Console]:
https://console.cloud.google.com/artifacts

[Cloud Firestore]:
https://cloud.google.com/firestore

[Firestore Console]:
https://console.cloud.google.com/firestore

[Cloud Build]:
https://cloud.google.com/build

[Cloud Build Triggers]:
https://cloud.google.com/build/docs/triggers

[Cloud Deploy]:
https://cloud.google.com/deploy

[Cloud Run]:
https://cloud.google.com/run

[Secret Manager]:
https://console.cloud.google.com/security/secret-manager

[Terraform]:
(https://www.terraform.io)
