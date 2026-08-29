export function findLocalJavaDeclarations(source: string, name: string): number[] {
  const escaped = name.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
  const patterns = [
    new RegExp(`\\b(?:class|interface|enum|record)\\s+(${escaped})\\b`, "g"),
    new RegExp(
      `(?:^|[;{}]\\s*|\\n\\s*)(?:@[\\w$.]+(?:\\([^)]*\\))?\\s*)*` +
      `(?:(?:public|protected|private|static|final|abstract|synchronized|native|default|strictfp)\\s+)*` +
      `(?:<[^>{}]+>\\s+)?(?:[\\w$.[\\]<>?,]+\\s+)+(${escaped})\\s*` +
      `\\([^;{}]*\\)\\s*(?:throws\\s+[^;{]+)?[;{]`,
      "gm",
    ),
    new RegExp(
      `(?:^|[;{}]\\s*|\\n\\s*)(?:@[\\w$.]+(?:\\([^)]*\\))?\\s*)*` +
      `(?:(?:public|protected|private|static|final|transient|volatile)\\s+)*` +
      `(?:[\\w$.[\\]<>?,]+\\s+)+(${escaped})\\s*(?:[=;,])`,
      "gm",
    ),
  ];
  const offsets = new Set<number>();
  for (const pattern of patterns) {
    for (const match of source.matchAll(pattern)) {
      const capture = match[1];
      if (capture && match.index !== undefined) {
        offsets.add(match.index + match[0].lastIndexOf(capture));
      }
    }
  }
  return [...offsets].sort((a, b) => a - b);
}

export function findLocalJavaConstructors(source: string, name: string): number[] {
  const escaped = name.replace(/[.*+?^${}()|[\]\\]/g, "\\$&");
  const pattern = new RegExp(
    `(?:^|[;{}]\\s*|\\n\\s*)(?:@[\\w$.]+(?:\\([^)]*\\))?\\s*)*` +
    `(?:(?:public|protected|private)\\s+)*(${escaped})\\s*` +
    `\\([^;{}]*\\)\\s*(?:throws\\s+[^;{]+)?\\{`,
    "gm",
  );
  const offsets = new Set<number>();
  for (const match of source.matchAll(pattern)) {
    if (match[1] && match.index !== undefined) {
      offsets.add(match.index + match[0].lastIndexOf(match[1]));
    }
  }
  return [...offsets].sort((a, b) => a - b);
}
