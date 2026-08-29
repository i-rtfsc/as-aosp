import * as fs from "node:fs";
import * as path from "node:path";
import { pathToFileURL } from "node:url";

export type JdtTargetProject = {
  id: string;
  sourceRoots: string[];
  libraries: string[];
  dependencies: string[];
};

export type JdtProjection = {
  root: string;
  projects: Array<{ id: string; name: string; directory: string }>;
};

export function writeJdtProjects(
  workspaceRoot: string,
  targets: JdtTargetProject[],
  generatedRoot?: string,
): JdtProjection {
  const destination = path.join(workspaceRoot, ".aosp-ide", "java", "projects");
  const temporary = destination + `.tmp-${process.pid}`;
  fs.rmSync(temporary, { recursive: true, force: true });
  fs.mkdirSync(temporary, { recursive: true });

  const projectNames = new Map<string, string>();
  for (const target of targets) projectNames.set(target.id, projectName(target.id));
  if (generatedRoot) projectNames.set("generated", "aosp-generated");

  const projects: JdtProjection["projects"] = [];
  for (const target of targets) {
    const name = projectNames.get(target.id)!;
    const dependencies = target.dependencies
      .map((id) => projectNames.get(id))
      .filter((value): value is string => Boolean(value));
    if (generatedRoot) dependencies.push("aosp-generated");
    const directory = writeProject(temporary, name, target.sourceRoots, target.libraries, dependencies);
    projects.push({ id: target.id, name, directory });
  }
  if (generatedRoot) {
    const directory = writeProject(temporary, "aosp-generated", [generatedRoot], [], []);
    projects.push({ id: "generated", name: "aosp-generated", directory });
  }

  fs.rmSync(destination, { recursive: true, force: true });
  fs.renameSync(temporary, destination);
  return {
    root: destination,
    projects: projects.map((project) => ({
      ...project, directory: path.join(destination, path.basename(project.directory)),
    })),
  };
}

function writeProject(
  root: string,
  name: string,
  sourceRoots: string[],
  libraries: string[],
  dependencies: string[],
): string {
  const directory = path.join(root, name);
  fs.mkdirSync(path.join(directory, ".settings"), { recursive: true });
  const roots = unique(sourceRoots).filter((value) => fs.existsSync(value));
  fs.writeFileSync(path.join(directory, ".project"), renderProject(name, roots), "utf8");
  fs.writeFileSync(
    path.join(directory, ".classpath"),
    renderClasspath(roots.length, unique(libraries), unique(dependencies)),
    "utf8",
  );
  fs.writeFileSync(
    path.join(directory, ".settings", "org.eclipse.jdt.core.prefs"),
    "eclipse.preferences.version=1\n" +
      "org.eclipse.jdt.core.compiler.codegen.targetPlatform=17\n" +
      "org.eclipse.jdt.core.compiler.compliance=17\n" +
      "org.eclipse.jdt.core.compiler.source=17\n",
    "utf8",
  );
  return directory;
}

export function renderProject(name: string, roots: string[]): string {
  const links = roots.map((root, index) => `
    <link>
      <name>${sourceLink(index)}</name>
      <type>2</type>
      <locationURI>${xml(pathToFileURL(root).toString())}</locationURI>
    </link>`).join("");
  return `<?xml version="1.0" encoding="UTF-8"?>
<projectDescription>
  <name>${xml(name)}</name>
  <comment>AOSP Workspace generated project</comment>
  <projects/>
  <buildSpec>
    <buildCommand><name>org.eclipse.jdt.core.javabuilder</name><arguments/></buildCommand>
  </buildSpec>
  <natures><nature>org.eclipse.jdt.core.javanature</nature></natures>
  <linkedResources>${links}
  </linkedResources>
</projectDescription>
`;
}

export function renderClasspath(
  sourceRootCount: number,
  libraries: string[],
  dependencies: string[],
): string {
  const sources = Array.from({ length: sourceRootCount }, (_, index) =>
    `  <classpathentry kind="src" path="${sourceLink(index)}"/>`).join("\n");
  const projects = dependencies.map((dependency) =>
    `  <classpathentry combineaccessrules="false" kind="src" path="/${xml(dependency)}"/>`)
    .join("\n");
  const jars = libraries.filter((library) => fs.existsSync(library)).map((library) =>
    `  <classpathentry kind="lib" path="${xml(library)}"/>`).join("\n");
  return `<?xml version="1.0" encoding="UTF-8"?>
<classpath>
${sources}
${projects}
${jars}
  <classpathentry kind="con" path="org.eclipse.jdt.launching.JRE_CONTAINER/org.eclipse.jdt.internal.debug.ui.launcher.StandardVMType/JavaSE-17"/>
  <classpathentry kind="output" path="bin"/>
</classpath>
`;
}

function sourceLink(index: number): string {
  return `src-${String(index).padStart(4, "0")}`;
}

function projectName(id: string): string {
  const normalized = id.toLowerCase().replace(/[^a-z0-9._-]+/g, "-").replace(/^-+|-+$/g, "");
  return `aosp-${normalized || "java-target"}`;
}

function unique(values: string[]): string[] {
  return [...new Set(values.filter(Boolean))].sort();
}

function xml(value: string): string {
  return value.replaceAll("&", "&amp;").replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;").replaceAll('"', "&quot;").replaceAll("'", "&apos;");
}
