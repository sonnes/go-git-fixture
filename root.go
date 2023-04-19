package gogitfixture

import (
	"fmt"
	"os"
)

// Version reads the version from the VERSION file
func Version() string {
	f, err := os.Open("VERSION")
	if err != nil {
		return "unknown"
	}

	defer f.Close()

	var version string
	_, err = fmt.Fscanf(f, "%s", &version)
	if err != nil {
		return "unknown"
	}

	return version
}
