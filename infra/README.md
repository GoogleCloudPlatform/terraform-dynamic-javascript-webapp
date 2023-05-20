# terraform-dynamic-javascript-webapp

## Description
### Tagline
This is an auto-generated module.

### Detailed
This module was generated from [terraform-google-module-template](https://github.com/terraform-google-modules/terraform-google-module-template/), which by default generates a module that simply creates a GCS bucket. As the module develops, this README should be updated.

The resources/services/activations/deletions that this module will create/trigger are:

- Create a GCS bucket with the provided name

### PreDeploy
To deploy this blueprint you must have an active billing account and billing permissions.


## Documentation

Upon successful provisioning, you should see an output similar to:

```bash
frontend_url = "http://11.222.333.444/"
neos_toc_url = "http://console.cloud.google.com?walkthrough_id=panels--sic--dynamic-javascript-web-app_toc"
run_service_name = "dev-journey"
```

## Usage

Basic usage of this module is as follows:

```hcl
module "dynamic_web_app" {
  source  = "terraform-google-modules/dynamic-javascript-webapp/google"
  version = "~> 0.1"
  project_id  = "<PROJECT ID>"
}
```

Functional examples are included in the
[examples](./examples/) directory.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| deployment\_name | Identifier for the deployment. Used in some resource names. | `string` | `"dev-journey"` | no |
| enable\_apis | Whether or not to enable underlying apis in this solution. | `bool` | `true` | no |
| initial\_run\_image | Initial image to deploy to Cloud Run service. | `string` | `"gcr.io/hsa-public/developer-journey/app"` | no |
| labels | A set of key/value label pairs to assign to the resources deployed by this solution. | `map(string)` | `{}` | no |
| project\_id | The project ID to deploy resources to. | `string` | n/a | yes |
| region | Google Cloud region | `string` | `"us-central1"` | no |

## Outputs

| Name | Description |
|------|-------------|
| frontend\_url | IP address to site. Load balancer expected to take ~5 minutes for it to warm up. |
| neos\_toc\_url | Tutorial URL |
| run\_service\_name | The name of the deployed Cloud Run service. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

These sections describe requirements for using this module.

### Software

The following dependencies must be available:

- [Terraform][terraform] v0.13
- [Terraform Provider for GCP][terraform-provider-gcp] plugin v3.0

### Service Account

A service account with the following roles must be used to provision
the resources of this module:

- Datastore Owner: `roles/datastore.owner`
- Cloud Run Invoker: `roles/run.invoker`
- Storage Object Viewer: `roles/storage.objectViewer`
- Secret Manager: `roles/secretmanager.secretAccessor`

The [Project Factory module][project-factory-module] and the
[IAM module][iam-module] may be used in combination to provision a
service account with the necessary roles applied.

### APIs

A project with the following APIs enabled must be used to host the
resources of this module:

- Google Cloud Asset API: `cloudasset.googleapis.com`
- Google Cloud Build API: `cloudbuild.googleapis.com`
- Google Cloud Resource Manager API: `cloudresourcemanager.googleapis.com`
- Google Cloud Compute API: `compute.googleapis.com`
- Google Cloud Firestore API: `firestore.googleapis.com`
- Google Cloud IAM API: `iam.googleapis.com`
- Google Cloud Run API: `run.googleapis.com`
- Google Cloud Secret Manager API: `secretmanager.googleapis.com`
- Google Cloud Service Usage API: `serviceusage.googleapis.com`
- Google Cloud Storage JSON API: `storage.googleapis.com`

The [Project Factory module][project-factory-module] can be used to
provision a project with the necessary APIs enabled.

## FAQ

**Why isn't the app popping up right away after provisioning has completed?**

Upon opening the given IP address (`frontend_url` output) from the load balancer,
you may have to wait around 5 minutes for your app to appear.

**Why am I seeing a security warning when I open the app up with the given Cloud Run url after attempting to "login"?**

If you are accessing the app through the Cloud Run given `https` (secure) url,
attempting to login will route you to the load balancer IP address (`frontend_url` output) which is currently `http` (not secured).
We recommend that when you access the app with the given IP address and not through the deployed url given through `Cloud Run`.

## Contributing

Refer to the [contribution guidelines](./CONTRIBUTING.md) for
information on contributing to this module.

[iam-module]: https://registry.terraform.io/modules/terraform-google-modules/iam/google
[project-factory-module]: https://registry.terraform.io/modules/terraform-google-modules/project-factory/google
[terraform-provider-gcp]: https://www.terraform.io/docs/providers/google/index.html
[terraform]: https://www.terraform.io/downloads.html

## Security Disclosures

Please see our [security disclosure process](./SECURITY.md).
