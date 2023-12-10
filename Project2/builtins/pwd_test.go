// // builtins/pwd_test.go
// package builtins

// import (
// 	"bytes"
// 	"testing"
// 	"os"
	
// )

// func TestPrintWorkingDirectory(t *testing.T) {
// 	t.Parallel()

// 	// Run PrintWorkingDirectory
// 	var out bytes.Buffer
// 	err := PrintWorkingDirectory(&out)
// 	if err != nil {
// 		t.Fatalf("Error printing working directory: %v", err)
// 	}

// 	// Check if the output contains a non-empty working directory
// 	actual := out.String()
// 	if len(actual) == 0 {
// 		t.Error("Printed working directory is empty")
// 	}

// 	// Print the working directory and printed working directory for debugging
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		t.Fatalf("Error getting current working directory: %v", err)
// 	}

// 	t.Logf("Current working directory: %s", wd)
// 	t.Logf("Printed working directory: %s", actual)
// }

// builtins/pwd_test.go
package builtins

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
	t.Parallel()

	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "pwd_test")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up: remove the temporary directory

	// Change the working directory to the temporary directory
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Error changing working directory: %v", err)
	}

	// Run PrintWorkingDirectory
	var out bytes.Buffer
	err = PrintWorkingDirectory(&out)
	if err != nil {
		t.Fatalf("Error printing working directory: %v", err)
	}

	// Check if the output contains a non-empty working directory
	actual := out.String()
	if len(actual) == 0 {
		t.Error("Printed working directory is empty")
	}

	// Print the working directory and printed working directory for debugging
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory: %v", err)
	}

	t.Logf("Current working directory: %s", wd)
	t.Logf("Printed working directory: %s", actual)
}
