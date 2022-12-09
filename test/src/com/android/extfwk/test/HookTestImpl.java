/*
 * Copyright (c) 2022 anqi.huang@outlook.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.android.extfwk.test;

import android.content.Context;
import android.os.RemoteException;

import com.journeyOS.server.godeye.GodEyeManager;
import com.journeyOS.server.godeye.GodEyeObserver;
import com.journeyOS.server.godeye.Scene;
import com.journeyOS.server.vrr.VRRManager;
import com.journeyOS.server.vrr.VrrInputMonitor;
import com.journeyOS.server.vrr.VrrSurfaceControlProxy;
import com.journeyOS.server.vrr.VrrThread;

import system.ext.utils.JosLog;

public class HookTestImpl implements HookTest {
    private static final String TAG = HookTestImpl.class.getSimpleName();

    public HookTestImpl() {
    }

    @Override
    public void test(Context context) {
        JosLog.d(VRRManager.VRR_TAG, TAG, "test() called with: context = [" + context + "]");

//        handler(context);
    }

    private void handler(Context context) {
        VrrInputMonitor.getDefault().startTouchMonitoring(context, VrrThread.getDefault().getLooper());
        VrrInputMonitor.getDefault().registerGestureEventListener(new VrrInputMonitor.OnGestureEvent() {
            @Override
            public void onTouch() {
                JosLog.d(VRRManager.VRR_TAG, TAG, "onTouch() called");
            }

            @Override
            public void onSpeed(float xVelocity, float yVelocity) {
                JosLog.d(VRRManager.VRR_TAG, TAG, "onSpeed() called with: xVelocity = [" + xVelocity + "], yVelocity = [" + yVelocity + "]");
            }
        });

        VrrSurfaceControlProxy.getDefault().getActiveDisplayModeId();

        long factors = GodEyeManager.SCENE_FACTOR_APP
                | GodEyeManager.SCENE_FACTOR_CAMERA
                | GodEyeManager.SCENE_FACTOR_VIDEO
                | GodEyeManager.SCENE_FACTOR_AUDIO
                | GodEyeManager.SCENE_FACTOR_BRIGHTNESS
                | GodEyeManager.SCENE_FACTOR_TEMPERATURE;

        GodEyeManager.getDefault().subscribeObserver(new GodEyeObserver() {
            @Override
            public void onSceneChanged(Scene scene) throws RemoteException {
                JosLog.d(GodEyeManager.GOD_EYE_TAG, TAG, "onSceneChanged() called with: scene = [" + scene + "]");
            }
        });
        GodEyeManager.getDefault().setFactor(factors);

//        MonitorManager.getInstance().init(context);
//        MonitorManager.getInstance().start(factors);
    }

}
