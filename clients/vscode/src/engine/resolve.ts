import * as path from "node:path";

export function platformTarget(platform: string = process.platform, architecture: string = process.arch): string {
  const normalizedPlatform = platform === "win32" ? "windows" : platform;
  const normalizedArch = architecture === "x64" ? "amd64" : architecture === "arm64" ? "arm64" : architecture;
  return `${normalizedPlatform}-${normalizedArch}`;
}

export function resolveEngineExecutable(options: {
  configured?: string;
  extensionPath: string;
  exists?: (value: string) => boolean;
  platform?: string;
  architecture?: string;
}): string {
  const exists = options.exists ?? ((value) => require("node:fs").existsSync(value));
  if (options.configured?.trim()) return options.configured.trim();
  const target = platformTarget(options.platform, options.architecture);
  const name = target.startsWith("windows") ? "aosp-index.exe" : "aosp-index";
  const bundled = path.join(options.extensionPath, "engine", target, name);
  if (exists(bundled)) return bundled;
  return name;
}
