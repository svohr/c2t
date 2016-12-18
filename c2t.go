
package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/biogo/hts/bam"
)

// Opens a bam file and prints the name of each entry to standard output.
func main() {
	help := flag.Bool("help", false, "Print this usage message.")
	flag.Parse()

	bamFilename := flag.Arg(0)

	if *help {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(bamFilename)

	var err error

	bamIn, err := os.Open(bamFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.\n", err)
		os.Exit(1)
	}
	bamIn.Close()

	bamReader, err := bam.NewReader(bamIn, 1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.\n", err)
		os.Exit(1)
	}

	it, err := bam.NewIterator(bamReader, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.\n", err)
		os.Exit(1)
	}
	for it.Next() {
		rec := it.Record()
		fmt.Println(rec.Name)
	}

	bamReader.Close()
}
