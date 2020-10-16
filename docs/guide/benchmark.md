# 压力测试

下面是使用objpoint类型做的压力测试

```
$ go test -bench=. -run=non                         

goos: darwin
goarch: amd64
pkg: github.com/jianfengye/collection
Benchmark_Append-12        	 3512688	       387 ns/op
Benchmark_Contain-12       	 6727482	       179 ns/op
Benchmark_Copy-12          	 7260177	       159 ns/op
Benchmark_Diff-12          	 2310327	       522 ns/op
Benchmark_Each-12          	 7784914	       154 ns/op
Benchmark_Every-12         	 7602790	       157 ns/op
Benchmark_ForPage-12       	 2355352	       515 ns/op
Benchmark_Filter-12        	 1356804	       876 ns/op
Benchmark_First-12         	19379992	        61.8 ns/op
Benchmark_Index-12         	19259961	        62.1 ns/op
Benchmark_IsEmpty-12       	162860646	         7.33 ns/op
Benchmark_IsNotEmpty-12    	163036106	         7.36 ns/op
Benchmark_Join-12          	 4705460	       255 ns/op
Benchmark_Last-12          	15544176	        76.8 ns/op
Benchmark_Merge-12         	 1372609	       872 ns/op
Benchmark_Map-12           	 2752177	       439 ns/op
Benchmark_Max-12           	 3218686	       372 ns/op
Benchmark_Min-12           	 3233270	       372 ns/op
Benchmark_Median-12        	 1379985	       869 ns/op
Benchmark_Nth-12           	 2360064	       503 ns/op
Benchmark_Pop-12           	 1454916	       834 ns/op
Benchmark_Push-12          	 3629934	       346 ns/op
Benchmark_Prepend-12       	   10000	    376298 ns/op
Benchmark_Pluck-12         	 2531895	       469 ns/op
Benchmark_Reject-12        	 4184707	       286 ns/op
Benchmark_Random-12        	  142698	      8397 ns/op
Benchmark_Reverse-12       	 1324262	       903 ns/op
Benchmark_Slice-12         	 2272142	       515 ns/op
Benchmark_Search-12        	 6484984	       186 ns/op
Benchmark_Sort-12          	 3627673	       333 ns/op
Benchmark_SortDesc-12      	 3565390	       331 ns/op
Benchmark_Shuffle-12       	  128826	      9320 ns/op
Benchmark_SortBy-12        	564669482	         2.13 ns/op
Benchmark_SortByDesc-12    	595491585	         2.03 ns/op
Benchmark_Unique-12        	 1219267	       979 ns/op
PASS
ok  	github.com/jianfengye/collection	59.484s
```