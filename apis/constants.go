/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apis

import (
	core "k8s.io/api/core/v1"
)

const (
	Finalizer = "auditor.kubeshield.to"

	// Specifies the path where auth is enabled
	AuthPathKey = "auditor.kubeshield.to/auth-path"

	// required fields:
	// - Secret.Data["token"] - a vault token
	SecretTypeTokenAuth core.SecretType = "auditor.kubeshield.to/token"

	// required for SecretTypeTokenAut
	TokenAuthTokenKey = "token"

	// required fields:
	// - Secret.Data["access_key_id"] - aws access key id
	// - Secret.Data["secret_access_key"] - aws access secret key
	//
	// optional fields:
	// - Secret.Annotations["auditor.kubeshield.to/aws.header-value"] - specifies the header value that required if X-Vault-AWS-IAM-Server-ID Header is set
	// - Secret.Annotations["auditor.kubeshield.to/auth-path"] - Specifies the path where aws auth is enabled
	SecretTypeAWSAuth core.SecretType = "auditor.kubeshield.to/aws"

	// required for SecretTypeAWSAuth
	AWSAuthAccessKeyIDKey = "access_key_id"
	// required for SecretTypeAWSAuth
	AWSAuthAccessSecretKey = "secret_access_key"
	// optional for SecretTypeAWSAuth
	AWSAuthSecurityTokenKey = "security_token"

	// Specifies the header value that required if X-Vault-AWS-IAM-Server-ID Header is set
	// optional for annotation for  SecretTypeAWSAuth
	AWSHeaderValueKey = "auditor.kubeshield.to/aws.header-value"

	// required fields:
	// - Secret.Data["sa.json"] - gcp access secret key
	//
	// optional fields:
	// - Secret.Annotations["auditor.kubeshield.to/auth-path"] - Specifies the path where gcp auth is enabled
	SecretTypeGCPAuth core.SecretType = "auditor.kubeshield.to/gcp"
	// required for SecretTypeGCPAuth
	GCPAuthSACredentialJson = "sa.json"

	// - Secret.Data["msiToken"] - azure managed service identity (MSI)  jwt token
	//
	// optional fields:
	// - Secret.Annotations["auditor.kubeshield.to/azure.subscription-id"] - The subscription ID for the machine that generated the MSI token. This information can be obtained through instance metadata.
	// - Secret.Annotations["auditor.kubeshield.to/azure.resource-group-name"] - The resource group for the machine that generated the MSI token. This information can be obtained through instance metadata.
	// - Secret.Annotations["auditor.kubeshield.to/azure.vm-name"] - The virtual machine name for the machine that generated the MSI token. This information can be obtained through instance metadata. If vmss_name is provided, this value is ignored.
	// - Secret.Annotations["auditor.kubeshield.to/azure.vmss-name"] - The virtual machine scale set name for the machine that generated the MSI token. This information can be obtained through instance metadata.
	SecretTypeAzureAuth = "auditor.kubeshield.to/azure"

	// required for SecretTypeAzureAuth
	AzureMSIToken = "msiToken"

	// optional for SecretTypeAzureAuth
	AzureSubscriptionId    = "auditor.kubeshield.to/azure.subscription-id"
	AzureResourceGroupName = "auditor.kubeshield.to/azure.resource-group-name"
	AzureVmName            = "auditor.kubeshield.to/azure.vm-name"
	AzureVmssName          = "auditor.kubeshield.to/azure.vmss-name"
)
