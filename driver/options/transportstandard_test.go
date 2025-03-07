package options_test

import (
	"errors"
	"testing"

	"github.com/scrapli/scrapligo/driver/generic"
	"github.com/scrapli/scrapligo/driver/options"

	"github.com/google/go-cmp/cmp"
	"github.com/scrapli/scrapligo/transport"
	"github.com/scrapli/scrapligo/util"
)

func testStandardTransportExtraCiphers(
	testName string,
	testCase *optionsStringSliceTestCase,
) func(t *testing.T) {
	return func(t *testing.T) {
		t.Logf("%s: starting", testName)

		err := options.WithStandardTransportExtraCiphers(testCase.ss)(testCase.o)
		if err != nil {
			if errors.Is(err, util.ErrIgnoredOption) && !testCase.isignored {
				t.Fatalf(
					"%s: option should be ignored, but returned different error",
					testName,
				)
			}

			return
		}

		oo, _ := testCase.o.(*transport.Standard)

		if !cmp.Equal(oo.ExtraCiphers, testCase.ss) {
			t.Fatalf(
				"%s: actual and expected transport extra ciphers do not match\nactual:"+
					" %v\nexpected:%v",
				testName,
				oo.ExtraCiphers,
				testCase.ss,
			)
		}
	}
}

func TestStandardTransportExtraCiphers(t *testing.T) {
	cases := map[string]*optionsStringSliceTestCase{
		"set-system-transport-open-args": {
			description: "simple set option test",
			ss:          []string{"some", "extra", "ciphers"},
			o:           &transport.Standard{},
			isignored:   false,
		},
		"ignored": {
			description: "skipped due to ignored type",
			ss:          []string{"some", "extra", "ciphers"},
			o:           &generic.Driver{},
			isignored:   true,
		},
	}

	for testName, testCase := range cases {
		f := testStandardTransportExtraCiphers(testName, testCase)

		t.Run(testName, f)
	}
}

func testStandardTransportExtraKexs(
	testName string,
	testCase *optionsStringSliceTestCase,
) func(t *testing.T) {
	return func(t *testing.T) {
		t.Logf("%s: starting", testName)

		err := options.WithStandardTransportExtraKexs(testCase.ss)(testCase.o)
		if err != nil {
			if errors.Is(err, util.ErrIgnoredOption) && !testCase.isignored {
				t.Fatalf(
					"%s: option should be ignored, but returned different error",
					testName,
				)
			}

			return
		}

		oo, _ := testCase.o.(*transport.Standard)

		if !cmp.Equal(oo.ExtraKexs, testCase.ss) {
			t.Fatalf(
				"%s: actual and expected transport extra kexs do not match\nactual:"+
					" %v\nexpected:%v",
				testName,
				oo.ExtraKexs,
				testCase.ss,
			)
		}
	}
}

func TestStandardTransportExtraKexs(t *testing.T) {
	cases := map[string]*optionsStringSliceTestCase{
		"set-system-transport-open-args": {
			description: "simple set option test",
			ss:          []string{"some", "extra", "kexs"},
			o:           &transport.Standard{},
			isignored:   false,
		},
		"ignored": {
			description: "skipped due to ignored type",
			ss:          []string{"some", "extra", "kexs"},
			o:           &generic.Driver{},
			isignored:   true,
		},
	}

	for testName, testCase := range cases {
		f := testStandardTransportExtraKexs(testName, testCase)

		t.Run(testName, f)
	}
}
