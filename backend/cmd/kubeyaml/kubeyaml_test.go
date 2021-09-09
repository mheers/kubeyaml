package main

import (
	"bytes"
	"os"
	"path"
	"testing"
)

func TestIntegrations(t *testing.T) {
	testcases := []struct {
		filename       string
		extraArgs      string
		shouldValidate bool
	}{
		// missing a selector.
		{filename: "issue-6.yaml", shouldValidate: false},
		// volume list item is lacking a name
		{filename: "issue-8.yaml", shouldValidate: false},
		// type Airflow is invalid. But we don't validate data
		{filename: "issue-9.yaml", shouldValidate: true},
		// first document is valid, second is not
		{filename: "issue-7.yaml", shouldValidate: false},
		// first two documents are valid, but there is an empty one (which fails)
		{filename: "issue-7_2.yaml", shouldValidate: false},
		// unless -ignore-empty is set
		{filename: "issue-7_2.yaml", extraArgs: "-ignore-empty", shouldValidate: true},
	}

	for _, tc := range testcases {
		t.Run(tc.filename, func(t *testing.T) {
			f, err := os.Open(path.Join("testdata", tc.filename))
			if err != nil {
				t.Fatal(err)
			}
			var b bytes.Buffer
			err = run(f, &b, "-silent", tc.extraArgs)
			if tc.shouldValidate && err != nil {
				t.Fatal(err)
			}
			if !tc.shouldValidate && err == nil {
				t.Fatal("expected error but didn't get one")
			}
		})
	}
	// open file
	// pipe to stdin
	// run main.go
	// expect no errors
}
