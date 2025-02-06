library(dplyr)
library(readr)
library(ggplot2)

# read in csv data and store as a data fram
file_path <- "Anscombe_quartet_data.csv"
df <- read_csv(file_path)

# create x and y matrix
anscombe_x <- list(
  df$x123,
  df$x123,
  df$x123,
  df$x4
)

anscombe_y <- list(
  df$y1,
  df$y2,
  df$y3,
  df$y4
)

# create a function for regression which output slope and intercept
regression <- function(x, y) {
  model <- lm(y ~ x)
  slope <- coef(model)[2]
  intercept <- coef(model)[1]
  return(c(slope, intercept))
}

# create results and timing vars
results <- list()
timings <- numeric(4)

#loop over data and perform regressions
for (i in 1:4) {
  start_time <- Sys.time()
  res <- regression(anscombe_x[[i]], anscombe_y[[i]])
  end_time <- Sys.time()
  
  exec_time <- as.numeric(difftime(end_time, start_time, units = "secs"))
  results[[i]] <- res
  timings[i] <- exec_time
}

# Print results
for (i in 1:4) {
  cat(sprintf("Dataset %d: Slope = %.5f, Intercept = %.5f\n", i, results[[i]][1], results[[i]][2]))
  cat(sprintf("Time = %.5f seconds\n", timings[i]))
}