provider "google" {
  project = var.project_id
  billing_project = var.project_id
}

provider "google-beta" {
  project = var.project_id
  billing_project = var.project_id
}

module "dynamic_web_app" {
  source     = "../.."
  project_id = var.project_id
}