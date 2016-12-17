
package main

import (
	"fmt"
	"flag"
//	"github.com/biogo/hts/bam"
)


func main() {
	flag.Parse()
	bamFilename := flag.Arg(0)

	fmt.Println(bamFilename)
}
