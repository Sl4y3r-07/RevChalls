package com.backdoor.sl4ydroid;

import androidx.appcompat.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.os.Bundle;
import android.widget.TextView;




public class MainActivity extends AppCompatActivity {
    static {
        System.loadLibrary("sl4ydroid");
    }
    private TextView displayTextView;
    private TextView textView;
    public void sh4dy(String message) {
        System.out.println(message);
        runOnUiThread(() -> {
            String currentText = displayTextView.getText().toString();
            String newText = currentText+ message;
            displayTextView.setText(newText);
        });
    }

    public void sl4y3r(String message) {
        System.out.println(message);
        runOnUiThread(() -> {
            String currentText = displayTextView.getText().toString();
            String newText = currentText  + message;
            displayTextView.setText(newText);
        });
    }
    public void it4chi(String message) {
        System.out.println(message);
        runOnUiThread(() -> {
            String currentText = displayTextView.getText().toString();
            String newText = currentText  + message;
            displayTextView.setText(newText);
        });
    }
    public void n4ut1lus(String message) {
        System.out.println(message);
        runOnUiThread(() -> {
            String currentText = displayTextView.getText().toString();
            String newText = currentText + message;
            displayTextView.setText(newText);
        });
    }
    public native void kim(String message);
    public native void nim(String message);
    public native void damn(String message);
    public native void k2(String message);

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        displayTextView = findViewById(R.id.displayTextView);
        displayTextView.setVisibility(View.GONE);
        textView = findViewById(R.id.textView);
        textView.setText(getResources().getString(R.string.message));

        kim(getResources().getString(R.string.k1));
        nim(getResources().getString(R.string.n1));
        damn(getResources().getString(R.string.d1));
        k2(getResources().getString(R.string.k21));
    }


}
