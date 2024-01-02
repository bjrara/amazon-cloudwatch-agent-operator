// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package opampbridge

import (
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/aws/amazon-cloudwatch-agent-operator/internal/manifests"
)

const (
	ComponentOpAMPBridge = "opentelemetry-opamp-bridge"
)

// Build creates the manifest for the OpAMPBridge resource.
func Build(params manifests.Params) ([]client.Object, error) {
	var resourceManifests []client.Object
	resourceFactories := []manifests.K8sManifestFactory{
		manifests.FactoryWithoutError(Deployment),
		manifests.Factory(ConfigMap),
		manifests.FactoryWithoutError(ServiceAccount),
		manifests.FactoryWithoutError(Service),
	}
	for _, factory := range resourceFactories {
		res, err := factory(params)
		if err != nil {
			return nil, err
		} else if manifests.ObjectIsNotNil(res) {
			resourceManifests = append(resourceManifests, res)
		}
	}
	return resourceManifests, nil
}
