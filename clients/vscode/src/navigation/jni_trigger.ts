export function shouldTriggerJniDefinition(language: string, context: string): boolean {
  if (language === "java" || language === "kotlin") return /\bnative\b/.test(context);
  return /\b(?:JNIEXPORT|JNINativeMethod|RegisterNatives|registerNativeMethods|jniRegisterNativeMethods)\b/.test(context);
}
