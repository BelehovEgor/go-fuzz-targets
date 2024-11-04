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
		var s *youtube.Service
		var part []string
		var abusereport *youtube.AbuseReport
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&s, &part, &abusereport)
		if s == nil || abusereport == nil {
			return
		}

		r := youtube.NewAbuseReportsService(s)
		abuseReportsInsertCall := r.Insert(part, abusereport)

		var fields []googleapi.Field
		var str string
		fz.Fill(&fields, &str)
		if abuseReportsInsertCall == nil {
			return
		}

		_, err := abuseReportsInsertCall.Do(quotaUser(str))
		if err != nil {
			return
		}
	})
}
