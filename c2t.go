
package main

import (
	"fmt"
	"flag"
	"os"
//	"github.com/biogo/hts/bam"
)


func main() {
	help := flag.Bool("help", false, "Print this usage message.")
	flag.Parse()

	bamFilename := flag.Arg(0)

	if *help {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(bamFilename)
}
