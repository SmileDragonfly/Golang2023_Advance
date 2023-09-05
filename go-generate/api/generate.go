// generate.go
//
//go:generate go run generate.go
package api

import (
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("app_build_time.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "package api")
	fmt.Fprintln(file, "")
	fmt.Fprintf(file, "var BuildTime string = \"%s\"", time.Now().Format("2006-01-02 15:04:05"))
}
