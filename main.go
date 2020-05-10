package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asvvvad/clap/clap"
	"github.com/atotto/clipboard"
)

func main() {
	var file string
	var stdin bool
	var print bool

	flag.StringVar(&file, "file", "", "Use text from file")
	flag.BoolVar(&stdin, "stdin", false, "Use text from stdin")
	flag.BoolVar(&print, "print", false, "Print result.")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Println("\n  Copy the text you want to tilde then run tildeit and paste the result")
		fmt.Println("\n  Options:")
		fmt.Println("  -------")
		flag.PrintDefaults()
		fmt.Println("\n~t~i~l~d~e~i~t~ by tilde.club/~asvvvad - github.com/asvvvad/tildeit")
	}

	flag.Parse()

	// The string that will contain the input text
	var text string
	var err error
	var bytes []byte

	if stdin {
		bytes, err = ioutil.ReadAll(os.Stdin)
		text = string(bytes)
	} else if file != "" {
		bytes, err = ioutil.ReadFile(file)
		text = string(bytes)
	} else {
		text, err = clipboard.ReadAll()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// The string that will contain the final result
	result := "~" + clap.Emoji(text, "~") + "~"

	if print {
		fmt.Println(result)
	} else {
		err := clipboard.WriteAll(result)
		if err != nil {
			fmt.Println(err)
			// Print out in case of unsupported system
			fmt.Println(result)
			os.Exit(1)
		}
		fmt.Println("~tilde'd~and~copied~")
	}
}
