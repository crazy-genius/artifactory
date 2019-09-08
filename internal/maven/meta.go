package maven

type snapshot struct {
	Timestamp string `xml:"timestamp"`
	BuildNumber int `xml:"buildNumber"`
}

type snapshotVersion struct {
	Extension string `xml:"extension"`
	Value string `xml:"value"`
	Updated int64 `xml:"updated"`
}

type versionong struct {
	LastUpdated int64 `xml:"lastUpdated"`
	Snapshot snapshot `xml:"snapshot"`
	SnapshotVersions []snapshotVersion `xml:"snapshotVersions"`
}

type Meta struct {
	ModelVersion string `xml:"model_version, attribute"`
	GroupId string `xml:"groupId"`
	ArtifactId string `xml:"artifactId"`
	Version string `xml:"version"`
	Versioning versionong `xml:"versioning"`
}
