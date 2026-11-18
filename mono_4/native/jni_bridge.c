/* JNI bridge with an unchecked length and a format-string sink. */
#include <jni.h>
#include <stdio.h>
#include <string.h>

JNIEXPORT void JNICALL Java_com_northwind_Native_copy(JNIEnv *env, jobject obj, jstring value) {
    const char *text = (*env)->GetStringUTFChars(env, value, 0);
    char buffer[128];
    /* CWE-120 */
    strcpy(buffer, text);
    /* CWE-134 */
    printf(buffer);
    (*env)->ReleaseStringUTFChars(env, value, text);
}

JNIEXPORT jint JNICALL Java_com_northwind_Native_hash(JNIEnv *env, jobject obj, jint count) {
    /* CWE-190 */
    return count * 1024 * 1024;
}
