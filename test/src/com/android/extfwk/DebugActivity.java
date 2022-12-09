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

package com.android.extfwk;

import android.app.Activity;
import android.content.Context;
import android.os.Bundle;
import android.util.Log;
import android.view.Display;
import android.view.Window;
import android.view.WindowManager;
import android.widget.Button;
import android.widget.LinearLayout;
import android.widget.ScrollView;
import android.widget.TextView;

import com.android.extfwk.test.HookTest;
import com.journeyOS.server.vrr.VRRManager;

import system.ext.utils.JosLog;


/**
 * @author anqi.huang@outlook.com
 */
public class DebugActivity extends Activity {
    private static final String TAG = DebugActivity.class.getSimpleName();
    private LinearLayout mLayout;
    private Context mContext;

    /**
     * {@inheritDoc}
     */
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        mContext = getApplicationContext();
        mLayout = new LinearLayout(this);
        mLayout.setOrientation(LinearLayout.VERTICAL);
        initView();
    }

    /**
     * {@inheritDoc}
     */
    @Override
    protected void onDestroy() {
        super.onDestroy();
    }

    private void initView() {
        TextView textView = new TextView(this);
        mLayout.addView(textView);

        Button button = new Button(this);
        button.setText("60 FPS");
        button.setOnClickListener(v -> {
            Log.d(TAG, "60 button click");
            setFps(60.01f);
        });
        mLayout.addView(button);

        button = new Button(this);
        button.setText("120 FPS");
        button.setOnClickListener(v -> {
            Log.d(TAG, "120 button click");
            setFps(120.49f);
        });
        mLayout.addView(button);

        button = new Button(this);
        button.setText("Test");
        button.setOnClickListener(v -> {
            Log.d(TAG, "test button click");
            HookTest.get().test(mContext);
        });
        mLayout.addView(button);

        ScrollView sv = new ScrollView(this);
        sv.addView(mLayout);
        setContentView(sv);
    }

    private void setFps(float fps) {
        //VrrDisplayModeDirector.getDefault().setRefreshRate(mContext, fps);
        Display.Mode[] modes = getDisplayModes();
        for (Display.Mode mode : modes) {
            JosLog.d(VRRManager.VRR_TAG, TAG, "setFps() mode id = [" + mode.getModeId() + "], refresh rate = [" + mode.getRefreshRate() + "]");
            if (Math.round(mode.getRefreshRate()) == Math.round(fps)) {
                setMode(this, mode);
                break;
            }
        }
    }

    private Display.Mode[] getDisplayModes() {
        Display primaryDisplay = getDisplay();
        Display.Mode[] modes = primaryDisplay.getSupportedModes();
        return modes;
    }

    private void setMode(Activity activity, Display.Mode mode) {
        Window window = activity.getWindow();
        WindowManager.LayoutParams params = window.getAttributes();
        params.preferredDisplayModeId = mode.getModeId();
        window.setAttributes(params); //通过该函数通知wms layout变化。
    }
}
