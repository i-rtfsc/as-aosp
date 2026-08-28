package rustproject

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/pathmap"
)

type Result struct {
	Data        []byte
	Crates      int
	Diagnostics []model.Diagnostic
}

func Normalize(data []byte, workspaceRoot, outRoot string, mappings []model.PathMapping) Result {
	result := Result{Diagnostics: []model.Diagnostic{}}
	var project map[string]any
	if err := json.Unmarshal(data, &project); err != nil {
		result.Diagnostics = append(result.Diagnostics, diagnostic(
			"RUST_PROJECT_PARSE_FAILED", err.Error(), "",
		))
		return result
	}
	rewrite := func(value any, field string) any {
		text, ok := value.(string)
		if !ok || strings.TrimSpace(text) == "" {
			return value
		}
		mapped, ok := mapPath(text, workspaceRoot, outRoot, mappings)
		if !ok {
			result.Diagnostics = append(result.Diagnostics, diagnostic(
				"RUST_PROJECT_PATH_UNMAPPED", "cannot map "+field, text,
			))
			return value
		}
		return mapped
	}
	project["sysroot"] = rewrite(project["sysroot"], "sysroot")
	crates, _ := project["crates"].([]any)
	result.Crates = len(crates)
	for _, raw := range crates {
		crate, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		crate["root_module"] = rewrite(crate["root_module"], "root_module")
		crate["proc_macro_dylib_path"] = rewrite(
			crate["proc_macro_dylib_path"], "proc_macro_dylib_path",
		)
		if env, ok := crate["env"].(map[string]any); ok {
			env["OUT_DIR"] = rewrite(env["OUT_DIR"], "env.OUT_DIR")
		}
		if source, ok := crate["source"].(map[string]any); ok {
			for _, key := range []string{"include_dirs", "exclude_dirs"} {
				values, _ := source[key].([]any)
				for index := range values {
					values[index] = rewrite(values[index], "source."+key)
				}
				source[key] = values
			}
		}
	}
	normalized, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		result.Diagnostics = append(result.Diagnostics, diagnostic(
			"RUST_PROJECT_WRITE_FAILED", err.Error(), "",
		))
		return result
	}
	result.Data = append(normalized, '\n')
	return result
}

func mapPath(value, workspaceRoot, outRoot string, mappings []model.PathMapping) (string, bool) {
	clean := strings.TrimSpace(strings.ReplaceAll(value, "\\", "/"))
	if clean == "" || strings.Contains(clean, "../") || strings.HasPrefix(clean, "../") {
		return "", false
	}
	if isAbsolute(clean) {
		candidate := filepath.Clean(filepath.FromSlash(clean))
		if under(candidate, workspaceRoot) || under(candidate, outRoot) {
			return candidate, true
		}
		if mapped, ok := (pathmap.Mapper{Mappings: mappings}).ExternalToLocal(
			model.ExternalPath(clean),
		); ok {
			return string(mapped), true
		}
		return "", false
	}
	if clean == "out" || strings.HasPrefix(clean, "out/") {
		if strings.TrimSpace(outRoot) == "" {
			return "", false
		}
		if clean == "out" {
			return filepath.Clean(outRoot), true
		}
		return filepath.Join(outRoot, filepath.FromSlash(strings.TrimPrefix(clean, "out/"))), true
	}
	if strings.TrimSpace(workspaceRoot) == "" {
		return "", false
	}
	return filepath.Join(workspaceRoot, filepath.FromSlash(clean)), true
}

func isAbsolute(value string) bool {
	return filepath.IsAbs(value) || strings.HasPrefix(value, "//") ||
		(len(value) >= 2 && value[1] == ':')
}

func under(pathName, root string) bool {
	if strings.TrimSpace(root) == "" {
		return false
	}
	relative, err := filepath.Rel(filepath.Clean(root), filepath.Clean(pathName))
	return err == nil && relative != ".." &&
		!strings.HasPrefix(relative, ".."+string(filepath.Separator))
}

func diagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{
		Code: code, Severity: model.SeverityWarning, Stage: "rust-project",
		Message: message, LocalPath: pathName,
	}
}
