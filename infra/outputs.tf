/**
 * Copyright 2023 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

output "frontend_url" {
  value = "http://${google_compute_global_address.default.address}/"
}

output "run_service_account" {
  description = "The email of the Cloud Run service account for the service."
  value       = google_service_account.cloud_run.email
}

output "run_service_name" {
  description = "The name of the deployed Cloud Run service."
  value       = google_cloud_run_v2_service.default.name
}
