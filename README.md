# go-ml-benchmarks

> Given I have a single struct in my Go service, how quickly can I get ML inference result?

TODO: illustration of latencies breakdowns
- [Go port of XGBoost and scikit-learn](https://github.com/dmitryikh/leaves)

TODO: matrix of latencies

- [ ] flatbuffers - unixsocket - python flatbuffers - sklearn - xgb
- [ ] json - unixsocket - python rapidjson - sklearn - xgb
- [ ] grpc - tcp - python - sklearn - xgb
- [ ] go-featureprocessing - cgo - xgb
- [ ] [cgo bindings fo XGBoost](https://github.com/Unity-Technologies/go-xgboost)

```
BenchmarkXGB_GoFeatureProcessing_GoLeaves                       25049485               412 ns/op
BenchmarkXGB_GoFeatureProcessing_UDS_RawBytes_Python_XGB           54088            230235 ns/op
BenchmarkXGB_HTTP_JSON_Python_Gunicorn_Flask_sklearn_XGB             501          22995328 ns/op
```

## Dataset and Model

We are using here classic [Titanic dataset](https://www.kaggle.com/c/titanic).
It contains numerical and categorical features which makes it representative of a typical scenario.
Data as well notebook that trains model and preprocessor is available in /data and /notebooks respectively.

## Some numbers for reference

How fast do you need to get?

```
                     ...
                   200ps - 4.6GHz single cycle time
                1ns      - L1 cache latency
               10ns      - L2/L3 cache SRAM latency
               20ns      - DDR4 CAS, first byte from memory latency
               20ns      - C++ raw hardcoded structs access
               80ns      - C++ FlatBuffers decode/traverse/dealloc
              100ns      - go-featureprocessing typical processing
              150ns      - PCIe bus latency
              171ns      - Go cgo call boundary, 2015
              200ns      - some High Frequency Trading FPGA claims
 ---------->  400ns      - gofeatureprocessing + leaves
              800ns      - Go Protocol Buffers Marshal
              837ns      - Go json-iterator/go json decode
           1µs           - Go Protocol Buffers Unmarshal
           1µs           - High Frequency Trading FPGA
           3µs           - Go JSON Marshal
           7µs           - Go JSON Unmarshal
           9µs           - Go XML Marshal
          10µs           - PCIe/NVLink startup time
          17µs           - Python JSON encode or decode times
          30µs           - UNIX domain socket, eventfd, fifo pipes latency
          30µs           - Go XML Unmarshal
         100µs           - Redis intrinsic latency
         100µs           - AWS DynamoDB + DAX
         100µs           - KDB+ queries
         100µs           - High Frequency Trading direct market access range
         200µs           - 1GB/s network air latency
         200µs           - Go garbage collector latency 2018
         500µs           - NGINX/Kong added latency
     10ms                - AWS DynamoDB
     10ms                - WIFI6 "air" latency
     15ms                - AWS Sagemaker latency
     30ms                - 5G "air" latency
    100ms                - typical roundtrip from mobile to backend
    200ms                - AWS RDS MySQL/PostgreSQL or AWS Aurora
 10s                     - AWS Cloudfront 1MB transfer time
```

## Reference

- [how-to article for Go leaves module](https://dev.to/blairhudson/machine-learning-microservices-python-and-xgboost-in-a-tiny-486kb-container-4on4)
- [cgo performance](https://about.sourcegraph.com/go/gophercon-2018-adventures-in-cgo-performance/)
- [cgo goroutines are not as performant](https://www.cockroachlabs.com/blog/the-cost-and-complexity-of-cgo/)
- [Datadog shows how to call CPython with cgo](https://www.datadoghq.com/blog/engineering/cgo-and-python/)
- [EuroPython2019 shows how to call CPython with cgo, suggests Linux fifo to speedup](https://ep2019.europython.eu/talks/Zktoaai-golang-to-python/)
- [Linux IPC comparison "Evaluation of Inter-Process Communication Mechanisms", Aditya, Kishore](http://pages.cs.wisc.edu/~adityav/Evaluation_of_Inter_Process_Communication_Mechanisms.pdf)
- [MLCommons MLPerf benchmarks](https://github.com/mlcommons/inference)
- [Redis latency](https://redis.io/topics/latency)
- [Huawei wifi6 latency](https://e.huawei.com/sg/products/enterprise-networking/wlan/wifi-6)
- [Verizon 5G latency](https://www.verizon.com/about/our-company/5g/5g-latency)
- [NGINX added latency](https://www.nginx.com/blog/nginx-controller-api-management-module-vs-kong-performance-comparison/)
- [AWS Sagemaker latency](https://aws.amazon.com/blogs/machine-learning/load-test-and-optimize-an-amazon-sagemaker-endpoint-using-automatic-scaling/)
- [AWS Cloudfront transfer rates](https://media.amazonwebservices.com/FS_WP_AWS_CDN_CloudFront.pdf)
- [Go garbage collector updates 2018](https://blog.golang.org/ismmkeynote)
- [HFT latency](https://en.wikipedia.org/wiki/Ultra-low_latency_direct_market_access)
- [AWS Aurora latency](https://aws.amazon.com/blogs/database/using-aurora-to-drive-3x-latency-improvement-for-end-users/)
- [UNIX local IPC latencies](http://kamalmarhubi.com/blog/2015/06/10/some-early-linux-ipc-latency-data/)
- [FPGA HFT latency](https://ieeexplore.ieee.org/document/6299067)
- [Google TPU latency](https://ai.googleblog.com/2019/08/efficientnet-edgetpu-creating.html)
- [HFT FPGA 200nanoseconds 2018](https://apnews.com/press-release/pr-businesswire/2edb1f8f12d64ab490ef0c180e648e24)
- [PCIe latency](https://www.cl.cam.ac.uk/research/srg/netos/projects/pcie-bench/neugebauer2018understanding.pdf)
- ["Evaluating Modern GPU Interconnect: PCIe, NVLink, NV-SLI, NVSwitch and GPUDirect", 2019](?)
- [Cache and DRAM latency](https://en.wikipedia.org/wiki/CPU_cache)
- [Go feature extractor](https://github.com/dustin-decker/featuremill)
- [Go port of sklearn](https://github.com/pa-m/sklearn)
