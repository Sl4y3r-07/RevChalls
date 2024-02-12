package com.backdoor.sl4ydroid;

import android.os.Bundle;
import android.widget.TextView;
import androidx.appcompat.app.AppCompatActivity;

/* loaded from: classes3.dex */
public class MainActivity extends AppCompatActivity {
    private TextView displayTextView;
    private TextView textView;

    public native void damn(String str);

    public native void k2(String str);

    public native void kim(String str);

    public native void nim(String str);

    static {
        System.loadLibrary("sl4ydroid");
    }

    public void sh4dy(final String message) {
        System.out.println(message);
        runOnUiThread(new Runnable() { // from class: com.backdoor.sl4ydroid.MainActivity$$ExternalSyntheticLambda1
            @Override // java.lang.Runnable
            public final void run() {
                MainActivity.this.m51lambda$sh4dy$0$combackdoorsl4ydroidMainActivity(message);
            }
        });
    }

    /* JADX INFO: Access modifiers changed from: package-private */
    /* renamed from: lambda$sh4dy$0$com-backdoor-sl4ydroid-MainActivity  reason: not valid java name */
    public /* synthetic */ void m51lambda$sh4dy$0$combackdoorsl4ydroidMainActivity(String message) {
        String currentText = this.displayTextView.getText().toString();
        String newText = currentText + message;
        this.displayTextView.setText(newText);
    }

    public void sl4y3r(final String message) {
        System.out.println(message);
        runOnUiThread(new Runnable() { // from class: com.backdoor.sl4ydroid.MainActivity$$ExternalSyntheticLambda2
            @Override // java.lang.Runnable
            public final void run() {
                MainActivity.this.m52lambda$sl4y3r$1$combackdoorsl4ydroidMainActivity(message);
            }
        });
    }

    /* JADX INFO: Access modifiers changed from: package-private */
    /* renamed from: lambda$sl4y3r$1$com-backdoor-sl4ydroid-MainActivity  reason: not valid java name */
    public /* synthetic */ void m52lambda$sl4y3r$1$combackdoorsl4ydroidMainActivity(String message) {
        String currentText = this.displayTextView.getText().toString();
        String newText = currentText + message;
        this.displayTextView.setText(newText);
    }

    public void it4chi(final String message) {
        System.out.println(message);
        runOnUiThread(new Runnable() { // from class: com.backdoor.sl4ydroid.MainActivity$$ExternalSyntheticLambda3
            @Override // java.lang.Runnable
            public final void run() {
                MainActivity.this.m49lambda$it4chi$2$combackdoorsl4ydroidMainActivity(message);
            }
        });
    }

    /* JADX INFO: Access modifiers changed from: package-private */
    /* renamed from: lambda$it4chi$2$com-backdoor-sl4ydroid-MainActivity  reason: not valid java name */
    public /* synthetic */ void m49lambda$it4chi$2$combackdoorsl4ydroidMainActivity(String message) {
        String currentText = this.displayTextView.getText().toString();
        String newText = currentText + message;
        this.displayTextView.setText(newText);
    }

    public void n4ut1lus(final String message) {
        System.out.println(message);
        runOnUiThread(new Runnable() { // from class: com.backdoor.sl4ydroid.MainActivity$$ExternalSyntheticLambda0
            @Override // java.lang.Runnable
            public final void run() {
                MainActivity.this.m50lambda$n4ut1lus$3$combackdoorsl4ydroidMainActivity(message);
            }
        });
    }

    /* JADX INFO: Access modifiers changed from: package-private */
    /* renamed from: lambda$n4ut1lus$3$com-backdoor-sl4ydroid-MainActivity  reason: not valid java name */
    public /* synthetic */ void m50lambda$n4ut1lus$3$combackdoorsl4ydroidMainActivity(String message) {
        String currentText = this.displayTextView.getText().toString();
        String newText = currentText + message;
        this.displayTextView.setText(newText);
    }

    /* JADX INFO: Access modifiers changed from: protected */
    @Override // androidx.fragment.app.FragmentActivity, androidx.activity.ComponentActivity, androidx.core.app.ComponentActivity, android.app.Activity
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        TextView textView = (TextView) findViewById(R.id.displayTextView);
        this.displayTextView = textView;
        textView.setVisibility(8);
        TextView textView2 = (TextView) findViewById(R.id.textView);
        this.textView = textView2;
        textView2.setText(getResources().getString(R.string.message));
        kim(getResources().getString(R.string.k1));
        nim(getResources().getString(R.string.n1));
        damn(getResources().getString(R.string.d1));
        k2(getResources().getString(R.string.k21));
    }
}