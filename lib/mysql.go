package lyndbdump

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// MySQL represents container for mysql arguments
type MySQL struct {
	// DB Host (e.g. 127.0.0.1)
	Host string
	// DB Port (e.g. 3306)
	Port int64
	// DB Name
	DB string
	// DB User
	User string
	// DB Password
	Password string
	// Extra mysqldump options
	// e.g []string{"--extended-insert"}
	Options []string
}

// Import imports data into Local database
func (x MySQL) Import(data string, opt LocalDb) {
	var stderr bytes.Buffer
	cmd := exec.Command("mysql", x.importOptions(opt)...)
	// Set < pipe variable
	stdin, err := cmd.StdinPipe()
	errChk(err)

	cmd.Stderr = &stderr
	cmd.Start()

	// Write data to pipe
	io.WriteString(stdin, data)
	stdin.Close()
	fmt.Println("Importing " + x.DB + " to localhost...")

	// Log mysql error
	if err := cmd.Wait(); err != nil {
		log.Fatal(stderr.String())
	} else {
		fmt.Println("Importing complete")
	}
}

// Dump - executes mysqldump and return the output
func (x MySQL) Dump() []byte {
	fmt.Println("Copying '" + x.DB + "' database from remote host...")
	return execCmd("mysqldump", x.dumpOptions())
}

// importOptions - return arguments for mysql
func (x MySQL) importOptions(opt LocalDb) []string {
	options := x.Options
	options = append(options, fmt.Sprintf(`-u%v`, opt.User))
	options = append(options, fmt.Sprintf(`-p%v`, opt.Pass))
	options = append(options, fmt.Sprintf(`-h%v`, opt.Host))
	options = append(options, fmt.Sprintf(`-P%v`, x.Port))
	options = append(options, x.DB)
	return options
}

// dumpOptions - returns arguments for mysqldump
func (x MySQL) dumpOptions() []string {
	options := x.Options
	options = append(options, fmt.Sprintf(`-h%v`, x.Host))
	options = append(options, fmt.Sprintf(`-P%v`, x.Port))
	options = append(options, fmt.Sprintf(`-u%v`, x.User))
	options = append(options, fmt.Sprintf(`-p%v`, x.Password))
	options = append(options, x.DB)
	return options
}
