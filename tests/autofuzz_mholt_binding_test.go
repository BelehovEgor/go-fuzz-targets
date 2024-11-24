package tests

// Edit if desired. Code generated by "fzgen github.com/mholt/binding".

import (
	"testing"

	"github.com/mholt/binding"
	"github.com/thepudds/fzgen/fuzzer"
)

func Fuzz_Errors_Add(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e *binding.Errors
		var fieldNames []string
		var classification string
		var message string
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&e, &fieldNames, &classification, &message)
		if e == nil {
			return
		}

		e.Add(fieldNames, classification, message)
	})
}

func Fuzz_Errors_Has(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e *binding.Errors
		var class string
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&e, &class)
		if e == nil {
			return
		}

		e.Has(class)
	})
}

func Fuzz_Errors_Len(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e *binding.Errors
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&e)
		if e == nil {
			return
		}

		e.Len()
	})
}

func Fuzz_Error_Error(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e binding.Error
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&e)

		e.Error()
	})
}

func Fuzz_Error_Fields(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e binding.Error
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&e)

		e.Fields()
	})
}

func Fuzz_Error_Kind(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e binding.Error
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&e)

		e.Kind()
	})
}

func Fuzz_Errors_Error(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e binding.Errors
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&e)

		e.Error()
	})
}

// skipping Fuzz_Errors_Handle because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Bind because parameters include func, chan, or unsupported interface: github.com/mholt/binding.FieldMapper

// skipping Fuzz_Form because parameters include func, chan, or unsupported interface: github.com/mholt/binding.FieldMapper

// skipping Fuzz_Json because parameters include func, chan, or unsupported interface: github.com/mholt/binding.FieldMapper

// skipping Fuzz_MultipartForm because parameters include func, chan, or unsupported interface: github.com/mholt/binding.FieldMapper

// skipping Fuzz_Validate because parameters include func, chan, or unsupported interface: github.com/mholt/binding.FieldMapper