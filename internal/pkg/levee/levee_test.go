// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/google/go-flow-levee/internal/pkg/debug"
	"golang.org/x/tools/go/analysis/analysistest"
)

var debugging *bool = flag.Bool("debug", false, "run the debug analyzer")

func TestLevee(t *testing.T) {
	dataDir := analysistest.TestData()
	if err := Analyzer.Flags.Set("config", dataDir+"/test-config.json"); err != nil {
		t.Error(err)
	}
	testsDir := filepath.Join(dataDir, "src/example.com/tests")
	patterns := findTestPatterns(t, testsDir)
	if *debugging {
		Analyzer.Requires = append(Analyzer.Requires, debug.Analyzer)
	}
	analysistest.Run(t, dataDir, Analyzer, patterns...)
}

func findTestPatterns(t *testing.T, testsDir string) (patterns []string) {
	t.Helper()
	files, err := ioutil.ReadDir(testsDir)
	if err != nil {
		t.Fatalf("Failed to read tests dir (%s): %v", testsDir, err)
	}
	for _, f := range files {
		patterns = append(patterns, filepath.Join(testsDir, f.Name()))
	}
	return
}
