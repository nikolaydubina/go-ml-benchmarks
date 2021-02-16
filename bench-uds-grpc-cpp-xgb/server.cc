#include <iostream>
#include <memory>
#include <string>

#include <grpcpp/grpcpp.h>
#include <grpcpp/health_check_service_interface.h>
#include <grpcpp/ext/proto_server_reflection_plugin.h>
#include "predictor.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using predictor::PredictRequest;
using predictor::PredictResponse;
using predictor::Predictor;

class PredictorImpl final : public Predictor::Service {
  Status SayHello(ServerContext* context, const PredictRequest* request, PredictResponse* reply) override {
    reply->set_prediction(11.2);
    return Status::OK;
  }
};

void RunServer() {
  std::string server_address("0.0.0.0:50051");
  PredictorImpl service;

  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  ServerBuilder builder;

  builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<Server> server(builder.BuildAndStart());

  server->Wait();
}

int main(int argc, char** argv) {
  RunServer();
  return 0;
}
