# Release History

## 3.0.0 (2022-11-21)
### Breaking Changes

- Type of `SynapseSparkJobReference.ReferenceName` has been changed from `*string` to `interface{}`

### Features Added

- New const `TeradataPartitionOptionHash`
- New const `SQLPartitionOptionNone`
- New const `OrcCompressionCodecLzo`
- New const `OraclePartitionOptionDynamicRange`
- New const `CopyBehaviorTypePreserveHierarchy`
- New const `SQLWriteBehaviorEnumInsert`
- New const `CompressionCodecZipDeflate`
- New const `StoredProcedureParameterTypeInt`
- New const `TeradataPartitionOptionDynamicRange`
- New const `ServicePrincipalCredentialTypeServicePrincipalKey`
- New const `CompressionCodecNone`
- New const `SapTablePartitionOptionPartitionOnInt`
- New const `CopyBehaviorTypeFlattenHierarchy`
- New const `SQLWriteBehaviorEnumUpsert`
- New const `CompressionCodecDeflate`
- New const `AvroCompressionCodecXz`
- New const `JSONFormatFilePatternSetOfObjects`
- New const `SQLPartitionOptionPhysicalPartitionsOfTable`
- New const `CompressionCodecLzo`
- New const `DynamicsAuthenticationTypeIfd`
- New const `SapTablePartitionOptionPartitionOnCalendarMonth`
- New const `NetezzaPartitionOptionDataSlice`
- New const `CompressionCodecLz4`
- New const `HdiNodeTypesHeadnode`
- New const `SapTablePartitionOptionNone`
- New const `SQLPartitionOptionDynamicRange`
- New const `JSONWriteFilePatternArrayOfObjects`
- New const `SapTablePartitionOptionPartitionOnTime`
- New const `CompressionCodecGzip`
- New const `SQLDWWriteBehaviorEnumInsert`
- New const `JSONWriteFilePatternSetOfObjects`
- New const `SapTablePartitionOptionPartitionOnCalendarYear`
- New const `CompressionCodecSnappy`
- New const `CopyBehaviorTypeMergeFiles`
- New const `AmazonRdsForOraclePartitionOptionDynamicRange`
- New const `AvroCompressionCodecDeflate`
- New const `SQLWriteBehaviorEnumStoredProcedure`
- New const `DynamicsAuthenticationTypeAADServicePrincipal`
- New const `OrcCompressionCodecNone`
- New const `StoredProcedureParameterTypeInt64`
- New const `CompressionCodecTar`
- New const `AvroCompressionCodecBzip2`
- New const `SapTablePartitionOptionPartitionOnCalendarDate`
- New const `HdiNodeTypesWorkernode`
- New const `NetezzaPartitionOptionNone`
- New const `AvroCompressionCodecNone`
- New const `OraclePartitionOptionPhysicalPartitionsOfTable`
- New const `OrcCompressionCodecSnappy`
- New const `StoredProcedureParameterTypeString`
- New const `SQLDWWriteBehaviorEnumUpsert`
- New const `DynamicsDeploymentTypeOnline`
- New const `SapHanaPartitionOptionSapHanaDynamicRange`
- New const `JSONFormatFilePatternArrayOfObjects`
- New const `StoredProcedureParameterTypeDecimal`
- New const `AvroCompressionCodecSnappy`
- New const `SapHanaPartitionOptionPhysicalPartitionsOfTable`
- New const `AmazonRdsForOraclePartitionOptionNone`
- New const `StoredProcedureParameterTypeDate`
- New const `SapHanaPartitionOptionNone`
- New const `DynamicsDeploymentTypeOnPremisesWithIfd`
- New const `CompressionCodecTarGZip`
- New const `OrcCompressionCodecZlib`
- New const `StoredProcedureParameterTypeGUID`
- New const `StoredProcedureParameterTypeBoolean`
- New const `CompressionCodecBzip2`
- New const `NetezzaPartitionOptionDynamicRange`
- New const `DatasetCompressionLevelOptimal`
- New const `AmazonRdsForOraclePartitionOptionPhysicalPartitionsOfTable`
- New const `HdiNodeTypesZookeeper`
- New const `TeradataPartitionOptionNone`
- New const `ServicePrincipalCredentialTypeServicePrincipalCert`
- New const `DatasetCompressionLevelFastest`
- New const `DynamicsAuthenticationTypeOffice365`
- New const `OraclePartitionOptionNone`
- New type alias `JSONFormatFilePattern`
- New type alias `DynamicsDeploymentType`
- New type alias `OraclePartitionOption`
- New type alias `SQLPartitionOption`
- New type alias `TeradataPartitionOption`
- New type alias `SQLDWWriteBehaviorEnum`
- New type alias `AmazonRdsForOraclePartitionOption`
- New type alias `DatasetCompressionLevel`
- New type alias `SQLWriteBehaviorEnum`
- New type alias `CopyBehaviorType`
- New type alias `DynamicsAuthenticationType`
- New type alias `StoredProcedureParameterType`
- New type alias `SapHanaPartitionOption`
- New type alias `OrcCompressionCodec`
- New type alias `NetezzaPartitionOption`
- New type alias `HdiNodeTypes`
- New type alias `ServicePrincipalCredentialType`
- New type alias `SapTablePartitionOption`
- New type alias `JSONWriteFilePattern`
- New type alias `AvroCompressionCodec`
- New type alias `CompressionCodec`
- New function `PossibleStoredProcedureParameterTypeValues() []StoredProcedureParameterType`
- New function `*CredentialOperationsClient.NewListByFactoryPager(string, string, *CredentialOperationsClientListByFactoryOptions) *runtime.Pager[CredentialOperationsClientListByFactoryResponse]`
- New function `PossibleSapHanaPartitionOptionValues() []SapHanaPartitionOption`
- New function `*Credential.GetCredential() *Credential`
- New function `*CredentialOperationsClient.Get(context.Context, string, string, string, *CredentialOperationsClientGetOptions) (CredentialOperationsClientGetResponse, error)`
- New function `PossibleHdiNodeTypesValues() []HdiNodeTypes`
- New function `PossibleSQLWriteBehaviorEnumValues() []SQLWriteBehaviorEnum`
- New function `PossibleCopyBehaviorTypeValues() []CopyBehaviorType`
- New function `PossibleNetezzaPartitionOptionValues() []NetezzaPartitionOption`
- New function `PossibleDynamicsAuthenticationTypeValues() []DynamicsAuthenticationType`
- New function `PossibleAmazonRdsForOraclePartitionOptionValues() []AmazonRdsForOraclePartitionOption`
- New function `*ManagedIdentityCredential.GetCredential() *Credential`
- New function `PossibleOraclePartitionOptionValues() []OraclePartitionOption`
- New function `*CopyTranslator.GetCopyTranslator() *CopyTranslator`
- New function `NewCredentialOperationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CredentialOperationsClient, error)`
- New function `*CredentialOperationsClient.CreateOrUpdate(context.Context, string, string, string, ManagedIdentityCredentialResource, *CredentialOperationsClientCreateOrUpdateOptions) (CredentialOperationsClientCreateOrUpdateResponse, error)`
- New function `PossibleSapTablePartitionOptionValues() []SapTablePartitionOption`
- New function `PossibleAvroCompressionCodecValues() []AvroCompressionCodec`
- New function `PossibleOrcCompressionCodecValues() []OrcCompressionCodec`
- New function `*ServicePrincipalCredential.GetCredential() *Credential`
- New function `PossibleTeradataPartitionOptionValues() []TeradataPartitionOption`
- New function `PossibleDatasetCompressionLevelValues() []DatasetCompressionLevel`
- New function `PossibleSQLPartitionOptionValues() []SQLPartitionOption`
- New function `*TabularTranslator.GetCopyTranslator() *CopyTranslator`
- New function `PossibleJSONFormatFilePatternValues() []JSONFormatFilePattern`
- New function `PossibleCompressionCodecValues() []CompressionCodec`
- New function `PossibleSQLDWWriteBehaviorEnumValues() []SQLDWWriteBehaviorEnum`
- New function `PossibleServicePrincipalCredentialTypeValues() []ServicePrincipalCredentialType`
- New function `PossibleDynamicsDeploymentTypeValues() []DynamicsDeploymentType`
- New function `*CredentialOperationsClient.Delete(context.Context, string, string, string, *CredentialOperationsClientDeleteOptions) (CredentialOperationsClientDeleteResponse, error)`
- New function `PossibleJSONWriteFilePatternValues() []JSONWriteFilePattern`
- New struct `AdditionalColumns`
- New struct `CopyTranslator`
- New struct `Credential`
- New struct `CredentialListResponse`
- New struct `CredentialOperationsClient`
- New struct `CredentialOperationsClientCreateOrUpdateOptions`
- New struct `CredentialOperationsClientCreateOrUpdateResponse`
- New struct `CredentialOperationsClientDeleteOptions`
- New struct `CredentialOperationsClientDeleteResponse`
- New struct `CredentialOperationsClientGetOptions`
- New struct `CredentialOperationsClientGetResponse`
- New struct `CredentialOperationsClientListByFactoryOptions`
- New struct `CredentialOperationsClientListByFactoryResponse`
- New struct `CredentialResource`
- New struct `DatasetDataElement`
- New struct `DatasetSchemaDataElement`
- New struct `GetDataFactoryOperationStatusResponse`
- New struct `IntegrationRuntimeStatusListResponse`
- New struct `ManagedIdentityCredential`
- New struct `ManagedIdentityCredentialResource`
- New struct `ManagedIdentityTypeProperties`
- New struct `Resource`
- New struct `ServicePrincipalCredential`
- New struct `ServicePrincipalCredentialTypeProperties`
- New struct `StoredProcedureParameter`
- New struct `SubResource`
- New struct `SubResourceDebugResource`
- New struct `TabularTranslator`
- New struct `TypeConversionSettings`
- New field `DisablePublish` in struct `FactoryVSTSConfiguration`
- New field `DisablePublish` in struct `FactoryGitHubConfiguration`
- New field `WorkspaceResourceID` in struct `AzureSynapseArtifactsLinkedServiceTypeProperties`
- New field `DisablePublish` in struct `FactoryRepoConfiguration`
- New field `PythonCodeReference` in struct `SynapseSparkJobActivityTypeProperties`
- New field `FilesV2` in struct `SynapseSparkJobActivityTypeProperties`
- New field `ScriptBlockExecutionTimeout` in struct `ScriptActivityTypeProperties`


