package tests

// Edit if desired. Code generated by "fzgen github.com/kataras/go-serializer/jsonp".

import (
	"testing"

	"github.com/kataras/go-serializer/jsonp"
	"github.com/thepudds/fzgen/fuzzer"
)

// skipping Fuzz_Serializer_Serialize because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Config_Merge(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var cfg []jsonp.Config
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&cfg)

		c := jsonp.DefaultConfig()
		c.Merge(cfg)
	})
}

func Fuzz_Config_MergeSingle(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var cfg jsonp.Config
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&cfg)

		c := jsonp.DefaultConfig()
		c.MergeSingle(cfg)
	})
}

func Fuzz_New(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var cfg []jsonp.Config
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&cfg)

		jsonp.New(cfg...)
	})
}
