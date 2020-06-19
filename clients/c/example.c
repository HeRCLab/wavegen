#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "client.h"

int main(int argc, char** argv) {
	if (argc != 2) {
		fprintf(stderr, "usage: example JSON_FILE\n");
		exit(1);
	}

	int handle;

	if (WGOpen(argv[1], &handle) != 0) {
		fprintf(stderr, "ERROR: %s\n", WGGetError());
		exit(1);
	}
	printf("Opened handle %i on file '%s'\n", handle, argv[1]);

	int size;
	if (WGSize(handle, &size) != 0) {
		fprintf(stderr, "ERROR: %s\n", WGGetError());
		exit(1);
	}
	printf("Handle contains %i samples\n", size);

	double* S = malloc(sizeof(double) * size);
	double* T = malloc(sizeof(double) * size);
	double SampleRate;

	if (WGSampleRate(handle, &SampleRate) != 0) {
		fprintf(stderr, "ERROR: %s\n", WGGetError());
		exit(1);
	}
	printf("Retrieved sample rate %f\n", SampleRate);

	if (WGCopyS(handle, S) != 0) {
		fprintf(stderr, "ERROR: %s\n", WGGetError());
		exit(1);
	}

	if (WGCopyT(handle, T) != 0) {
		fprintf(stderr, "ERROR: %s\n", WGGetError());
		exit(1);
	}

	if (WGClose(handle) != 0) {
		fprintf(stderr, "ERROR: %s\n", WGGetError());
		exit(1);
	}

	printf("First 10 samples... ");
	for (int i = 0 ; i < 10 ; i++) {
		printf("S[%i]=%f T[%i]=%f\n", i, S[i], i, T[i]);
	}

	free(S);
	free(T);
}


