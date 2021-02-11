# go-ml-benchmarks

> Given I have input data for a single struct in a Go service, how quickly can I get inference result?

TODO: illustration of latencies breakdowns
TODO: matrix of latencies

- [ ] flatbuffers - unixsocket - python flatbuffers - sklearn - xgb
- [ ] json - unixsocket - python rapidjson - sklearn - xgb
- [ ] grpc - tcp - python - sklearn - xgb
- [ ] go-featureprocessing - cgo - xgb

```
BenchmarkXGB_gofeatureprocessing_goleaves               29092437               398 ns/op
BenchmarkXGB_Python_UDS_RawBytes_NewConnection             52749            227066 ns/op
BenchmarkXGB_Python_JSON_Gunicorn_Flask_sklearn_xgb          508          22985758 ns/op
```

## Some numbers for reference

TODO

## Reference and Related Work

Go modules
- https://github.com/pytorch/serve
- [Go port of XGBoost and scikit-learn](https://github.com/dmitryikh/leaves)
- [cgo bindings fo XGBoost](https://github.com/Unity-Technologies/go-xgboost)

Articles
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
