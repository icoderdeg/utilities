package utilities

import ("fmt"
        "os/exec"
        "bytes"
        "testing"
)

//go test -v
//func TestReadPdf(t *testing.T) 
func TestDisplayGoVersion(t *testing.T) {
	cmd := exec.Command("go", "version")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
	   t.Error(err)
	}
	fmt.Print(string(cmdOutput.Bytes()))
}


