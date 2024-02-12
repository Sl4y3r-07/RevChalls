package com.example.open_sesame;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

public class MainActivity extends AppCompatActivity {
    private EditText editTextUsername;
    private EditText editTextPassword;
    private Button buttonLogin;

    // Replace these with your correct username and password
    private static final String valid_user = "Jack Ma";
    private static final int[] valid_password = {52, 108, 49, 98, 97, 98, 97}; 

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        editTextUsername = findViewById(R.id.editTextUsername);
        editTextPassword = findViewById(R.id.editTextPassword);
        buttonLogin = findViewById(R.id.buttonLogin);

        buttonLogin.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                validateCredentials();
            }
        });
    }

    private void validateCredentials() {
        String username = editTextUsername.getText().toString().trim();
        String password = editTextPassword.getText().toString().trim();

        if (username.equals(valid_user) && n4ut1lus(password)) {
            String digitsExtracted = sh4dy(password);
            int digits = sl4y3r(digitsExtracted);
            String a ="U|]rURuoU^PoR_FDMo@X]uBUg";
            String result =  "flag{"+flag(Integer.toString(digits), a)+"}";

        } else {
            String result = "Invalid credentials. Please try again.";
            showToast(result);
        }
    }

    private boolean n4ut1lus(String input) {
        int[] asciiValues = it4chi(input);
        if (asciiValues.length != valid_password.length) {
            return false;
        }

        for (int i = 0; i < asciiValues.length; i++) {
            if (asciiValues[i] != valid_password[i]) {
                return false;
            }
        }
        return true;
    }

    private int[] it4chi(String input) {
        int[] asciiValues = new int[input.length()];
        for (int i = 0; i < input.length(); i++) {
            asciiValues[i] = (int) input.charAt(i);
        }
        return asciiValues;
    }

    private String sh4dy(String input) {
        StringBuilder extractedDigits = new StringBuilder();

        for (int i = 0; i < input.length(); i++) {
            char currentChar = input.charAt(i);
            if (Character.isDigit(currentChar)) {
                extractedDigits.append(currentChar);
            }
        }
        return extractedDigits.toString();
    }
    private int sl4y3r(String extractedDigits) {
        return Integer.parseInt(extractedDigits)-1;
    }

    private String flag(String a, String b) {
        StringBuilder result = new StringBuilder();
        for (int i = 0; i < b.length(); i++) {
            char xoredChar = (char) (b.charAt(i) ^ a.charAt(i % a.length()));
            result.append(xoredChar);
        }
        return result.toString();
    }

    private void showToast(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
    }
}
