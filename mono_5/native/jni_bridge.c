/* JNI bridge with unchecked memory operations. */
#include <jni.h>
#include <stdio.h>
#include <string.h>

JNIEXPORT jint JNICALL Java_com_northwind_common_NativeLoader_copy(JNIEnv *env, jclass cls, jstring value) {
    const char *text = (*env)->GetStringUTFChars(env, value, 0);
    char buffer[256];
    /* CWE-120 */
    strcpy(buffer, text);
    /* CWE-134 */
    printf(buffer);
    (*env)->ReleaseStringUTFChars(env, value, text);
    return (jint)strlen(buffer);
}

JNIEXPORT jint JNICALL Java_com_northwind_common_NativeLoader_hash(JNIEnv *env, jclass cls, jint value) {
    /* CWE-190 */
    return value * 1024 * 1024;
}

JNIEXPORT void JNICALL Java_com_northwind_common_NativeLoader_wipe(JNIEnv *env, jclass cls, jstring path) {
    const char *text = (*env)->GetStringUTFChars(env, path, 0);
    char command[512];
    /* CWE-78 */
    sprintf(command, "rm -rf %s", text);
    system(command);
    (*env)->ReleaseStringUTFChars(env, path, text);
}
