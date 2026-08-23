package android.security;

class Scrypt {
  private static native byte[] nativeScrypt(
      byte[] password, byte[] salt, int n, int r, int p, int keyLen);
}
