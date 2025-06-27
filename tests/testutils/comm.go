package testutils

import (
	"fmt"
	"os"
	"strings"
)

// GoTestComm returns the comm of the `go test` process from /proc/$PID/comm
func GoTestComm() string {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/comm", os.Getpid()))
	if err != nil {
		return "integration.tes"
	}
	return strings.TrimSuffix(string(data), "\n")
}
