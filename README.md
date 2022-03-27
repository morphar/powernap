[![Go Reference](https://pkg.go.dev/badge/github.com/morphar/powernap.svg)](https://pkg.go.dev/github.com/morphar/powernap)

# powernap
A more precise sleep for Go.

[time.Sleep](https://pkg.go.dev/time#Sleep) only promises to sleep for **at least** the duration specified.  
The amount of overshooting can be quite large and totally random.  
When microsecond precision is needed, `powernap` can be the solution you need.  
Depending on your system, it becomes relatively precise above 1-10 µs.   

`powernap` was created for the cases, where you're willing to give up CPU cycles for precision.

When possible, a safe part of the duration is done with [time.Sleep](https://pkg.go.dev/time#Sleep), which probably happens around 1-10 ms, the rest is done by looping until the total duration has elapsed. 

### WARNING!
`powernap` can undershoot and makes no promises of sleeping "at least" for the duration given.  
The promise is: Spend CPU cycles to get as close as possible to the wanted duration, but be smart about it.

## Benchmark
It can be hard to describe the precise effect, so the following benchmark will show you the differences.  
You can run this benchmark yourself with `go test -bench .`
The last value `ns/op` should be as close as possible to the test names, e.g.: `100ns`.

Concurrently on 8 cores:
```bash
$ go test -run none -bench Sleep
goos: darwin
goarch: amd64
pkg: github.com/morphar/powernap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkNativeSleep/1ns-8         	 3668739	       324.8 ns/op
BenchmarkNativeSleep/10ns-8        	 3734082	       320.1 ns/op
BenchmarkNativeSleep/100ns-8       	 3241970	       370.3 ns/op
BenchmarkNativeSleep/1µs-8         	  256581	      4750 ns/op
BenchmarkNativeSleep/10µs-8        	   68962	     17406 ns/op
BenchmarkNativeSleep/100µs-8       	   10000	    120359 ns/op
BenchmarkNativeSleep/1ms-8         	     972	   1263703 ns/op
BenchmarkNativeSleep/10ms-8        	     100	  10777938 ns/op
BenchmarkNativeSleep/100ms-8       	      10	 101013503 ns/op
BenchmarkNativeSleep/1s-8          	       1	1001136413 ns/op
BenchmarkSleep/1ns-8               	304529434	         3.826 ns/op
BenchmarkSleep/10ns-8              	312554546	         3.814 ns/op
BenchmarkSleep/100ns-8             	313016788	         3.922 ns/op
BenchmarkSleep/1µs-8               	 1326577	       915.8 ns/op
BenchmarkSleep/10µs-8              	  119400	     10098 ns/op
BenchmarkSleep/100µs-8             	   10000	    100199 ns/op
BenchmarkSleep/1ms-8               	    1198	   1000308 ns/op
BenchmarkSleep/10ms-8              	     100	  10000445 ns/op
BenchmarkSleep/100ms-8             	      10	 100008614 ns/op
BenchmarkSleep/1s-8                	       1	1000001382 ns/op
BenchmarkSleepTight/1ns-8          	305970579	         3.951 ns/op
BenchmarkSleepTight/10ns-8         	311431710	         3.840 ns/op
BenchmarkSleepTight/100ns-8        	311112458	         3.911 ns/op
BenchmarkSleepTight/1µs-8          	 1328527	       905.6 ns/op
BenchmarkSleepTight/10µs-8         	  120620	     10033 ns/op
BenchmarkSleepTight/100µs-8        	   10000	    100705 ns/op
BenchmarkSleepTight/1ms-8          	    1200	   1000417 ns/op
BenchmarkSleepTight/10ms-8         	     100	  10000474 ns/op
BenchmarkSleepTight/100ms-8        	      10	 100003390 ns/op
BenchmarkSleepTight/1s-8           	       1	1000001256 ns/op
```

Single core:
```bash
dan@DansMac powernap % go test -run none -bench Sleep -cpu 1
goos: darwin
goarch: amd64
pkg: github.com/morphar/powernap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkNativeSleep/1ns         	 5196897	       246.4 ns/op
BenchmarkNativeSleep/10ns        	 5188018	       234.8 ns/op
BenchmarkNativeSleep/100ns       	 4249675	       271.0 ns/op
BenchmarkNativeSleep/1µs         	  179382	      9319 ns/op
BenchmarkNativeSleep/10µs        	   73255	     17105 ns/op
BenchmarkNativeSleep/100µs       	   10000	    121338 ns/op
BenchmarkNativeSleep/1ms         	    1053	   1189934 ns/op
BenchmarkNativeSleep/10ms        	     100	  10744706 ns/op
BenchmarkNativeSleep/100ms       	      10	 100730162 ns/op
BenchmarkNativeSleep/1s          	       1	1001129364 ns/op
BenchmarkSleep/1ns               	306946417	         3.837 ns/op
BenchmarkSleep/10ns              	306756348	         3.830 ns/op
BenchmarkSleep/100ns             	309193824	         3.897 ns/op
BenchmarkSleep/1µs               	 1242044	       950.1 ns/op
BenchmarkSleep/10µs              	  119083	     10107 ns/op
BenchmarkSleep/100µs             	   10000	    100318 ns/op
BenchmarkSleep/1ms               	    1198	   1000572 ns/op
BenchmarkSleep/10ms              	     100	  10001685 ns/op
BenchmarkSleep/100ms             	      10	 100004715 ns/op
BenchmarkSleep/1s                	       1	1000001456 ns/op
BenchmarkSleepTight/1ns          	299916702	         3.885 ns/op
BenchmarkSleepTight/10ns         	308808326	         3.813 ns/op
BenchmarkSleepTight/100ns        	307965261	         3.917 ns/op
BenchmarkSleepTight/1µs          	 1268040	       950.3 ns/op
BenchmarkSleepTight/10µs         	  119745	     10036 ns/op
BenchmarkSleepTight/100µs        	   10000	    100139 ns/op
BenchmarkSleepTight/1ms          	    1197	   1000479 ns/op
BenchmarkSleepTight/10ms         	     100	  10000683 ns/op
BenchmarkSleepTight/100ms        	      10	 100000765 ns/op
BenchmarkSleepTight/1s           	       1	1000001293 ns/op
```

## Acknowledgement
The basic idea came from the [Rust](https://www.rust-lang.org) crate (library) [spin_sleep](https://crates.io/crates/spin_sleep).  

