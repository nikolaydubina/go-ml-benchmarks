PWD := $(shell echo $$PWD)
MY_INSTALL_DIR := $(shell echo $$HOME/.local)
SYSTEM ?= $(shell uname | cut -f 1 -d_)
CXX = g++
CPPFLAGS += `pkg-config --cflags protobuf grpc`
CXXFLAGS += -std=c++11
ifeq ($(SYSTEM),Darwin)
LDFLAGS += -L/usr/local/lib `pkg-config --libs protobuf grpc++` -pthread -lgrpc++_reflection -ldl
else
LDFLAGS += -L/usr/local/lib `pkg-config --libs protobuf grpc++` -pthread -Wl,--no-as-needed -lgrpc++_reflection -Wl,--as-needed -ldl
endif
PROTOC = protoc
GRPC_CPP_PLUGIN = grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH ?= `which $(GRPC_CPP_PLUGIN)`
PROTOS_PATH = ../proto

vpath %.proto $(PROTOS_PATH)

init:
	mkdir -p $(MY_INSTALL_DIR)
	export PATH="$$PATH:$(MY_INSTALL_DIR)/bin"
ifeq ($(SYSTEM),Darwin)
	brew install protobuf autoconf automake libtool pkg-config
	-brew install smake
endif
	-git clone --recurse-submodules -b v1.35.0 https://github.com/grpc/grpc
	cd grpc; \
		mkdir -p cmake/build; \
		pushd cmake/build; \
		cmake -DgRPC_INSTALL=ON -DgRPC_BUILD_TESTS=OFF -DCMAKE_INSTALL_PREFIX=$(MY_INSTALL_DIR) ../..; \
		make -j; \
		make install; \
		popd

build:
	mkdir -p cmake/build
	cd cmake/build
	cmake -DCMAKE_PREFIX_PATH=$(MY_INSTALL_DIR) ../..
	make -j

predictor: predictor.pb.o predictor.grpc.pb.o predictor.o
	$(CXX) $^ $(LDFLAGS) -o $@

.PRECIOUS: %.grpc.pb.cc
%.grpc.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --grpc_out=. --plugin=protoc-gen-grpc=$(GRPC_CPP_PLUGIN_PATH) $<

.PRECIOUS: %.pb.cc
%.pb.cc: %.proto
	$(PROTOC) -I $(PROTOS_PATH) --cpp_out=. $<

clean:
	rm -f *.o *.pb.cc *.pb.h server