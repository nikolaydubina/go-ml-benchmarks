#include <iostream>
#include <memory>
#include <string>
#include <cstdlib>
#include <stdio.h>
#include <exception>

#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>
#include "predictor.grpc.pb.h"

#include <xgboost/c_api.h>

#define N_FEATURES 12

class PredictorImpl final : public predictor::Predictor::Service {
  private:

  BoosterHandle booster;

  public:

  PredictorImpl() {
    const char* model_path = std::getenv("MODEL_PATH");
    XGBoosterCreate(0, 0, &booster);
    XGBoosterLoadModel(booster, model_path);
  }

  ~PredictorImpl() {
    XGBoosterFree(booster);
  }

  grpc::Status PredictProcessed(grpc::ServerContext* context, const predictor::PredictProcessedRequest* request, predictor::PredictResponse* reply) override {
    bst_ulong out_len = 0;
    const float* out_result = NULL;

    if (request->features_size() != N_FEATURES) {
      return grpc::Status::CANCELLED;
    }

    float features[N_FEATURES];
    for (int i = 0; i < request->features_size(); i++) {
      features[i] = (float)request->features(i);
    }

    DMatrixHandle dmatrix;
    XGDMatrixCreateFromMat(features, 1, N_FEATURES, -1, &dmatrix);

    XGBoosterPredict(booster, dmatrix, 0, 0, 0, &out_len, &out_result);
    XGDMatrixFree(dmatrix);

    if (out_len != 1) {
      return grpc::Status::CANCELLED;
    }

    reply->set_prediction(out_result[0]);
    
    return grpc::Status::OK;
  }
};

void RunServer() {
  std::string server_address("unix:///tmp/test.sock");
  PredictorImpl service;

  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  grpc::ServerBuilder builder;

  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<grpc::Server> server(builder.BuildAndStart());

  server->Wait();
}

int main(int argc, char** argv) {
  RunServer();
  return 0;
}
