const path = require("node:path");
const { runTests } = require("@vscode/test-electron");

async function main() {
  await runTests({
    extensionDevelopmentPath: path.resolve(__dirname, ".."),
    extensionTestsPath: path.resolve(__dirname, "suite"),
    launchArgs: ["--disable-extensions"],
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
