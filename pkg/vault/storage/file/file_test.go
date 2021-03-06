/*
Copyright The KubeVault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"fmt"
	"testing"

	api "kubevault.dev/operator/apis/kubevault/v1alpha1"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/fake"
	ofst "kmodules.xyz/offshoot-api/api/v1"
)

func TestOptions_GetStorageConfig(t *testing.T) {
	kfake := fake.NewSimpleClientset()
	vaultServer := &api.VaultServer{}
	opts, err := NewOptions(kfake, vaultServer, &api.FileSpec{
		Path: "/test",
		VolumeClaimTemplate: ofst.PersistentVolumeClaim{
			PartialObjectMeta: ofst.PartialObjectMeta{
				Name: "test-pvc",
			},
			Spec: v1.PersistentVolumeClaimSpec{},
		},
	})
	assert.Nil(t, err)

	out := `
storage "file" {
path = "/test"
}
`
	t.Run("file system storage config", func(t *testing.T) {
		got, err := opts.GetStorageConfig()
		assert.Nil(t, err)
		if !assert.Equal(t, out, got) {
			fmt.Println("expected:", out)
			fmt.Println("got:", got)
		}
	})
}
