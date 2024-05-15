package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/Explorer1092/nuclei/v3/pkg/templates"
	"github.com/Explorer1092/nuclei/v3/pkg/templates/signer"
	"github.com/Explorer1092/nuclei/v3/pkg/testutils"
)

var workflowTestcases = []TestCaseInfo{
	{Path: "workflow/basic.yaml", TestCase: &workflowBasic{}},
	{Path: "workflow/condition-matched.yaml", TestCase: &workflowConditionMatched{}},
	{Path: "workflow/condition-unmatched.yaml", TestCase: &workflowConditionUnmatch{}},
	{Path: "workflow/matcher-name.yaml", TestCase: &workflowMatcherName{}},
	{Path: "workflow/complex-conditions.yaml", TestCase: &workflowComplexConditions{}},
	{Path: "workflow/http-value-share-workflow.yaml", TestCase: &workflowHttpKeyValueShare{}},
	{Path: "workflow/dns-value-share-workflow.yaml", TestCase: &workflowDnsKeyValueShare{}},
	{Path: "workflow/code-value-share-workflow.yaml", TestCase: &workflowCodeKeyValueShare{}, DisableOn: isCodeDisabled}, // isCodeDisabled declared in code.go
	{Path: "workflow/multiprotocol-value-share-workflow.yaml", TestCase: &workflowMultiProtocolKeyValueShare{}},
	{Path: "workflow/shared-cookie.yaml", TestCase: &workflowSharedCookies{}},
}

func init() {
	// sign code templates (unless they are disabled)
	if !isCodeDisabled() {
		// allow local file access to load content of file references in template
		// in order to sign them for testing purposes
		templates.TemplateSignerLFA()

		// testCertFile and testKeyFile are declared in code.go
		tsigner, err := signer.NewTemplateSignerFromFiles(testCertFile, testKeyFile)
		if err != nil {
			panic(err)
		}

		// only the code templates are necessary to be signed
		var templatesToSign = []string{
			"workflow/code-template-1.yaml",
			"workflow/code-template-2.yaml",
		}
		for _, templatePath := range templatesToSign {
			if err := templates.SignTemplate(tsigner, templatePath); err != nil {
				log.Fatalf("Could not sign template %v got: %s\n", templatePath, err)
			}
		}
	}
}

type workflowBasic struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowBasic) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 2)
}

type workflowConditionMatched struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowConditionMatched) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}

type workflowConditionUnmatch struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowConditionUnmatch) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 0)
}

type workflowMatcherName struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowMatcherName) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}

type workflowComplexConditions struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowComplexConditions) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}

	for _, result := range results {
		if !strings.Contains(result, "test-matcher-3") {
			return fmt.Errorf("incorrect result: the \"basic-get-third:test-matcher-3\" and only that should be matched!\nResults:\n\t%s", strings.Join(results, "\n\t"))
		}
	}
	return expectResultsCount(results, 2)
}

type workflowHttpKeyValueShare struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowHttpKeyValueShare) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/path1", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "href=\"test-value\"")
	})
	router.GET("/path2", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		body, _ := io.ReadAll(r.Body)
		fmt.Fprintf(w, "%s", body)
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}

type workflowDnsKeyValueShare struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowDnsKeyValueShare) Execute(filePath string) error {
	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, "http://scanme.sh", debug)
	if err != nil {
		return err
	}

	// no results - ensure that the variable sharing works
	return expectResultsCount(results, 1)
}

type workflowCodeKeyValueShare struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowCodeKeyValueShare) Execute(filePath string) error {
	// provide the Certificate File that the code templates are signed with
	certEnvVar := signer.CertEnvVarName + "=" + testCertFile

	results, err := testutils.RunNucleiArgsWithEnvAndGetResults(debug, []string{certEnvVar}, "-workflows", filePath, "-target", "input", "-code")
	if err != nil {
		return err
	}

	return expectResultsCount(results, 1)
}

type workflowMultiProtocolKeyValueShare struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowMultiProtocolKeyValueShare) Execute(filePath string) error {
	router := httprouter.New()
	// the response of path1 contains a domain that will be extracted and shared with the second template
	router.GET("/path1", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "href=\"blog.projectdiscovery.io\"")
	})
	// path2 responds with the value of the "extracted" query parameter, e.g.: /path2?extracted=blog.projectdiscovery.io => blog.projectdiscovery.io
	router.GET("/path2", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintf(w, "%s", r.URL.Query().Get("extracted"))
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}

	return expectResultsCount(results, 2)
}

type workflowSharedCookies struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowSharedCookies) Execute(filePath string) error {
	handleFunc := func(name string, w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		cookie := &http.Cookie{Name: name, Value: name}
		http.SetCookie(w, cookie)
	}

	var gotCookies []string
	router := httprouter.New()
	router.GET("/http1", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handleFunc("http1", w, r, p)
	})
	router.GET("/http2", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handleFunc("http2", w, r, p)
	})
	router.GET("/headless1", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handleFunc("headless1", w, r, p)
	})
	router.GET("/http3", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		for _, cookie := range r.Cookies() {
			gotCookies = append(gotCookies, cookie.Name)
		}
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	_, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug, "-headless")
	if err != nil {
		return err
	}

	return expectResultsCount(gotCookies, 3)
}
