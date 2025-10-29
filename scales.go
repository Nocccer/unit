package unit

var scales = map[string]float64{
	"p": 1e-12,
	"n": 1e-9,
	"µ": 1e-6,
	"m": 1e-3,
	"":  1,
	"k": 1e-3,
	"M": 1e-6,
	"G": 1e-9,
	"T": 1e-12,
}

var times = map[string]float64{
	"ns": 1e-9,
	"µs": 1e-6,
	"ms": 1e-3,
	"s":  1,
	"m":  60,
	"h":  3600,
	"d":  86400,
}
