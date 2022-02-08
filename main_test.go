package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type addTest struct {
	testCaseName   string
	ip             string
	expectedOutput string
}

var testCases = []addTest{
	{"testCase1", "128.0.0.1", "validIPv4\n"},
	{"testCase2", "125.512.100.abc", "invalidIP\n"},
	{"testCase3", "256.256.256.256", "invalidIP\n"},
	{"testCase4", "192.168.01.1", "invalidIP\n"},
	{"testCase5", "000.12.234.23.23", "invalidIP\n"},
	{"testCase6", "192.168.1.0", "validIPv4\n"},
	{"testCase7", "2001:db8:85a3:0:0:8A2E:0370:7334", "validIPv6\n"},
	{"testCase8", "2001:0db8:85a3:0:0:8A2E:0370:7334", "validIPv6\n"},
	{"testCase9", "2001:db8:85a3:0000:0000:8a2e:0370:7334", "validIPv6\n"},
	{"testCase10", "2001:0db8:85a3::8A2E:037j:7334", "invalidIP\n"},
	{"testCase11", "2F33:12a0:3Ea0:0302", "invalidIP\n"},
	{"testCase12", "I.Am.not.an.ip", "invalidIP\n"},
}

func TestIp(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.testCaseName, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/?ip="+test.ip, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			ip(rec, req)
			res := rec.Result()
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status OK, got %v", res.Status)
			}
			output, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}
			actualOutput := string(output)
			if actualOutput != test.expectedOutput {
				t.Fatalf("expected %v, got %v", test.expectedOutput, actualOutput)
			}
		})
	}
}
