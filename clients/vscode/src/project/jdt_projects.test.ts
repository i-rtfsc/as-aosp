import assert from "node:assert/strict";
import test from "node:test";
import { renderClasspath, renderProject } from "./jdt_projects";

test("renders distinct linked source folders", () => {
  const project = renderProject("aosp-services", ["/a/services/core/java", "/a/services/usb/java"]);
  assert.match(project, /<name>src-0000<\/name>/);
  assert.match(project, /<name>src-0001<\/name>/);
  assert.match(project, /file:\/\/\/a\/services\/core\/java/);
});

test("renders project and library classpath entries", () => {
  const classpath = renderClasspath(2, [], ["aosp-java-foundation", "aosp-generated"]);
  assert.match(classpath, /kind="src" path="src-0000"/);
  assert.match(classpath, /kind="src" path="\/aosp-java-foundation"/);
  assert.match(classpath, /kind="output" path="bin"/);
});
