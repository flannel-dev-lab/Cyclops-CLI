# Cyclops CLI

[![Go Report Card](https://goreportcard.com/badge/github.com/flannel-dev-lab/Cyclops-CLI)](https://goreportcard.com/report/github.com/flannel-dev-lab/Cyclops-CLI)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/13ffc7db08264bb093c03d7e06263db4)](https://www.codacy.com/manual/vmanikes/Cyclops-CLI?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=flannel-dev-lab/Cyclops-CLI&amp;utm_campaign=Badge_Grade)

Helps in building a starter template to getting started with [cyclops](https://github.com/flannel-dev-lab/cyclops)

## Features
    - Creates the base directories for config, database and router

## How to run
    - Download the last `bin.zip` from [here](https://github.com/flannel-dev-lab/Cyclops-CLI/releases). 
    - Extract the zip file and choose the binary with respect to your operating system
    - Once downloaded, Run `./cyclops-cli -bootstrap {PATH TO PROJECT}`

## Want to build your self
    - Clone the repo
    - Run `GOOS={OS} GOARCH={ARCH} go build -o cyclops-cli CyclopsCLI.go`