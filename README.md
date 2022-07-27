# Geth IPC / RPC Bench
Benchmarks comparing Geth's IPC and RPC endpoints. To start the benchmark, run `go test -bench=. -count 10`. The current results are listed below. 

## Testing environment hardware  
- Ryzen 5 3600G
- 32G DDR4-3200 RAM
- 1TB WD Black SN750 NVME


## Results

```

goos: linux
goarch: amd64
pkg: geth_ipc_rpc_bench
cpu: AMD Ryzen 5 3400G with Radeon Vega Graphics    
BenchmarkIPC-8   	    6104	    222998 ns/op
BenchmarkIPC-8   	    5257	    229704 ns/op
BenchmarkIPC-8   	    4974	    225137 ns/op
BenchmarkIPC-8   	    4657	    233986 ns/op
BenchmarkIPC-8   	    4971	    270283 ns/op
BenchmarkIPC-8   	   10000	    194288 ns/op
BenchmarkIPC-8   	    5424	    232378 ns/op
BenchmarkIPC-8   	    5090	    232598 ns/op
BenchmarkIPC-8   	    5032	    207802 ns/op
BenchmarkIPC-8   	    8047	    218459 ns/op

BenchmarkRPC-8   	    4752	    219878 ns/op
BenchmarkRPC-8   	    4905	    241542 ns/op
BenchmarkRPC-8   	    7490	    230226 ns/op
BenchmarkRPC-8   	    6016	    168425 ns/op
BenchmarkRPC-8   	    4558	    248492 ns/op
BenchmarkRPC-8   	    4699	    239182 ns/op
BenchmarkRPC-8   	    6621	    213950 ns/op
BenchmarkRPC-8   	    5548	    247520 ns/op
BenchmarkRPC-8   	    5222	    242062 ns/op
BenchmarkRPC-8   	    4585	    236007 ns/op
PASS
ok  	geth_ipc_rpc_bench	26.216s


//==================================================

Average IPC: 226763.3 ns/op

Average RPC: 228728.4 ns/op




```
