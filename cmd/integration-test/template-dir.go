package main

import (
	"os"

<<<<<<< HEAD:v2/cmd/integration-test/template-dir.go
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 7f556f8e33080a9eb2e52a9c14fa2e16f32f62c3:cmd/integration-test/template-dir.go
	errorutil "github.com/projectdiscovery/utils/errors"
)

var templatesDirTestCases = []TestCaseInfo{
	{Path: "protocols/dns/cname-fingerprint.yaml", TestCase: &templateDirWithTargetTest{}},
}

type templateDirWithTargetTest struct{}

// Execute executes a test case and returns an error if occurred
func (h *templateDirWithTargetTest) Execute(filePath string) error {
	tempdir, err := os.MkdirTemp("", "nuclei-update-dir-*")
	if err != nil {
		return errorutil.NewWithErr(err).Msgf("failed to create temp dir")
	}
	defer os.RemoveAll(tempdir)

	results, err := testutils.RunNucleiTemplateAndGetResults(filePath, "8x8exch02.8x8.com", debug, "-ud", tempdir)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}
