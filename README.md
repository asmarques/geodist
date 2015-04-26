# geodist
[![Build Status](https://travis-ci.org/asmarques/geodist.svg)](https://travis-ci.org/asmarques/geodist)
[![GoDoc](https://godoc.org/github.com/asmarques/geodist?status.svg)](https://godoc.org/github.com/asmarques/geodist)

Implementation of the [Haversine](http://en.wikipedia.org/wiki/Haversine_formula) and [Vincenty](http://en.wikipedia.org/wiki/Vincenty%27s_formulae) methods in Go for calculating geographical distances between points on the surface of the Earth.

## Installation

```bash
go get github.com/asmarques/geodist
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/asmarques/geodist"
)

func main() {
	lis := geodist.Point{38.781311, -9.135918}
	sfo := geodist.Point{37.618817, -122.375427}

	d := geodist.HaversineDistance(lis, sfo)
	fmt.Printf("Haversine: %.2f km\n", d)

	d, err := geodist.VincentyDistance(lis, sfo)
	if err != nil {
		fmt.Printf("Failed to converge: %v", err)
	}

	fmt.Printf("Vincenty: %.6f km\n", d)
}
```

## License

[MIT](LICENSE)