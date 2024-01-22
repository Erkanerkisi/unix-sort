# usort

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Description

An implementation of Unix sort in Go.

## Table of Contents

- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Getting Started

Instructions on how to get your project up and running.

### Prerequisites

Ensure you have Go installed.

### Installation

Step-by-step guide on how to install your project.

```bash
# Build the usort executable
go build -o usort
```

### Example commands

Choose a sort type:

* 0: quicksort
* 1: bubblesort
* 2: mergesort
* 3: randomsort

```bash
# Run usort with unique data and using randomsort
./usort -u -sort 3 words.txt

# Run usort with unique data, quicksort, and display the first 10 lines
./usort -u -sort 1 -head 10 words.txt
