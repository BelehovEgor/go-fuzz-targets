package main

// Edit if desired. Code generated by "fzgen github.com/rs/zerolog".

import (
	"reflect"
	"testing"

	"github.com/BelehovEgor/fzgen/fuzzer"
	zerolog "github.com/rs/zerolog"
)

// #1

// can return nil of updated struct
// fuzz successful 216

func Fuzz_N51_Event_Floats64_default(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e *zerolog.Event
		var key string
		var f []float64
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&e, &key, &f)
		if err != nil || e == nil {
			return
		}

		// Put here your precondition of func arguments...

		e.Floats64(key, f)

		// Put here your postcondition of func results...
	})
}

// Prompt

/*
You write fuzzing test using golang

Purpose: extend target func result check in fuzzing target

Requirements:
        + return only fuzz target code
        + you can past code only instead comments like "Put here your"
        + no explanation
        + process all edge cases
        + if arguments is invalid, target function shouldn't be call, this case should be skipped
        + if there is an explicit exception creation, skip only them by their message, the rest should cause a fuzzing test error
        + situations that should not occur during the execution of the function should end with t.Error
        + use simple strings for t.Error or t.Log without any variables format
        + don't use not exported fields in validation checks (its starts with lowercase letter)

Target func from package github.com/rs/zerolog:

func (e *Event) Floats64(key string, f []float64) *Event {
        if e == nil {
                return e
        }
        e.buf = enc.AppendFloats64(enc.AppendKey(e.buf, key), f, FloatingPointPrecision)
        return e
}

Current fuzzing target test:

func Fuzz_N51_Event_Floats64(f *testing.F) {
        f.Fuzz(func(t *testing.T, data []byte) {
                var e *zerolog.Event
                var key string
                var f []float64
                fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
                err := fz.Fill2(&e, &key, &f)
                if err != nil || e == nil {
                        return
                }

                // Put here your precondition of func arguments...

                e.Floats64(key, f)

                // Put here your postcondition of func results...
        })
}

Current import aliases:

import (
        "context"
        "io"
        "net"
        "reflect"
        "testing"
        "time"

        "github.com/BelehovEgor/fzgen/fuzzer"
        zerolog "github.com/rs/zerolog"
)
*/

// #LLM claude-3-7-sonnet-20250219
// count fixes: 1
// trying use unexported
// #summary:
// invalid precondition.
// invalid post condition.

func Fuzz_N51_Event_Floats64_claude(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e *zerolog.Event
		var key string
		var f []float64
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&e, &key, &f)
		if err != nil || e == nil {
			return
		}

		// Save the original buffer to compare later
		// originalBuf := make([]byte, len(e.buf))
		// copy(originalBuf, e.buf)

		result := e.Floats64(key, f)

		// Check that the method returns the event itself
		if result != e {
			t.Error("Method should return the event itself")
		}

		// Check that the buffer has been modified
		// if len(e.buf) <= len(originalBuf) {
		// 	t.Error("Buffer should be extended after adding floats64")
		// }

		// Check that the key was added to the buffer
		// keyBytes := []byte(key)
		// found := false
		// for i := 0; i <= len(e.buf)-len(keyBytes); i++ {
		// 	if reflect.DeepEqual(e.buf[i:i+len(keyBytes)], keyBytes) {
		// 		found = true
		// 		break
		// 	}
		// }
		// if !found && len(key) > 0 {
		// 	t.Error("Key not found in the buffer")
		// }
	})
}

// #LLM claude-3-7-sonnet-20250219-thinking-32k
// count fixes: 0
// #summary:
// valid precondition.
// valid postcondition.
// fuzzing success 222
func Fuzz_N51_Event_Floats64_claude_thinking(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e *zerolog.Event
		var key string
		var f []float64
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&e, &key, &f)
		if err != nil || e == nil {
			return
		}

		defer func() {
			if r := recover(); r != nil {
				t.Error("Floats64 method should not panic")
			}
		}()

		result := e.Floats64(key, f)

		if result != e {
			t.Error("Method should return the same event instance")
		}
	})
}

