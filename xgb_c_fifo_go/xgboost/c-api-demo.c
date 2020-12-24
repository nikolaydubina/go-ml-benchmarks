/**
 * Using FIFO pipes to compute XGBoost inference.
 * Input is expected to be little-endian binary of float32.
 */
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h> 
#include <sys/types.h> 
#include <sys/un.h>
#include <unistd.h>
#include <xgboost/c_api.h>

#define FIFO_FEATURES "../features.fifo"
#define FIFO_PREDICTIONS "../predictions.fifo"
#define NFEATURES 13

int main(int argc, char *argv[]) {
	DMatrixHandle data;
	BoosterHandle booster;

	XGBoosterCreate(0, 0, &booster);
	XGBoosterLoadModel(booster, "model.xgb");
  
    mkfifo(FIFO_FEATURES, 0666); 
    mkfifo(FIFO_PREDICTIONS, 0666); 
  
    float features[NFEATURES];
    while (1) { 
		int fd_features = open(FIFO_FEATURES, O_RDONLY); 
        read(fd_features, features, sizeof(features));
		close(fd_features);
		
		XGDMatrixCreateFromMat(features, 1, NFEATURES, 0.0, &data);

		bst_ulong predictions_len = 0;
		const float *predictions = NULL;
		XGBoosterPredict(booster, data, 0, 0, 0, &predictions_len, &predictions);
  
		int fd_predictions = open(FIFO_PREDICTIONS, O_WRONLY); 
        write(fd_predictions, predictions, sizeof(predictions)); 
		close(fd_predictions);
    } 
	
	XGDMatrixFree(data);
	XGBoosterFree(booster);
}