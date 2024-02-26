# Project Details:

The bootstrap was the statistical method of choice for this project. The popular Iris dataset was selected to test the statistical analysis comparing R with Go. The bootstrap is performed on the first two columns of the data, which represent sepal width and sepal length. The bootstrap function then conducts 1000 iterations of resampling the data with replacement and calculating various statistics, including correlation of width and length, sepal width median value, sepal length median value, sepal width mean, and sepal length mean. After running both programs, the results yield very similar values for all calculated statistics. The R program was constructed with the assistance of a DataCamp tutorial found here: https://www.datacamp.com/tutorial/bootstrap-r

## Package Details:

The package of choice used for the R program is the `boot` library, included in base R. The `pryr` library was installed to run the memory monitoring processes. Documentation for `pryr` can be found at: [GitHub Repository](https://github.com/hadley/pryr). The bootstrap was built as its own Go program without relying on a third-party package for that function, similar to the operations in the R program. The 'math/rand' package was imported for the random sample in the bootstrap function, and the third-party packages `github.com/gonum/stat` and `github.com/montanaflynn/stats` were used for the statistical calculations. It is worth noting that the search for a bootstrap package on go.dev yielded mixed results. The first two hits provided a package for bootstrapping a server, while the third result discussed implementing bootstrapping logic with lackluster documentation. For the sake of time, I did not delve deeply into selecting the correct bootstrap package but instead integrated the bootstrap function into the program. As demonstrated in the next section, this is one area where CPU could potentially be conserved with a third-party package.

### Memory & CPU Usage:

The R program has memory and time built into the bootstrap program. When run, the output yields

"System time:"
   user  system elapsed
   0.26    0.02    0.39
"Memory usage:"
40.2 MB

In contrast to R, the Go program has a seperate testing and benchmark program built in the `stats_test.go`program to analyize run time and memory usage as well as unit testing. When the benchmark is run the output yields:

5	 220400740 ns/op	18714083 B/op	   25688 allocs/op
PASS
ok  	goStats/appliedStats	2.335s

Memory usage: 18.71 MB and 2.335 secs per iteration. However, when the executable is run using the bash command:

$ start=$(date +%s%N)
./appliedStats.exe

Time taken: 532 ms

This results show that the Go program is comparable in speed to the R program.

Logging was conducted specifically for reading in the data and conducting the bootstrap operation in the main function. Here is what the output yields.

2024/02/24 19:03:12 Data read successfully

2024/02/24 19:03:12 Bootstrap process completed

Profile.prof program was run using the go tool pprof --text profile.prof command, and the output is stored in the text file called cpu_use.txt. From the output, `flat` values are 0, which means that no single function used a significant amount of CPU time directly. However, the `cum` values show that the `main.Bootstrap`, `main.main`, and `runtime.main` functions used 100% of the CPU time, and `main.statFunc`, `math/rand.(*Rand).Intn`, and `math/rand.Intn` used 50% of the CPU time. This suggests that the`main.Bootstrap` function is the main driver of CPU usage in this program, and within that function, `main.statFunc` and the random number generation functions (`math/rand.(*Rand).Intn` and `math/rand.Intn`) are significant contributors.

## Summary to Managment:

The recommendation to management is to switch to Go for applied statistical computing programs, especially if the need for scalability and performance is of most concern. In addition, R consumes more memory/compute power to perform the same operations as seen in the bootstrap programs. If the company is striving to save on cloud cost, Go will help achieve this objective. For example Google Cloud advertises that you only pay for what you used. If Go is consuming less than 50% of the memory compared to R, that alone will cut costs substantially in short order. From a developer standpoint, the only draw back for the switch is seen with R's robust ecosystem of statistically libraries, however cost savings in cloud computing will be dramatic with the switch to Go.