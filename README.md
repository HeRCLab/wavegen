# WaveGen

[![HeRCLab](https://circleci.com/gh/HeRCLab/wavegen.svg?style=svg)](https://app.circleci.com/pipelines/github/HeRCLab/wavegen?branch=master)

[![HeRCLab](https://goreportcard.com/badge/github.com/HeRCLab/wavegen)(https://goreportcard.com/report/github.com/HeRCLab/wavegen)

This tool is used to generate composite waves as test data for one of our
projects. It allows the user to specify a collection of noise levels,
amplitudes, frequencies, and phases, and generates each requested sin wave and
noise, then sums them together for output in a standard format.

## Installation

### From Binary

Simply download the `.deb` package and install with dpkg.

### From Source

Run `go install ./cmd/wavegen/...`, this will install wavegen into your
`$GOPATH`.

If you would prefer a system-wide installation, run `make ; sudo make install`
(requires help2man and ronn installed).

### Release

To generate a release:

Pre-requisites:
* Golang version 1.13 or better
* help2man
* ronn
* checkinstall

On Ubuntu 20.04: `sudo apt insttall help2man ronn golang-go`

Run `./build_release.sh`, you must have permission to run commands with `sudo`.
This will generate a generic binary tarball, as well as a Debian source
package.

## Usage

See `man wavegen`.

## License

[LICENSE](./LICENSE)
