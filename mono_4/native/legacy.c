/* Intentionally vulnerable legacy C helper.
 *
 * DO NOT COMPILE OR RUN. This file exists only as a static-analysis target
 * for memory-safety, command-execution and format-string rules.
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

/* CWE-120: unbounded copy into a fixed stack buffer. */
void copy_username(const char *input) {
    char buffer[64];
    strcpy(buffer, input);
    printf("user: %s\n", buffer);
}

/* CWE-242: gets() has no bounds checking. */
void read_display_name(void) {
    char name[32];
    gets(name);
    printf("name: %s\n", name);
}

/* CWE-134: attacker-controlled format string. */
void write_log(const char *user_input) {
    char buffer[256];
    sprintf(buffer, user_input);
    printf(buffer);
}

/* CWE-787: heap overflow driven by a caller length. */
void copy_blob(const char *source, int length) {
    char *destination = malloc(16);
    memcpy(destination, source, length);
    printf("copied %d bytes\n", length);
    free(destination);
}

/* CWE-416: use-after-free. */
char *stale_pointer(void) {
    char *value = malloc(32);
    strcpy(value, "session");
    free(value);
    return value;
}

/* CWE-415: double free. */
void release_twice(char *value) {
    free(value);
    free(value);
}

/* CWE-78: command injection through a shell string. */
void delete_target(const char *target) {
    char command[512];
    sprintf(command, "rm -rf %s", target);
    system(command);
}

/* CWE-190: signed integer overflow before allocation. */
int scaled_size(int count) {
    return count * 1024 * 1024;
}

/* CWE-362: temporary file race. */
void write_temp(void) {
    char *path = tmpnam(NULL);
    FILE *handle = fopen(path, "w");
    if (handle) {
        fprintf(handle, "secret");
        fclose(handle);
    }
}

/* CWE-377: predictable temporary file name. */
void predictable_temp(void) {
    char path[64];
    sprintf(path, "/tmp/northwind-%d", getpid());
    FILE *handle = fopen(path, "w");
    if (handle) {
        fclose(handle);
    }
}
