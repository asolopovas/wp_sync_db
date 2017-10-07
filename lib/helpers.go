package lyndbdump

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func errChk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func execCmd(app string, options []string) []byte {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(app, options...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		log.Fatal(stderr.String())
	}
	return stdout.Bytes()
}

func writeToFile(filename string, data []byte) {

	dir := "./tmp"
	if !exists(dir) {
		fmt.Println("'./tmp' directory created...")
		os.MkdirAll(dir, 700)
	}
	err := ioutil.WriteFile(dir+"/"+filename, data, 0644)
	errChk(err)
}
