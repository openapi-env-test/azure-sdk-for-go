# Unreleased

## Breaking Changes

### Signature Changes

#### Struct Fields

1. OrchestrationServiceStateInput.ServiceName changed type from OrchestrationServiceNames to *string

## Additive Changes

### New Constants

1. SecurityProfileType.SecurityProfileTypeEncryptedVMGuestStateOnlyWithPmk
1. SecurityProfileType.SecurityProfileTypeEncryptedWithCmk
1. SecurityProfileType.SecurityProfileTypeEncryptedWithPmk

### New Funcs

1. PossibleSecurityProfileTypeValues() []SecurityProfileType

### Struct Changes

#### New Structs

1. OSDiskImageSecurityProfile

#### New Struct Fields

1. OSDiskImageEncryption.SecurityProfile
