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
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/gcloud"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/tft"
	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/infra/blueprint-test/pkg/utils"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

var providedProject = flag.String("override_project", "", "use the specified project, instead of the one from the setup stage. Intended for development use only.")

func getProjectId(tst *tft.TFBlueprintTest) string {
	if *providedProject != "" {
		return *providedProject
	}
	return tst.GetTFSetupStringOutput("project_id")
}

func TestSimpleExample(t *testing.T) {
	example := tft.NewTFBlueprintTest(t)

	example.DefineVerify(func(assert *assert.Assertions) {
		// DefaultVerify asserts no resource changes exist after apply.
		// It helps ensure that a second "terraform apply" wouldn't result in resource deletions/replacements.
		example.DefaultVerify(assert)

		projectID := getProjectId(example)
		t.Logf("Using Project: %s\n", projectID)
		server_service_name := terraform.OutputRequired(t, example.GetTFOptions(), "run_service_name")

		frontend := example.GetStringOutput("frontend_url")
		t.Logf("frontend url: %s\n", frontend)
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

		// Check that load balancer endpoint returns 200
		{
			resp, err := http.Get(frontend)
			if err != nil {
				t.Errorf("http request to frontend (%s) failed: %v", frontend, err)
			}
			if resp.StatusCode != 200 {
				t.Errorf("unexpected response from frontend(%s): %s", frontend, resp.Status)
			}
		}
		// Check that our loadbalancer is wired correctly.
		{
			feUrl, err := url.Parse(frontend)
			if err != nil {
				t.Fatalf("invalid frontend url found: %s", frontend)
			}
			lbjson := gcloud.Run(t, "compute url-maps list",
				gcloud.WithCommonArgs([]string{
					"--filter", "hostRules.hosts=" + feUrl.Host,
					"--project", projectID,
					"--format", "json",
				}))
			lbBackendService := lbjson.Get("0.defaultService").String()
			lbBackendBucket := lbjson.Get("0.pathMatchers.0.pathRules.#(paths.#(=\"/google-cloud-icons/*\")).service").String()
			fmt.Printf("LB Backend Service: %s\n", lbBackendService)
			fmt.Printf("LB Backend Bucket: %s\n", lbBackendBucket)

			// Verify Cloud Run backend wiring.
			// url-map -> backend-service -> NEG -> Cloud Run Service
			backendJson := gcloud.Run(t, "compute backend-services describe "+lbBackendService,
				gcloud.WithCommonArgs([]string{
					"--project", projectID,
					"--format", "json",
				}))
			neg := backendJson.Get("backends.0.group").String()
			fmt.Println(neg)
			negJson := gcloud.Run(t, "compute network-endpoint-groups describe "+neg,
				gcloud.WithCommonArgs([]string{
					"--project", projectID,
					"--format", "json",
				}))
			assert.Equal(server_service_name, negJson.Get("cloudRun.service").String())

			// Verify GCS bucket wiring.
			// url-map -> backend-bucket -> gcs bucket name
			beBucketJson := gcloud.Run(t, "compute backend-buckets describe "+lbBackendBucket,
				gcloud.WithCommonArgs([]string{
					"--project", projectID,
					"--format", "json",
				}))
			// test that the bucket has files in it
			bucketName := beBucketJson.Get("bucketName").String()
			files := gcloud.Run(t, fmt.Sprintf("alpha storage ls gs://%s/google-cloud-icons/", bucketName),
				gcloud.WithCommonArgs([]string{
					"--project", projectID,
					"--json",
				})).Array()
			assert.NotEmpty(files, "expected files in gcs bucket (%s), found %s", bucketName, files)

		}
		// Hit the url that verifies Firestore connectivity
		{
			fscheckUrl, err := url.JoinPath(frontend, "/api/fscheck")
			if err != nil {
				t.Errorf("invalid url for fscheck: %v", err)
			}
			resp, err := http.Get(fscheckUrl)
			if err != nil {
				t.Errorf("http request to frontend (%s) failed: %v", frontend, err)
			}
			if resp.StatusCode != 200 {
				t.Errorf("unexpected response from frontend(%s): %s", frontend, resp.Status)
			}
		}
	})
	example.Test()
}
