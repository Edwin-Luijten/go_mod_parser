package go_mod_parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	contents, _ := ioutil.ReadFile("./_testdata/go_mod.test")
	p, err := Parse("go.mod", contents, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	if p.Go.Version != "1.12" {
		t.Log("invalid version")
		t.Fail()
	}

	if p.Module.Mod.Path != "github.com/edwin-luijten/go_mod_parser" {
		t.Log("invalid path")
		t.Fail()
	}

	if p.Require[0].Mod.Path != "rsc.io/sampler" {
		t.Log("invalid require path")
		t.Fail()
	}

	if p.Require[0].Mod.Version != "v1.3.0" {
		t.Log("invalid require version")
		t.Fail()
	}
}
