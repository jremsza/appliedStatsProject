library(boot)
library(pryr)

df <- read.csv("iris.csv")

# Measure processing time
timing <- system.time({
  stat_func <- function(data, indices, cor.type){
    dt <- data[indices, ]
    c(
      cor(dt[, 1], dt[, 2], method = cor.type),
      median(dt[, 1]),
      median(dt[, 2]),
      mean(dt[, 1]),
      mean(dt[, 2])
    )
  }

  set.seed(1234)
  bootstrap <- boot(df, stat_func, R=1000, cor.type="pearson")
  print(bootstrap)
})

# Print system time
print("System time:")
print(timing)

print("Memory usage:")
print(mem_used())