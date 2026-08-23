static const JNINativeMethod sMethods[] = {
  {"nativeScrypt", "([B[BIIII)[B", (void*)nativeScrypt},
};

int registerMethods(JNIEnv* env) {
  return jniRegisterNativeMethods(
      env, "android/security/Scrypt", sMethods, 1);
}
