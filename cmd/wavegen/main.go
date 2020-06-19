package main

import (
	"fmt"
	"os"

	"github.com/herclab/wavegen/pkg/wavegen"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("wavegen", "synthetic wave generation utility")

	versionFlag := parser.Flag("v", "version", &argparse.Options{Help: "Display version number and exit."})

	/****** generate sub-commands *****************************************/

	generateCmd := parser.NewCommand("generate", "generate synthetic data")

	generateAmplitudes := generateCmd.FloatList("a", "amplitudes", &argparse.Options{Help: "List of floating point amplitudes.", Default: []float64{1}})

	generatePhases := generateCmd.FloatList("p", "phases", &argparse.Options{Help: "List of floating point phases.", Default: []float64{0}})

	generateFrequencies := generateCmd.FloatList("f", "frequencies", &argparse.Options{Help: "List of floating point frequencies.", Default: []float64{1}})

	generateNoises := generateCmd.StringList("n", "noises", &argparse.Options{Help: "List of string noise types.", Default: []string{}})

	generateNoiseMagnitudes := generateCmd.FloatList("m", "noisemagnitudes", &argparse.Options{Help: "List of floating point noise magnitudes.", Default: []float64{}})

	generateSampleRate := generateCmd.Float("s", "samplerate", &argparse.Options{Help: "Floating point sample rate (Hz).", Default: float64(1000)})

	generateOffset := generateCmd.Float("O", "offset", &argparse.Options{Help: "Floating point offset (seconds).", Default: float64(0)})

	generateDuration := generateCmd.Float("d", "duration", &argparse.Options{Help: "Floating point duration (seconds).", Default: float64(1)})

	generateGlobalNoise := generateCmd.String("N", "globalnoise", &argparse.Options{Help: "String global noise type.", Default: "none"})

	generateGlobalNoiseMagnitude := generateCmd.Float("M", "globalnoisemagnitude", &argparse.Options{Help: "Floating point global noise magnitudes", Default: 0.0})

	generateOutput := generateCmd.String("o", "output", &argparse.Options{Help: "Specify output file, or '-' for stdout.", Default: "-"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Fprint(os.Stderr, parser.Usage(err))
		os.Exit(1)
	}

	if *versionFlag {
		fmt.Printf("wavegen v0.0.1-git")
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
