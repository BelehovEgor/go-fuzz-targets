package main

// Edit if desired. Code generated by "fzgen --llm=groq -o=ascii85.go encoding/ascii85".
// false
// package alias
// unused variable

import (
	ascii85 "encoding/ascii85"
	"io"
	"reflect"
	"testing"

	"github.com/BelehovEgor/fzgen/fuzzer"
)

func Fuzz_N1_CorruptInputError_Error(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e ascii85.CorruptInputError
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&e)
		if err != nil {
			return
		}

		e.Error()
	})
}

func Fuzz_N2_Decode(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var dst []byte
		var src []byte
		var flush bool
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&dst, &src, &flush)
		if err != nil {
			return
		}

		ndst, nsrc, err := ascii85.Decode(dst, src, flush)
		if err != nil {
			if _, ok := err.(ascii85.CorruptInputError); ok {
				return
			}
			t.Error("Unexpected error")
		}

		if ndst < 0 || nsrc < 0 {
			t.Error("Negative result")
		}

		if nsrc > len(src) {
			t.Error("More bytes read than source")
		}

		if err == nil && (ndst == 0 || nsrc == 0) {
			t.Error("Zero result without error")
		}
	})
}

func Fuzz_N3_Encode(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var dst []byte
		var src []byte
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&dst, &src)
		if err != nil {
			return
		}

		if len(dst) < len(src)*5/4 {
			t.Skip("Destination buffer is too small")
		}

		if len(src) == 0 {
			if ascii85.Encode(dst, src) != 0 {
				t.Error("Encode with empty source should return 0")
			}
			return
		}

		n := ascii85.Encode(dst, src)
		if n <= 0 {
			t.Error("Encode should return positive number of bytes written")
		}

		if n > len(dst) {
			t.Error("Encode wrote more bytes than the destination buffer can hold")
		}
	})
}

func Fuzz_N4_MaxEncodedLen(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var n int
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&n)
		if err != nil {
			return
		}

		if n < 0 {
			t.Skip("Negative input skipped")
		}

		result := ascii85.MaxEncodedLen(n)
		if result < 0 {
			t.Error("Result should not be negative")
		}
	})
}

func Fuzz_N5_NewDecoder(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var r io.Reader
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&r)
		if err != nil {
			return
		}

		if r == nil {
			t.Skip("io.Reader is nil")
		}

		defer func() {
			if r := recover(); r != nil {
				if r == "explicit exception message" {
					t.Skip("Explicit exception")
				} else {
					t.Error("Unexpected panic")
				}
			}
		}()

		dec := ascii85.NewDecoder(r)
		if dec == nil {
			t.Error("NewDecoder returned nil")
		}
	})
}

func Fuzz_N6_NewEncoder(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var w io.Writer
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&w)
		if err != nil {
			return
		}

		if w == nil {
			t.Skip("Writer is nil")
		}

		defer func() {
			if r := recover(); r != nil {
				if r == "explicit exception message" {
					t.Skip("Explicit exception occurred")
				} else {
					t.Error("Panic occurred")
				}
			}
		}()

		encoder := ascii85.NewEncoder(w)
		if encoder == nil {
			t.Error("Encoder is nil")
		}
	})
}
func fabric_interface_empty_string(impl string) interface{} {
	return impl
}

var FabricFuncsForCustomTypes map[string][]reflect.Value

func TestMain(m *testing.M) {
	FabricFuncsForCustomTypes = make(map[string][]reflect.Value)
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_string))
	FabricFuncsForCustomTypes["io.WriteCloser"] = append(FabricFuncsForCustomTypes["io.WriteCloser"], reflect.ValueOf(ascii85.NewEncoder))
	FabricFuncsForCustomTypes["io.Reader"] = append(FabricFuncsForCustomTypes["io.Reader"], reflect.ValueOf(ascii85.NewDecoder))
	m.Run()
}
