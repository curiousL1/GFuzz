package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {

	// Fuzzing Target

	GoModDir     string   `long:"goModDir" description:"Directory contains go.mod"`
	TestFuncs    []string `long:"testFunc" description:"Only run specific test functions in the tests"`
	TestBinGlobs []string `long:"testBin" description:"A list of globs for Go test bins."`

	// Fuzzer

	OutputDir string `long:"outputDir" description:"Directory for fuzzing output"`
	Parallel  int    `long:"parallel" description:"Number of workers to fuzz parallel" default:"5"`
	InstStats string `long:"instStats" description:"This parameter consumes a file path to a statistics file generated by isnt."`
	Version   bool   `long:"version" description:"Print version and exit"`

	// Fuzzing

	GlobalTuple  bool `long:"globalTuple" description:"Whether prev_location is global or per channel"`
	ScoreSdk     bool `long:"scoreSdk" description:"Recording/scoring if channel comes from Go SDK"`
	ScoreAllPrim bool `long:"scoreAllPrim" description:"Recording/scoring other primitives like Mutex together with channel"`
	TimeDivideBy int  `long:"timeDivideBy" description:"Durations in time/sleep.go will be divided by this int number"`
}

func parseFlags() {

	if _, err := flags.Parse(&opts); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

}
