// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multiple_buckets

import (
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/utils"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSimpleExample(t *testing.T) {
	example := tft.NewTFBlueprintTest(t)

	example.DefineVerify(func(assert *assert.Assertions) {
		// example.DefaultVerify(assert) // disables default verify

		projectID := example.GetTFSetupStringOutput("project_id")
		server_service_name := terraform.OutputRequired(t, example.GetTFOptions(), "run_service_name")

		{
			// Check that expected Cloud Run service is deployed
			cloudRunServices := gcloud.Run(t, "run services list", gcloud.WithCommonArgs([]string{"--filter", "metadata.name=" + server_service_name, "--project", projectID, "--format", "json"})).Array()
			nbServices := len(cloudRunServices)
			assert.Equal(1, nbServices, "expected a single Cloud Run services called %s to be deployed, found %d services", server_service_name, nbServices)
			match := utils.GetFirstMatchResult(t, cloudRunServices, "kind", "Service")
			serviceURL := match.Get("status.url").String()
			assert.Truef(strings.HasSuffix(serviceURL, ".run.app"), "unexpected service URL %q", serviceURL)
			t.Log("Cloud Run service is running at", serviceURL)
		}

		{
			// TODO: Check that load balancer endpoint returns 200
		}

		{
			// TODO: Check Firestore accessible from frontend service
		}
	})
	example.Test()
}
