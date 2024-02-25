# Project Details:

The bootstrap was the statistical method of choice for this project. The popular Iris dataset was chosen as the data to test the statisical analysis used to compare R with Go. The bootstrap is performed on the first 2 columns of the data that are seapl width and sepal length. The bootstrap function then performs 1000 iterations of resampling the data with replacment and calcuating varisou statistics that include, correlation of width and length, sepal width median value, sepal length median value, sepal width mean, and sepal length mean. After running both programs the reults do yeild very simialar values for all statistics calculated.

## Package Details:

The package of choice used for the R program is the `boot` library that is included in base R. The `pryr` library was installed to run the CPU and memory monitoring processes. The bootstrap was built as its own Go program without relying on a thirdparty for that function, like was seen in R. The `math/rand` package was imported for the use of the random sample in the bootstrap function and the third party packages `github.com/gonum/stat` and `github.com/montanaflynn/stats` were used for the calculations of the statistics themselves. It is worth noting that the search of a bootstrap package on go.dev yeilded mixed results. The first two hits gave a package for bootstraping a server, while the third result discussed implementing bootstrapping logic with lackluster documentation. For the sake of time I did not deep dive into selecting the correct bootstrap package but instead built the bootstrap function into the program. As demonstrated in the next section, this is one area that CPU could potentionaly be conserved, with a third party package.

### Memory & CPU Usage:

The R program has the memory and time built into the program. When run the putput yields

"System time:"
   user  system elapsed
   0.26    0.02    0.39
"Memory usage:"
40.2 MB

In contrast to R, the Go program has a seperate testing and benchmark program built in the `stats_test.go`program to analyize run time and memory usage as well as unit testing. When the benchmark is run it yields:

5	 220400740 ns/op	18714083 B/op	   25688 allocs/op
PASS
ok  	goStats/appliedStats	2.335s

Memory usage: 18.71 MB and 2.335 secs per iteration. However, when the executable is run using the bash command:
$ start=$(date +%s%N)
./appliedStats.exe

Time taken: 532 ms

Logging was conducted specifically for reading in the data and conducting the bootstrap operation in the main function. Here is what the output yielded.

2024/02/24 19:03:12 Data read successfully
2024/02/24 19:03:12 Bootstrap process completed

Profile.prof program was run using the go tool pprof --text profile.prof command, and the output is stored in the text file called cpu_use.txt. From the output, `flat` values are 0, which means that no single function used a significant amount of CPU time directly. However, the `cum` values show that the `main.Bootstrap`, `main.main`, and `runtime.main` functions used 100% of the CPU time, and `main.statFunc`, `math/rand.(*Rand).Intn`, and `math/rand.Intn` used 50% of the CPU time. This suggests that the`main.Bootstrap` function is the main driver of CPU usage in this program, and within that function, `main.statFunc` and the random number generation functions (`math/rand.(*Rand).Intn` and `math/rand.Intn`) are significant contributors.

## Summary to Managment:

The recomendation is to switch to Go for applied statistical computing if the need for scalability and perfromance is of most concern. In addition, R consumes more memory/compute power to perform the same operations as seen in the bootstrap programs. If the company is striving to save on cloud cost, Go will help achieve this objective. For example Google Cloud advertises that you only pay for what you used. If Go is consuming less than 50% of the memory compared to R, that alone will cut costs substantially in short order. From a developer standpoint, the only draw back for the switch is seen with R's robust ecosystem of statistically libraries, however cost savings in cloud computing will be dramitic with the switch to Go. 