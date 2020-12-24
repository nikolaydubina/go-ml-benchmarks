#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <unistd.h>
#include <xgboost/c_api.h>

#define SOCKET_NAME "../xgb-go.socket"
#define BUFFER_SIZE 1
#define NCOLUMNS 1

int
main(int argc, char *argv[])
{
	// load xgb model
	BoosterHandle booster;
	XGBoosterCreate(0, 0, &booster);
	XGBoosterLoadModel(booster, "model.xgb");
	DMatrixHandle data;

	// server
	struct sockaddr_un name;
	int down_flag = 0;
	int ret;
	int connection_socket;
	int data_socket;
	int result;
	float buffer[BUFFER_SIZE];

	/* Create local socket. */
	connection_socket = socket(AF_UNIX, SOCK_SEQPACKET, 0);
	if (connection_socket == -1) {
		perror("socket");
		exit(EXIT_FAILURE);
	}

	/*
	* For portability clear the whole structure, since some
	* implementations have additional (nonstandard) fields in
	* the structure.
	*/
	memset(&name, 0, sizeof(name));

	/* Bind socket to socket name. */
	name.sun_family = AF_UNIX;
	strncpy(name.sun_path, SOCKET_NAME, sizeof(name.sun_path) - 1);

	ret = bind(connection_socket, (const struct sockaddr *) &name, sizeof(name));
	if (ret == -1) {
		perror("bind");
		exit(EXIT_FAILURE);
	}

	/*
	* Prepare for accepting connections. The backlog size is set
	* to 20. So while one request is being processed other requests
	* can be waiting.
	*/
	ret = listen(connection_socket, 20);
	if (ret == -1) {
		perror("listen");
		exit(EXIT_FAILURE);
	}

	// loop for connections
	for (;;) {
		data_socket = accept(connection_socket, NULL, NULL);
		if (data_socket == -1) {
			perror("accept");
			exit(EXIT_FAILURE);
		}

		// loop for data
		for (;;) {
			/* Wait for next data packet. */
			ret = read(data_socket, buffer, sizeof(buffer));
			if (ret == -1) {
				perror("read");
				exit(EXIT_FAILURE);
			}
		}
		
		// predictions
		XGDMatrixCreateFromMat(buffer, 1, NCOLUMNS, 0, data);

		bst_ulong out_len = 0;
		const float *out_result = NULL;
		XGBoosterPredict(booster, data, 0, 0, 0, &out_len, &out_result);
		printf("prediction %1.4f\n", out_result[1]);

		/* Send result. */
		ret = write(data_socket, out_result, out_len);
		if (ret == -1) {
			perror("write");
			exit(EXIT_FAILURE);
		}

		close(data_socket);
	}

	close(connection_socket);

	/* Unlink the socket. */
	unlink(SOCKET_NAME);
	exit(EXIT_SUCCESS);
	
	// Cleanup XGBoost
	XGDMatrixFree(data);
	XGBoosterFree(booster);
}