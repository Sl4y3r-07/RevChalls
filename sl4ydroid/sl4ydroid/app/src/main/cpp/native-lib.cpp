// NativeLibrary.cpp

#include <jni.h>
#include <string>
#include <iostream>
#include <vector>
#include <iomanip>
#include <sstream>

std::string firsts()
{
    std::string encoded ="b5)c]d/ZP1:\"";
    std::string str1 ="R00rkee"; std::string str2 ="R1zz";

    for (int i = 0;i<=encoded.length()-1; i++) {
        encoded[i]= encoded[i]^str1[i%(str1.length())]^str2[i%str2.length()];
    }
    return encoded;
}

std::string s1 ="bb7f8f9e69";
std::string s2 ="bb7b8f832383";
std::string r1 ="a1638f997898";
std::string r2 ="a623ce8e219dd233";
std::string v2="May_1??";
std::string us3_m3 = firsts();
std::string try_m3 = "can I help you?";
std::string m1 = "d23ba52679c5e1";

std::string zamn(const std::string &input) {
    std::string bin;
    for (size_t i = 0; i < input.length(); i += 2) {
        std::string bytes = input.substr(i, 2);
        char byte = (char)std::stoi(bytes, nullptr, 16);
        bin += byte;
    }
    return bin;
}
std::string ring(const std::string &input, const std::string &key) {
    std::string decrypt;
    std::string encrypt = zamn(input);
    std::vector<int> S(256);
    int i = 0, j = 0;


    for (int i = 0; i < 256; ++i)
        S[i] = i;

    for (int i = 0; i < 256; ++i) {
        j = (j + S[i] + key[i % key.length()]) % 256;
        std::swap(S[i], S[j]);
    }

    i = 0;
    j = 0;


    for (char c : encrypt) {
        i = (i + 1) % 256;
        j = (j + S[i]) % 256;
        std::swap(S[i], S[j]);
        int K = S[(S[i] + S[j]) % 256];
        decrypt += char(c ^ K);
    }

    return decrypt;
}
extern "C" JNIEXPORT void JNICALL Java_com_backdoor_sl4ydroid_MainActivity_kim(JNIEnv *env, jobject obj, jstring message) {
    jclass dynamicMethodClass = env->GetObjectClass(obj);
    std::string name =ring(s1,us3_m3);
    const char* methodName = name.c_str();
    jmethodID dynamicMethod = env->GetMethodID(dynamicMethodClass, methodName , "(Ljava/lang/String;)V");
    const char *nativeMessage = env->GetStringUTFChars(message, JNI_FALSE);
    std::string encrypted(nativeMessage);
    env->ReleaseStringUTFChars(message, nativeMessage);

     for (char &c: encrypted) {
            c = c + 8;
            c = c^7;
     }

    jstring jniMessage = env->NewStringUTF(encrypted.c_str());
    env->CallVoidMethod(obj, dynamicMethod, jniMessage);
    env->DeleteLocalRef(jniMessage);
}
extern "C" JNIEXPORT void JNICALL Java_com_backdoor_sl4ydroid_MainActivity_nim(JNIEnv *env, jobject obj, jstring message) {
    jclass dynamicMethodClass = env->GetObjectClass(obj);
    std::string name =ring(s2,us3_m3);
    const char* methodName = name.c_str();
    jmethodID dynamicMethod = env->GetMethodID(dynamicMethodClass, methodName, "(Ljava/lang/String;)V");
    const char *nativeMessage = env->GetStringUTFChars(message, JNI_FALSE);
    std::string encoded(nativeMessage);
    env->ReleaseStringUTFChars(message, nativeMessage);

    for (int i = 0; i <= encoded.length() - 1; i++) {
        std::string ki =  ring(m1,try_m3);
        encoded[i] = encoded[i] ^ki[i%ki.length()]^19;
    }

    jstring jniMessage = env->NewStringUTF(encoded.c_str());
    env->CallVoidMethod(obj, dynamicMethod, jniMessage);
    env->DeleteLocalRef(jniMessage);
}

extern "C" JNIEXPORT void JNICALL Java_com_backdoor_sl4ydroid_MainActivity_damn(JNIEnv *env, jobject obj, jstring message) {
    jclass dynamicMethodClass = env->GetObjectClass(obj);
    std::string name =ring(r1,us3_m3);
    const char* methodName = name.c_str();
    jmethodID dynamicMethod = env->GetMethodID(dynamicMethodClass, methodName, "(Ljava/lang/String;)V");
    const char *nativeMessage = env->GetStringUTFChars(message, JNI_FALSE);
    std::string encoded(nativeMessage);
    env->ReleaseStringUTFChars(message, nativeMessage);

    std::string copy_encoded = encoded;
    for (int i = copy_encoded.length()-1; i>=0; i--) {
        encoded[(encoded.length()-1)-i]= copy_encoded[i]^12;
    }

    jstring jniMessage = env->NewStringUTF(encoded.c_str());
    env->CallVoidMethod(obj, dynamicMethod, jniMessage);
    env->DeleteLocalRef(jniMessage);
}

extern "C" JNIEXPORT void JNICALL Java_com_backdoor_sl4ydroid_MainActivity_k2(JNIEnv *env, jobject obj, jstring message) {
    jclass dynamicMethodClass = env->GetObjectClass(obj);
    std::string name =ring(r2,us3_m3);
    const char* methodName = name.c_str();
    jmethodID dynamicMethod = env->GetMethodID(dynamicMethodClass, methodName, "(Ljava/lang/String;)V");
    const char *nativeMessage = env->GetStringUTFChars(message, JNI_FALSE);
    std::string encoded(nativeMessage);
    env->ReleaseStringUTFChars(message, nativeMessage);

    for (int i=0;i<=encoded.length()-1;i++) {
       encoded[i]= encoded[i]^v2[i%(v2.length())];
    }

    jstring jniMessage = env->NewStringUTF(encoded.c_str());
    env->CallVoidMethod(obj, dynamicMethod, jniMessage);
    env->DeleteLocalRef(jniMessage);
}


