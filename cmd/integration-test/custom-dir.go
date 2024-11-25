package main

import (
	"os"

<<<<<<< HEAD
	"github.com/Explorer1092/nuclei/v3/pkg/testutils"
=======
<<<<<<< HEAD:v2/cmd/integration-test/custom-dir.go
	"github.com/Explorer1092/nuclei/v2/pkg/testutils"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/testutils"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:cmd/integration-test/custom-dir.go
>>>>>>> projectdiscovery-main
)

type customConfigDirTest struct{}

var customConfigDirTestCases = []TestCaseInfo{
	{Path: "protocols/dns/cname-fingerprint.yaml", TestCase: &customConfigDirTest{}},
}

// Execute executes a test case and returns an error if occurred
func (h *customConfigDirTest) Execute(filePath string) error {
	customTempDirectory, err := os.MkdirTemp("", "")
	if err != nil {
		return err
	}
	defer os.RemoveAll(customTempDirectory)
	results, err := testutils.RunNucleiBareArgsAndGetResults(debug, []string{"NUCLEI_CONFIG_DIR=" + customTempDirectory}, "-t", filePath, "-u", "8x8exch02.8x8.com")
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return nil
	}
	files, err := os.ReadDir(customTempDirectory)
	if err != nil {
		return err
	}
	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return expectResultsCount(fileNames, 4)
}
