# Geth IPC / RPC Bench
Benchmarks comparing Geth's IPC and RPC endpoints. To start the benchmark, run `go test -bench=. -count 10`. Note that you will need to enable the `--http` flag when running Geth to expose the RPC endpoint. The current results are listed below. 

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

BenchmarkIPC-8   	    7995	    226644 ns/op
BenchmarkIPC-8   	    5187	    221686 ns/op
BenchmarkIPC-8   	    5044	    222848 ns/op
BenchmarkIPC-8   	    5132	    222848 ns/op
BenchmarkIPC-8   	    5107	    221179 ns/op
BenchmarkIPC-8   	    5457	    228048 ns/op
BenchmarkIPC-8   	    5875	    220197 ns/op
BenchmarkIPC-8   	    5378	    219127 ns/op
BenchmarkIPC-8   	    5496	    220681 ns/op
BenchmarkIPC-8   	    5396	    220749 ns/op

BenchmarkRPC-8   	    2004	    530629 ns/op
BenchmarkRPC-8   	    1958	    565381 ns/op
BenchmarkRPC-8   	    2067	    560948 ns/op
BenchmarkRPC-8   	    2040	    571529 ns/op
BenchmarkRPC-8   	    2047	    571289 ns/op
BenchmarkRPC-8   	    2253	    533267 ns/op
BenchmarkRPC-8   	    1924	    552074 ns/op
BenchmarkRPC-8   	    2023	    572780 ns/op
BenchmarkRPC-8   	    2020	    585387 ns/op
BenchmarkRPC-8   	    1971	    592632 ns/op

PASS
ok  	geth_ipc_rpc_bench	26.592s


//==================================================

Average IPC: 222400.7 ns/op

Average RPC: 513591.6 ns/op




```
