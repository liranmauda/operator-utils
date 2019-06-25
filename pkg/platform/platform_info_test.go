package platform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOCPVersionHelpers(t *testing.T) {

	ocpTestVersions := []struct {
		version string
		major   string
		minor   string
		build   string
	}{
		{"3.11.69", "3", "11", "69"},
		{"4.1.0-rc.1", "4", "1", "0-rc.1"},
		{"1.2.3.4.5.6", "1", "2", "3.4.5.6"},
	}

	for _, v := range ocpTestVersions {

		info := PlatformInfo{OCPVersion: v.version}
		assert.Equal(t, v.major, info.OCPMajorVersion(), "OCPMajorVersion mismatch")
		assert.Equal(t, v.minor, info.OCPMinorVersion(), "OCPMinorVersion mismatch")
		assert.Equal(t, v.build, info.OCPBuildVersion(), "OCPBuildVersion mismatch")
	}
}

func TestK8SVersionHelpers(t *testing.T) {

	ocpTestVersions := []struct {
		version string
		major   string
		minor   string
	}{
		{"1.11+", "1", "11+"},
		{"1.13+", "1", "13+"},
	}

	for _, v := range ocpTestVersions {

		info := PlatformInfo{K8SVersion: v.version}
		assert.Equal(t, v.major, info.K8SMajorVersion(), "K8SMajorVersion mismatch")
		assert.Equal(t, v.minor, info.K8SMinorVersion(), "K8SMinorVersion mismatch")
	}
}

func TestPlatformInfo_ApproximateOpenShiftVersion(t *testing.T) {

	cases := []struct {
		label              string
		info               PlatformInfo
		expectedOCPVersion string
	}{
		{
			label:              "case 1",
			info:               PlatformInfo{},
			expectedOCPVersion: "",
		},
		{
			label:              "case 2",
			info:               PlatformInfo{Name: Kubernetes},
			expectedOCPVersion: "",
		},
		{
			label:              "case 3",
			info:               PlatformInfo{Name: OpenShift, K8SVersion: "1.10+"},
			expectedOCPVersion: "3.10",
		},
		{
			label:              "case 4",
			info:               PlatformInfo{Name: OpenShift, K8SVersion: "1.11+"},
			expectedOCPVersion: "3.11",
		},
		{
			label:              "case 5",
			info:               PlatformInfo{Name: OpenShift, K8SVersion: "1.13+"},
			expectedOCPVersion: "4.1",
		},
		{
			label:              "case 6",
			info:               PlatformInfo{Name: OpenShift, K8SVersion: "1.99"},
			expectedOCPVersion: "",
		},
		{
			label:              "case 7",
			info:               PlatformInfo{Name: OpenShift},
			expectedOCPVersion: "",
		},
	}

	for _, v := range cases {
		v.info.ApproximateOpenShiftVersion()
		assert.Equal(t, v.expectedOCPVersion, v.info.OCPVersion, v.label+": expected OCP version to match")
	}
}

func TestPlatformInfo_String(t *testing.T) {

	info := PlatformInfo{Name: OpenShift, OCPVersion: "1.1.1+", K8SVersion: "456", OS: "foo/bar"}

	assert.Equal(t, "PlatformInfo [Name: OpenShift, OCPVersion: 1.1.1+, K8SVersion: 456, OS: foo/bar]",
		info.String(), "PlatformInfo String() yields malformed result of %s", info.String())
}
