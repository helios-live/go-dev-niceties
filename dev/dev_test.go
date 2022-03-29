package dev

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	er "github.com/rotisserie/eris"
)

var update = flag.Bool("update", false, "updates fixtures if true")

func TestTrace(t *testing.T) {
	b, err := ioutil.ReadFile("./test-fixtures/output1.golden")
	if err != nil && !*update {
		t.Skipf("Don't have file: %s", err)
	}

	err = er.New("error bad request")
	err = er.Wrapf(err, "this is a level 2 error '%v'", "oupss")
	err = er.Wrapf(err, "this is a level 3 error '%v'", "hopa")
	err = er.Wrapf(err, "this is a top level error '%v'", "hopa")

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Trace(err)

	w.Close()

	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if *update {
		ioutil.WriteFile("./test-fixtures/output1.golden", out, 0777)
		b = out
	}

	expected := string(b)
	got := string(out)

	if got[20:] != expected[20:] {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, out)
	}
}
