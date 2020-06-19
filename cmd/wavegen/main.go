package main

import (
	"fmt"
	"os"

	"github.com/herclab/wavegen/pkg/wavegen"

	"github.com/akamensky/argparse"

	"github.com/kingishb/go-gnuplot"
)

func main() {
	parser := argparse.NewParser("wavegen", "synthetic wave generation utility")

	versionFlag := parser.Flag("v", "version", &argparse.Options{Help: "Display version number and exit."})

	/****** generate sub-commands *****************************************/

	defaultAmplitudes := true
	defaultPhases := true
	defaultFrequencies := true
	defaultNoises := true
	defaultNoiseMagnitudes := true
	defaultSampleRate := true
	defaultOffset := true
	defaultDuration := true
	defaultGlobalNoise := true
	defaultGlobalNoiseMagnitude := true

	generateCmd := parser.NewCommand("generate", "generate synthetic data")

	generateAmplitudes := generateCmd.FloatList("a", "amplitudes",
		&argparse.Options{
			Help:    "List of floating point amplitudes.",
			Default: []float64{1},
			Validate: func(args []string) error {
				defaultAmplitudes = false
				return nil
			},
		})

	generatePhases := generateCmd.FloatList("p", "phases",
		&argparse.Options{
			Help:    "List of floating point phases.",
			Default: []float64{0},
			Validate: func(args []string) error {
				defaultPhases = false
				return nil
			},
		})

	generateFrequencies := generateCmd.FloatList("f", "frequencies",
		&argparse.Options{
			Help:    "List of floating point frequencies.",
			Default: []float64{1},
			Validate: func(args []string) error {
				defaultFrequencies = false
				return nil
			},
		})

	generateNoises := generateCmd.StringList("n", "noises",
		&argparse.Options{
			Help:    "List of string noise types.",
			Default: []string{},
			Validate: func(args []string) error {
				defaultNoises = false
				return nil
			},
		})

	generateNoiseMagnitudes := generateCmd.FloatList("m", "noisemagnitudes",
		&argparse.Options{
			Help:    "List of floating point noise magnitudes.",
			Default: []float64{},
			Validate: func(args []string) error {
				defaultNoiseMagnitudes = false
				return nil
			},
		})

	generateSampleRate := generateCmd.Float("s", "samplerate",
		&argparse.Options{
			Help:    "Floating point sample rate (Hz).",
			Default: float64(1000),
			Validate: func(args []string) error {
				defaultSampleRate = false
				return nil
			},
		})

	generateOffset := generateCmd.Float("O", "offset",
		&argparse.Options{
			Help:    "Floating point offset (seconds).",
			Default: float64(0),
			Validate: func(args []string) error {
				defaultOffset = false
				return nil
			},
		})

	generateDuration := generateCmd.Float("d", "duration",
		&argparse.Options{
			Help:    "Floating point duration (seconds).",
			Default: float64(1),
			Validate: func(args []string) error {
				defaultDuration = false
				return nil
			},
		})

	generateGlobalNoise := generateCmd.String("N", "globalnoise",
		&argparse.Options{
			Help:    "String global noise type.",
			Default: "none",
			Validate: func(args []string) error {
				defaultGlobalNoise = false
				return nil
			},
		})

	generateGlobalNoiseMagnitude := generateCmd.Float("M", "globalnoisemagnitude",
		&argparse.Options{
			Help:    "Floating point global noise magnitudes",
			Default: 0.0,
			Validate: func(args []string) error {
				defaultGlobalNoiseMagnitude = false
				return nil
			},
		})

	generateOutput := generateCmd.String("o", "output", &argparse.Options{Help: "Specify output file, or '-' for stdout.", Default: "-"})

	generateDisplay := generateCmd.Flag("D", "display", &argparse.Options{Help: "Also interactively display the generated data."})

	generateLoad := generateCmd.String("l", "load", &argparse.Options{Help: "Load an existing wavegen file and use it's parameters rather than the defaults. Any parameters specified on the CLI take precedence."})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprint(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	if *versionFlag {
		fmt.Printf("wavegen v0.0.1-git\n")
		os.Exit(0)
	}

	/****** generate sub-commands *****************************************/
	if generateCmd.Happened() {
		param := &wavegen.WaveParameters{
			SampleRate:           *generateSampleRate,
			Offset:               *generateOffset,
			Duration:             *generateDuration,
			Frequencies:          *generateFrequencies,
			Phases:               *generatePhases,
			Amplitudes:           *generateAmplitudes,
			Noises:               *generateNoises,
			NoiseMagnitudes:      *generateNoiseMagnitudes,
			GlobalNoise:          *generateGlobalNoise,
			GlobalNoiseMagnitude: *generateGlobalNoiseMagnitude,
		}

		if *generateLoad != "" {
			loaded, err := wavegen.ReadJSON(*generateLoad)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to load parameters file '%s': %v\n",
					*generateLoad, err)
				os.Exit(1)
			}

			if defaultSampleRate {
				param.SampleRate = loaded.Parameters.SampleRate
			}

			if defaultOffset {
				param.Offset = loaded.Parameters.Offset
			}

			if defaultDuration {
				param.Duration = loaded.Parameters.Duration
			}

			if defaultFrequencies {
				param.Frequencies = loaded.Parameters.Frequencies
			}

			if defaultPhases {
				param.Phases = loaded.Parameters.Phases
			}

			if defaultAmplitudes {
				param.Amplitudes = loaded.Parameters.Amplitudes
			}

			if defaultNoises {
				param.Noises = loaded.Parameters.Noises
			}

			if defaultNoiseMagnitudes {
				param.NoiseMagnitudes = loaded.Parameters.NoiseMagnitudes
			}

			if defaultGlobalNoise {
				param.GlobalNoise = loaded.Parameters.GlobalNoise
			}

			if defaultGlobalNoiseMagnitude {
				param.GlobalNoiseMagnitude = loaded.Parameters.GlobalNoiseMagnitude
			}

		}

		err := param.ValidateParameters()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid parameters specified: %v\n", err)
			os.Exit(1)
		}

		sig, err := param.GenerateSyntheticData()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while generating signal: %v\n", err)
			os.Exit(1)
		}
		wf := wavegen.WaveFile{
			Version:    0,
			Parameters: param,
			Signal:     sig,
		}

		if *generateDisplay {
			p, err := gnuplot.NewPlotter("", true, false)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating plot: %v\n", err)
				os.Exit(1)
			}

			err = p.Cmd("set terminal qt")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error setting terminal to 'qt': %v\n", err)
				os.Exit(1)
			}

			err = p.PlotXY(wf.Signal.T, wf.Signal.S, "generated signal")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error plotting signal : %v\n", err)
				os.Exit(1)
			}

			// Because `reread` won't work in this environment, we
			// simply pause for one year instead, which works just
			// as well. This ensures that the graph is still
			// intractable after we exit.
			p.Cmd("pause 31540000")

		}

		if *generateOutput == "-" {
			data, err := wf.ToJSON()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to generate JSON: %v\n", err)
				os.Exit(1)
			}
			fmt.Print(string(data))
			fmt.Print("")
		} else {
			err := wf.WriteJSON(*generateOutput)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to write output: %v\n", err)
				os.Exit(1)
			}
		}

	} else {
		err := fmt.Errorf("no command specified")
		fmt.Fprint(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}
}
