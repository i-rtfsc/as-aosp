package rustproject

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestNormalizeMapsSourceAndOutPaths(t *testing.T) {
	result := Normalize([]byte(`{
  "sysroot":"prebuilts/rust",
  "crates":[{
    "root_module":"system/lib.rs",
    "proc_macro_dylib_path":"out/soong/macro.so",
    "env":{"OUT_DIR":"out/soong/generated"},
    "source":{"include_dirs":["system"],"exclude_dirs":[]}
  }]
}`), "/workspace/aosp", "/workspace/out", nil)
	if len(result.Diagnostics) != 0 || result.Crates != 1 {
		t.Fatalf("result=%+v", result)
	}
	var project map[string]any
	if err := json.Unmarshal(result.Data, &project); err != nil {
		t.Fatal(err)
	}
	crate := project["crates"].([]any)[0].(map[string]any)
	if crate["root_module"] != filepath.Join("/workspace/aosp", "system/lib.rs") ||
		crate["proc_macro_dylib_path"] != filepath.Join("/workspace/out", "soong/macro.so") {
		t.Fatalf("crate=%+v", crate)
	}
}
