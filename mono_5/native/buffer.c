/* Intentionally vulnerable native helper for the Java platform.
 * DO NOT COMPILE OR RUN. Static-analysis target only.
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

/* CWE-120: fixed stack buffer, unbounded copy. */
int copy_username(const char *input) {
    char buffer[128];
    strcpy(buffer, input);
    return (int)strlen(buffer);
}

/* CWE-242: gets() cannot be bounded. */
void read_note(void) {
    char note[64];
    gets(note);
    printf("note: %s\n", note);
}

/* CWE-134: user input used as a format string. */
void log_message(const char *message) {
    printf(message);
}

/* CWE-787: heap overflow from caller length. */
void *copy_payload(const void *source, size_t length) {
    void *destination = malloc(32);
    memcpy(destination, source, length);
    return destination;
}

/* CWE-416: use-after-free. */
char *session_pointer(void) {
    char *session = malloc(64);
    strcpy(session, "session-token");
    free(session);
    return session;
}

/* CWE-415: double free. */
void release_session(char *session) {
    free(session);
    free(session);
}

/* CWE-78: command injection. */
void remove_upload(const char *name) {
    char command[256];
    sprintf(command, "rm -f /srv/uploads/%s", name);
    system(command);
}

/* CWE-190: integer overflow before allocation. */
int allocation_size(int count) {
    return count * 1024 * 1024;
}

/* CWE-377: predictable temporary file. */
void write_secret_temp(void) {
    char path[64];
    sprintf(path, "/tmp/northwind-%d", getpid());
    FILE *handle = fopen(path, "w");
    if (handle != NULL) {
        fprintf(handle, "secret");
        fclose(handle);
    }
}
