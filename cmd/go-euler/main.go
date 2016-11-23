package main

//go:generate ./index

import (
	"fmt"
	"os"
	"runtime/pprof"

	flag "github.com/spf13/pflag"

	"github.com/phensley/go-euler"
)

var (
	debug   = flag.BoolP("debug", "d", false, "debug mode")
	verbose = flag.BoolP("verbose", "v", false, "verbose mode")
	reveal  = flag.BoolP("reveal", "r", false, "reveal the answer")
	profile = flag.StringP("profile", "p", "", "enable cpu profiling")
	elapsed = flag.Float64P("elapsed", "e", 0, "show answers that took longer than N.NNN seconds to produce")
)

func main() {
	flag.Parse()

	euler.Debug = *debug
	euler.Verbose = *verbose

	if *profile != "" {
		file, err := os.Create(*profile)
		euler.FatalOnError(err, "os.Create")
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	if euler.Debug {
		euler.ShowSolutions()
		fmt.Println()
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Printf("Usage: %s {all | ###}\n", os.Args[0])
		os.Exit(1)
	}

	all := false
	for _, arg := range args {
		if arg == "all" {
			all = true
			break
		}
	}
	if all {
		for ctx := range euler.SolveAll() {
			if ctx.Elapsed().Seconds() >= *elapsed {
				showAnswer(ctx)
			}
		}
	} else {
		for _, arg := range args {
			ctx := euler.Solve(arg)
			if ctx.Elapsed().Seconds() >= *elapsed {
				showAnswer(ctx)
			}
		}
	}
}

func showAnswer(ctx *euler.Context) {
	if euler.Debug {
		fmt.Println(ctx)
		return
	}
	if !ctx.Exists() {
		fmt.Printf("%s no solution\n", ctx.ID())
		return
	}

	m := "unfinished"
	if ctx.IsSolved() {
		m = "    ok    "
	} else if ctx.IsAnswered() {
		m = "   wrong  "
	}

	fmt.Printf("%s  %s  %.3f  %#v", ctx.ID(), m, ctx.Elapsed().Seconds(), ctx.Description())
	if *reveal && ctx.IsAnswered() {
		fmt.Print(": ", ctx.Answer())
	}
	fmt.Println()
}