## 1.3.0 (2022-09-07)
### Features Added

- New const `NotebookParameterTypeBool`
- New const `NotebookReferenceTypeNotebookReference`
- New const `NotebookParameterTypeString`
- New const `SparkJobReferenceTypeSparkJobDefinitionReference`
- New const `NotebookParameterTypeInt`
- New const `BigDataPoolReferenceTypeBigDataPoolReference`
- New const `NotebookParameterTypeFloat`
- New type alias `NotebookParameterType`
- New type alias `SparkJobReferenceType`
- New type alias `NotebookReferenceType`
- New type alias `BigDataPoolReferenceType`
- New function `*AzureSynapseArtifactsLinkedService.GetLinkedService() *LinkedService`
- New function `PossibleBigDataPoolReferenceTypeValues() []BigDataPoolReferenceType`
- New function `PossibleNotebookParameterTypeValues() []NotebookParameterType`
- New function `*SynapseSparkJobDefinitionActivity.GetExecutionActivity() *ExecutionActivity`
- New function `*GoogleSheetsLinkedService.GetLinkedService() *LinkedService`
- New function `*SynapseNotebookActivity.GetExecutionActivity() *ExecutionActivity`
- New function `PossibleNotebookReferenceTypeValues() []NotebookReferenceType`
- New function `PossibleSparkJobReferenceTypeValues() []SparkJobReferenceType`
- New function `*SynapseNotebookActivity.GetActivity() *Activity`
- New function `*SynapseSparkJobDefinitionActivity.GetActivity() *Activity`
- New struct `AzureSynapseArtifactsLinkedService`
- New struct `AzureSynapseArtifactsLinkedServiceTypeProperties`
- New struct `BigDataPoolParametrizationReference`
- New struct `GoogleSheetsLinkedService`
- New struct `GoogleSheetsLinkedServiceTypeProperties`
- New struct `NotebookParameter`
- New struct `SynapseNotebookActivity`
- New struct `SynapseNotebookActivityTypeProperties`
- New struct `SynapseNotebookReference`
- New struct `SynapseSparkJobActivityTypeProperties`
- New struct `SynapseSparkJobDefinitionActivity`
- New struct `SynapseSparkJobReference`


