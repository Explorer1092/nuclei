package extractors

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Explorer1092/nuclei/v3/pkg/operators/common/dsl"
	"github.com/Knetic/govaluate"
	"github.com/itchyny/gojq"
<<<<<<< HEAD
=======
<<<<<<< HEAD:v2/pkg/operators/extractors/compile.go
	"github.com/Explorer1092/nuclei/v2/pkg/operators/common/dsl"
=======
	"github.com/projectdiscovery/nuclei/v3/pkg/operators/common/dsl"
>>>>>>> 419f08f61ce5ca2d3f0eae9fe36dc7c44c1f532a:pkg/operators/extractors/compile.go
>>>>>>> projectdiscovery-main
)

// CompileExtractors performs the initial setup operation on an extractor
func (e *Extractor) CompileExtractors() error {
	// Set up the extractor type
	computedType, err := toExtractorTypes(e.GetType().String())
	if err != nil {
		return fmt.Errorf("unknown extractor type specified: %s", e.Type)
	}
	e.extractorType = computedType
	// Compile the regexes
	for _, regex := range e.Regex {
		compiled, err := regexp.Compile(regex)
		if err != nil {
			return fmt.Errorf("could not compile regex: %s", regex)
		}
		e.regexCompiled = append(e.regexCompiled, compiled)
	}
	for i, kval := range e.KVal {
		e.KVal[i] = strings.ToLower(kval)
	}

	for _, query := range e.JSON {
		query, err := gojq.Parse(query)
		if err != nil {
			return fmt.Errorf("could not parse json: %s", query)
		}
		compiled, err := gojq.Compile(query)
		if err != nil {
			return fmt.Errorf("could not compile json: %s", query)
		}
		e.jsonCompiled = append(e.jsonCompiled, compiled)
	}

	for _, dslExp := range e.DSL {
		compiled, err := govaluate.NewEvaluableExpressionWithFunctions(dslExp, dsl.HelperFunctions)
		if err != nil {
			return &dsl.CompilationError{DslSignature: dslExp, WrappedError: err}
		}
		e.dslCompiled = append(e.dslCompiled, compiled)
	}

	if e.CaseInsensitive {
		if e.GetType() != KValExtractor {
			return fmt.Errorf("case-insensitive flag is supported only for 'kval' extractors (not '%s')", e.Type)
		}
		for i := range e.KVal {
			e.KVal[i] = strings.ToLower(e.KVal[i])
		}
	}

	return nil
}
