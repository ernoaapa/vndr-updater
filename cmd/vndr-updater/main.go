package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ernoaapa/vndr-updater/pkg/vndr"
)

var config string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s <source file>", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&config, "config", "vendor.conf", "target config file")
}

func validateArgs() {
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(2)
	}
}

func main() {
	flag.Parse()
	validateArgs()

	targetDeps, err := vndr.ReadConfig(config)
	if err != nil {
		log.Printf("Failed to read target config: %a", err)
		os.Exit(3)
	}

	sourceDeps, err := vndr.ReadConfig((flag.Arg(0)))
	if err != nil {
		log.Printf("Failed to read source config: %a", err)
		os.Exit(4)
	}

	result := []vndr.DepEntry{}
	for _, targetDep := range targetDeps {
		sourceDep := findDep(sourceDeps, targetDep.ImportPath)
		if sourceDep != nil {
			result = append(result, *sourceDep)
		} else {
			result = append(result, targetDep)
		}
	}

	if err := vndr.WriteConfig(result, config); err != nil {
		log.Printf("Failed to read source config: %a", err)
		os.Exit(4)
	}
}

func findDep(deps []vndr.DepEntry, importPath string) *vndr.DepEntry {
	for _, dep := range deps {
		if dep.ImportPath == importPath {
			return &dep
		}
	}
	return nil
}