## 1.2.0 (2022-06-15)
### Features Added

- New field `ClientSecret` in struct `RestServiceLinkedServiceTypeProperties`
- New field `Resource` in struct `RestServiceLinkedServiceTypeProperties`
- New field `Scope` in struct `RestServiceLinkedServiceTypeProperties`
- New field `TokenEndpoint` in struct `RestServiceLinkedServiceTypeProperties`
- New field `ClientID` in struct `RestServiceLinkedServiceTypeProperties`


## 1.1.0 (2022-05-30)
### Features Added

- New function `GlobalParameterResource.MarshalJSON() ([]byte, error)`
- New struct `GlobalParameterListResponse`
- New struct `GlobalParameterResource`
- New struct `GlobalParametersClientCreateOrUpdateOptions`
- New struct `GlobalParametersClientCreateOrUpdateResponse`
- New struct `GlobalParametersClientDeleteOptions`
- New struct `GlobalParametersClientDeleteResponse`
- New struct `GlobalParametersClientGetOptions`
- New struct `GlobalParametersClientGetResponse`
- New struct `GlobalParametersClientListByFactoryOptions`
- New struct `GlobalParametersClientListByFactoryResponse`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).## 1.2.0 (2022-06-15)
### Features Added

- New field `ClientSecret` in struct `RestServiceLinkedServiceTypeProperties`
- New field `Resource` in struct `RestServiceLinkedServiceTypeProperties`
- New field `Scope` in struct `RestServiceLinkedServiceTypeProperties`
- New field `TokenEndpoint` in struct `RestServiceLinkedServiceTypeProperties`
- New field `ClientID` in struct `RestServiceLinkedServiceTypeProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).