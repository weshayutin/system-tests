package nhcparams

import "time"

const (
	// UpgradeSubName is the Subscription name used in the NHC upgrade test.
	UpgradeSubName = "nhc-upgrade-sub"
	// SNRUpgradeSubName is the Subscription name used to install SNR when the
	// NHC upgrade test needs to bootstrap its remediation provider.
	SNRUpgradeSubName = "snr-nhc-upgrade-sub"

	// SNRPackage is the OLM package name for the Self Node Remediation operator.
	SNRPackage = "self-node-remediation"
	// SNRCatalog is the CatalogSource that provides the SNR operator.
	SNRCatalog = "community-operators"
	// SNRCSVNamePattern identifies SNR CSVs created by the bootstrap Subscription.
	SNRCSVNamePattern = "self-node-remediation"
	// SNRDeploymentName is the SNR controller Deployment name.
	SNRDeploymentName = "self-node-remediation-controller-manager"

	// NHCUpgradeTestName is the NHC CR name used at the remediation checkpoints
	// of the upgrade test.
	NHCUpgradeTestName = "nhc-upgrade-test"

	// UpgradeRemediationCompletionTimeout is the remediation-completion timeout
	// used at the upgrade test's checkpoints.
	UpgradeRemediationCompletionTimeout = 20 * time.Minute
)
