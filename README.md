# WaveGen

[![HeRCLab](https://circleci.com/gh/HeRCLab/wavegen.svg?style=svg)](https://app.circleci.com/pipelines/github/HeRCLab/wavegen?branch=master) [![HeRCLab](https://goreportcard.com/badge/github.com/HeRCLab/wavegen)](https://goreportcard.com/report/github.com/HeRCLab/wavegen) [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/herclab/wavegen)

This tool is used to generate composite waves as test data for one of our
projects. It allows the user to specify a collection of noise levels,
amplitudes, frequencies, and phases, and generates each requested sin wave and
noise, then sums them together for output in a standard format.

## Installation

### From Binary

Download the `.deb` or generic tarball binaries from the GitHub releases page
and install as appropriate.

### From Source

Run `go install ./cmd/wavegen/...`, this will install wavegen into your
`$GOPATH`.

If you would prefer a system-wide installation, run `make ; sudo make install`
(requires `help2man` and `ronn` installed).

### Release

To generate a release:

Pre-requisites:
* Golang version 1.13 or better
* help2man
* ronn
* checkinstall
* gnuplot

On Ubuntu 20.04: `sudo apt insttall help2man ronn golang-go checkinstall gnuplot`

Run `./build_release.sh`, you must have permission to run commands with `sudo`.
This will generate a generic binary tarball, as well as a Debian binary package
package.

**NOTE**: attempting to compile a release without having `gnuplot` install may
cause `checkinstall` to fail with inscrutable errors.

**NOTE**: it is rare one would want to do this manually, the CI will
automatically generate and upload binaries for tagged releases matching `x.y.z`
format.

## Usage

See `man wavegen`.

`gnuplot` must be installed.

## License

[LICENSE](./LICENSE)