// #LLM chatgpt-4o-latest-20250129
// count fixes: 2
// use unexsited
// #summary:
// valid precondition.
// invalid postcondition.

func Fuzz_N51_Event_Floats64_gpt(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var e *zerolog.Event
		var key string
		var f []float64
		fz := fuzzer.NewFuzzerV2(data, FabricFuncsForCustomTypes, t, fuzzer.Constructors)
		err := fz.Fill2(&e, &key, &f)
		if err != nil || e == nil {
			return
		}

		if key == "" {
			return
		}

		// unnecessary
		defer func() {
			if r := recover(); r != nil {
				t.Error("Panic occurred")
			}
		}()

		//beforeLen := len(e.Floats64(key, f).Buf)

		e.Floats64(key, f)

		// afterLen := len(e.Buf)
		// if afterLen <= beforeLen {
		// 	t.Error("Buffer length did not increase")
		// }
	})
}

// Fabrics

func fabric_interface_zerolog_Sampler_LevelSampler(impl zerolog.LevelSampler) zerolog.Sampler {
	return impl
}

func fabric_interface_zerolog_Sampler_BurstSampler(impl *zerolog.BurstSampler) zerolog.Sampler {
	return impl
}

func fabric_interface_zerolog_Sampler_BasicSampler(impl *zerolog.BasicSampler) zerolog.Sampler {
	return impl
}

func fabric_interface_zerolog_LogArrayMarshaler_Array(impl *zerolog.Array) zerolog.LogArrayMarshaler {
	return impl
}

func fabric_interface_zerolog_LevelWriter_TriggerLevelWriter(impl *zerolog.TriggerLevelWriter) zerolog.LevelWriter {
	return impl
}

func fabric_interface_zerolog_LevelWriter_LevelWriterAdapter(impl zerolog.LevelWriterAdapter) zerolog.LevelWriter {
	return impl
}

func fabric_interface_zerolog_LevelWriter_FilteredLevelWriter(impl *zerolog.FilteredLevelWriter) zerolog.LevelWriter {
	return impl
}

func fabric_interface_zerolog_Hook_LevelHook(impl zerolog.LevelHook) zerolog.Hook {
	return impl
}

func fabric_interface_empty_string(impl string) interface{} {
	return impl
}

func fabric_interface_empty_TriggerLevelWriter(impl zerolog.TriggerLevelWriter) interface{} {
	return impl
}

func fabric_interface_empty_TestWriter(impl zerolog.TestWriter) interface{} {
	return impl
}

func fabric_interface_empty_Logger(impl zerolog.Logger) interface{} {
	return impl
}

func fabric_interface_empty_LevelWriterAdapter(impl zerolog.LevelWriterAdapter) interface{} {
	return impl
}

func fabric_interface_empty_LevelSampler(impl zerolog.LevelSampler) interface{} {
	return impl
}

func fabric_interface_empty_LevelHook(impl zerolog.LevelHook) interface{} {
	return impl
}

func fabric_interface_empty_FilteredLevelWriter(impl zerolog.FilteredLevelWriter) interface{} {
	return impl
}

func fabric_interface_empty_Event(impl zerolog.Event) interface{} {
	return impl
}

func fabric_interface_empty_Context(impl zerolog.Context) interface{} {
	return impl
}

func fabric_interface_empty_ConsoleWriter(impl zerolog.ConsoleWriter) interface{} {
	return impl
}

func fabric_interface_empty_BurstSampler(impl zerolog.BurstSampler) interface{} {
	return impl
}

func fabric_interface_empty_BasicSampler(impl zerolog.BasicSampler) interface{} {
	return impl
}

func fabric_interface_empty_Array(impl zerolog.Array) interface{} {
	return impl
}

var FabricFuncsForCustomTypes map[string][]reflect.Value

