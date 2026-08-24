const crypto = require("node:crypto");
const fs = require("node:fs");
const path = require("node:path");

const root = path.resolve(process.argv[2]);
const artifacts = [];
for (const kind of ["engine", "jetbrains", "vscode"]) {
  const directory = path.join(root, kind);
  const stack = [directory];
  while (stack.length) {
    const current = stack.pop();
    for (const entry of fs.readdirSync(current, { withFileTypes: true })) {
      const absolute = path.join(current, entry.name);
      if (entry.isDirectory()) stack.push(absolute);
      else {
        const relative = path.relative(root, absolute).split(path.sep).join("/");
        const sha256 = crypto.createHash("sha256").update(fs.readFileSync(absolute)).digest("hex");
        const artifact = { path: relative, kind, sha256 };
        if (kind === "engine") {
          const target = relative.split("/")[1];
          const separator = target.lastIndexOf("-");
          artifact.os = target.slice(0, separator);
          artifact.arch = target.slice(separator + 1);
          artifact.executable = relative.endsWith(".exe") ? "aosp-index.exe" : "aosp-index";
        }
        artifacts.push(artifact);
      }
    }
  }
}
artifacts.sort((left, right) => left.path.localeCompare(right.path));
fs.writeFileSync(path.join(root, "manifest.json"), `${JSON.stringify({
  version: process.env.AOSP_IDE_VERSION || "0.1.0-dev",
  artifacts,
}, null, 2)}\n`);
