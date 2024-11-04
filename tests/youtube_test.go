package tests

import (
	"github.com/thepudds/fzgen/fuzzer"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
	"testing"
)

type quotaUser string

func (q quotaUser) Get() (string, string) { return "quotaUser", string(q) }

func Fuzz_AbuseReportsInsertCall_Do(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		var c *youtube.AbuseReportsInsertCall
		var s []googleapi.Field
		var str string
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&c, &s, &str)
		if c == nil {
			return
		}

		_, err := c.Do(quotaUser(str))
		if err != nil {
			return
		}
	})
}
