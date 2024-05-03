/**
 * Copyright 2024 Google LLC
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

# The "existing" Firestore database that the JSS shouldn't try to recreate.
# We need to make sure the JSS doesn't produce an error similar to
# "Database already exists. Please use another database_id"
resource "google_firestore_database" "database" {
  project                     = module.project.project_id
  name                        = "(default)"
  location_id                 = "nam5"
  type                        = "FIRESTORE_NATIVE"
  concurrency_mode            = "PESSIMISTIC"
  app_engine_integration_mode = "DISABLED"
}

module "dynamic_web_app" {
  source     = "../.."
  project_id = var.project_id
  depends_on = [
    google_firestore_database.database
  ]
}
