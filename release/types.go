package release

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Release defaults
const (
	DefaultDeprecationHours = 48 // 2 days
	DefaultAutoUpdateHours  = 24
)

// ReleaseSummary contains concise release info.
type ReleaseSummary struct {
	Commit  string `json:"commit,omitempty"`
	IPFS    string `json:"ipfs,omitempty"`
	Version string `json:"version,omitempty"`
}

// ReleaseInfo contains the release response from the updater.
type ReleaseInfo struct {
	FromBuild bool            `json:"fromBuild"`
	IPFS      string          `json:"ipfs"`
	Manifest  ReleaseManifest `json:"manifest"`
}

// String implements fmt.Stringer interface.
func (releaseInfo *ReleaseInfo) String() string {
	if releaseInfo == nil {
		return ""
	}
	b, _ := json.Marshal(releaseInfo)
	return string(b)
}

// ReleaseInfoFromString parses the string.
func ReleaseInfoFromString(s string) *ReleaseInfo {
	if len(s) == 0 {
		log.Warn("empty release info")
		return nil
	}
	var releaseInfo ReleaseInfo
	if err := json.Unmarshal([]byte(s), &releaseInfo); err != nil {
		return &releaseInfo
	}
	if len(releaseInfo.Manifest.Release.Commit) > 0 {
		LogReleaseInfo(&releaseInfo)
	}
	return &releaseInfo
}

// LogReleaseInfo logs the release info.
func LogReleaseInfo(releaseInfo *ReleaseInfo) {
	if releaseInfo == nil {
		return
	}
	log.WithFields(log.Fields{
		"commit":    releaseInfo.Manifest.Release.Commit,
		"version":   releaseInfo.Manifest.Release.Version,
		"timestamp": releaseInfo.Manifest.Release.Timestamp,
		"ipfs":      releaseInfo.IPFS,
		"fromBuild": releaseInfo.FromBuild,
	}).Info("release info")
}

// MakeSummaryFromReleaseInfo transforms the release info into a more compact and common form.
func MakeSummaryFromReleaseInfo(releaseInfo *ReleaseInfo) *ReleaseSummary {
	if releaseInfo == nil {
		return nil
	}
	return &ReleaseSummary{
		Commit:  releaseInfo.Manifest.Release.Commit,
		IPFS:    releaseInfo.IPFS,
		Version: releaseInfo.Manifest.Release.Version,
	}
}

// ReleaseManifest contains the latest info about the latest scanner version.
type ReleaseManifest struct {
	Release Release `json:"release" yaml:"release"`
}

// Release contains release data.
type Release struct {
	Timestamp  string          `json:"timestamp" yaml:"timestamp"`
	Repository string          `json:"repository" yaml:"repository"`
	Version    string          `json:"version" yaml:"version"`
	Commit     string          `json:"commit" yaml:"commit"`
	Services   ReleaseServices `json:"services" yaml:"services"`
	Config     ReleaseConfig   `json:"config" yaml:"config"`
}

// ReleaseServices are the services to run for scanner node.
type ReleaseServices struct {
	Updater    string `json:"updater" yaml:"updater"`
	Supervisor string `json:"supervisor" yaml:"supervisor"`
}

// ReleaseConfig contains release config information.
type ReleaseConfig struct {
	AutoUpdateInHours int               `json:"autoUpdateInHours" yaml:"autoUpdateInHours"`
	DeprecationPolicy DeprecationPolicy `json:"deprecationPolicy" yaml:"deprecationPolicy"`
}

// DeprecationPolicy defines the deprecation policy of this release.
type DeprecationPolicy struct {
	SupportedVersions []string `json:"supportedVersions" yaml:"supportedVersions"`
	ActivatesInHours  int      `json:"activatesInHours" yaml:"activatesInHours"`
}
