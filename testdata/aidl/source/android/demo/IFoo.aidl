package android.demo;

import android.other.IBar;

interface IFoo {
  void ping(in int value);
  IBar getBar();
}
