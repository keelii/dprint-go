package dprint_go

import (
	"testing"
)

func TestDprintFormat(t *testing.T) {
	if ret, _ := FormatText("test.ts", "var a=1", DprintConfig{}); ret != "var a = 1\n" {
		t.Error("FormatText failed", ret)
	}

	if ret, _ := FormatText("test.ts", "var a=1", DprintConfig{
		p: PluginConfig{SemiColons: "prefer"},
	}); ret != "var a = 1;\n" {
		t.Error("FormatText failed", ret)
	}

	if ret, _ := FormatText("test.ts", "var a='1'", DprintConfig{
		p: PluginConfig{QuoteStyle: "alwaysDouble"},
	}); ret != "var a = \"1\"\n" {
		t.Error("FormatText failed", ret)
	}

	if ret, _ := FormatText("test.ts", "if (true) {\n1}", DprintConfig{}); ret != "if (true) {\n  1\n}\n" {
		t.Error("FormatText failed", ret)
	}
	if ret, _ := FormatText("test.ts", "if (true) {\n1}", DprintConfig{
		g: GlobalConfiguration{
			IndentWidth: 4,
		},
	}); ret != "if (true) {\n    1\n}\n" {
		t.Error("FormatText failed", ret)
	}
	if ret, _ := FormatText("test.ts", "if (true) {\n1}", DprintConfig{
		g: GlobalConfiguration{
			UseTabs: true,
		},
	}); ret != "if (true) {\n\t1\n}\n" {
		t.Error("FormatText failed", ret)
	}
}