func TestMain(m *testing.M) {
	FabricFuncsForCustomTypes = make(map[string][]reflect.Value)
	FabricFuncsForCustomTypes["zerolog.LevelWriter"] = append(FabricFuncsForCustomTypes["zerolog.LevelWriter"], reflect.ValueOf(fabric_interface_zerolog_LevelWriter_TriggerLevelWriter))
	FabricFuncsForCustomTypes["zerolog.LevelWriter"] = append(FabricFuncsForCustomTypes["zerolog.LevelWriter"], reflect.ValueOf(fabric_interface_zerolog_LevelWriter_LevelWriterAdapter))
	FabricFuncsForCustomTypes["zerolog.LevelWriter"] = append(FabricFuncsForCustomTypes["zerolog.LevelWriter"], reflect.ValueOf(fabric_interface_zerolog_LevelWriter_FilteredLevelWriter))
	FabricFuncsForCustomTypes["zerolog.Hook"] = append(FabricFuncsForCustomTypes["zerolog.Hook"], reflect.ValueOf(fabric_interface_zerolog_Hook_LevelHook))
	FabricFuncsForCustomTypes["zerolog.LogArrayMarshaler"] = append(FabricFuncsForCustomTypes["zerolog.LogArrayMarshaler"], reflect.ValueOf(fabric_interface_zerolog_LogArrayMarshaler_Array))
	FabricFuncsForCustomTypes["zerolog.Sampler"] = append(FabricFuncsForCustomTypes["zerolog.Sampler"], reflect.ValueOf(fabric_interface_zerolog_Sampler_LevelSampler))
	FabricFuncsForCustomTypes["zerolog.Sampler"] = append(FabricFuncsForCustomTypes["zerolog.Sampler"], reflect.ValueOf(fabric_interface_zerolog_Sampler_BurstSampler))
	FabricFuncsForCustomTypes["zerolog.Sampler"] = append(FabricFuncsForCustomTypes["zerolog.Sampler"], reflect.ValueOf(fabric_interface_zerolog_Sampler_BasicSampler))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_LevelHook))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_FilteredLevelWriter))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_ConsoleWriter))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_BurstSampler))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_LevelWriterAdapter))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_BasicSampler))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_Array))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_TriggerLevelWriter))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_TestWriter))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_Logger))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_LevelSampler))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_Event))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_Context))
	FabricFuncsForCustomTypes["interface {}"] = append(FabricFuncsForCustomTypes["interface {}"], reflect.ValueOf(fabric_interface_empty_string))
	FabricFuncsForCustomTypes["io.Writer"] = append(FabricFuncsForCustomTypes["io.Writer"], reflect.ValueOf(zerolog.SyncWriter))
	FabricFuncsForCustomTypes["zerolog.Level"] = append(FabricFuncsForCustomTypes["zerolog.Level"], reflect.ValueOf(zerolog.ParseLevel))
	FabricFuncsForCustomTypes["zerolog.Logger"] = append(FabricFuncsForCustomTypes["zerolog.Logger"], reflect.ValueOf(zerolog.Nop))
	FabricFuncsForCustomTypes["zerolog.LevelHook"] = append(FabricFuncsForCustomTypes["zerolog.LevelHook"], reflect.ValueOf(zerolog.NewLevelHook))
	FabricFuncsForCustomTypes["zerolog.Logger"] = append(FabricFuncsForCustomTypes["zerolog.Logger"], reflect.ValueOf(zerolog.New))
	FabricFuncsForCustomTypes["zerolog.LevelWriter"] = append(FabricFuncsForCustomTypes["zerolog.LevelWriter"], reflect.ValueOf(zerolog.MultiLevelWriter))
	FabricFuncsForCustomTypes["zerolog.Level"] = append(FabricFuncsForCustomTypes["zerolog.Level"], reflect.ValueOf(zerolog.GlobalLevel))
	FabricFuncsForCustomTypes["zerolog.Event"] = append(FabricFuncsForCustomTypes["zerolog.Event"], reflect.ValueOf(zerolog.Dict))
	FabricFuncsForCustomTypes["zerolog.Logger"] = append(FabricFuncsForCustomTypes["zerolog.Logger"], reflect.ValueOf(zerolog.Ctx))
	FabricFuncsForCustomTypes["zerolog.Array"] = append(FabricFuncsForCustomTypes["zerolog.Array"], reflect.ValueOf(zerolog.Arr))
	m.Run()
}
