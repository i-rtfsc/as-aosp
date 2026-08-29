package sourcecatalog

import (
	"strings"

	"github.com/i-rtfsc/as-aosp/core/modules"
)

func classifySourceModule(module Module) Module {
	module.BuildRole = roleForType(module.Type)
	module.OutputKinds = outputKindsForType(module.Type)
	module.DefaultVisible = isDefaultVisible(module.BuildRole, module.Capabilities)
	return module
}

func classifyBuildModule(module Module, facts modules.Module) Module {
	if module.BuildRole == "" {
		module = classifySourceModule(module)
	}
	if module.Type == "" && len(facts.Types) > 0 {
		module.Type = facts.Types[0]
	}
	module.BuildRole = mergeRole(module.BuildRole, roleForTypes(facts.Types))
	module.OutputKinds = uniqueStrings(append(module.OutputKinds, outputKindsForFacts(facts)...))
	module.DefaultVisible = module.DefaultVisible ||
		isDefaultVisible(module.BuildRole, module.Capabilities)
	return module
}

func roleForTypes(types []string) BuildRole {
	role := BuildRole("")
	for _, value := range types {
		role = mergeRole(role, roleForType(value))
	}
	return role
}

func roleForType(value string) BuildRole {
	lower := strings.ToLower(strings.TrimSpace(value))
	if lower == "" {
		return RoleBuildSupport
	}
	if strings.Contains(lower, "test") {
		return RoleTest
	}
	if strings.Contains(lower, "filegroup") || strings.Contains(lower, "genrule") ||
		strings.Contains(lower, "se_build") || strings.Contains(lower, "license") ||
		strings.Contains(lower, "config") || strings.Contains(lower, "prebuilt") {
		return RoleBuildSupport
	}
	if strings.Contains(lower, "app") || strings.Contains(lower, "binary") ||
		strings.Contains(lower, "executable") {
		return RoleProduct
	}
	if strings.Contains(lower, "library") || strings.Contains(lower, "lib") ||
		strings.Contains(lower, "rust") {
		return RoleLibrary
	}
	return RoleBuildSupport
}

func mergeRole(current, candidate BuildRole) BuildRole {
	if current == "" {
		return candidate
	}
	if candidate == "" || current == candidate {
		return current
	}
	if current == RoleTest || candidate == RoleTest {
		return RoleTest
	}
	if current == RoleProduct || candidate == RoleProduct {
		return RoleProduct
	}
	if current == RoleLibrary || candidate == RoleLibrary {
		return RoleLibrary
	}
	return RoleBuildSupport
}

func outputKindsForType(value string) []string {
	lower := strings.ToLower(value)
	switch {
	case strings.Contains(lower, "app"):
		return []string{"apk"}
	case strings.Contains(lower, "shared_library") || strings.Contains(lower, "library_shared"):
		return []string{"so"}
	case strings.Contains(lower, "static_library") || strings.Contains(lower, "library_static"):
		return []string{"a"}
	case strings.Contains(lower, "binary") || strings.Contains(lower, "executable"):
		return []string{"bin"}
	case strings.Contains(lower, "java") || strings.Contains(lower, "kotlin"):
		return []string{"jar"}
	case strings.Contains(lower, "filegroup") || strings.Contains(lower, "genrule") ||
		strings.Contains(lower, "se_build"):
		return []string{"generated"}
	default:
		return nil
	}
}

func outputKindsForFacts(facts modules.Module) []string {
	result := outputKindsForType(strings.Join(facts.Types, " "))
	if len(facts.Jars) > 0 {
		result = append(result, "jar")
	}
	if len(facts.SrcJars) > 0 {
		result = append(result, "srcjar")
	}
	if len(facts.Generated) > 0 {
		result = append(result, "generated")
	}
	return uniqueStrings(result)
}

func isDefaultVisible(role BuildRole, capabilities []string) bool {
	if role == RoleTest || role == RoleBuildSupport {
		return false
	}
	return len(capabilities) > 0 || role == RoleProduct || role == RoleLibrary
}
