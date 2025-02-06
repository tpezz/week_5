# Read the dataset
data <- read.csv("Anscombe_quartet_data.csv")

# Perform linear regressions
model1 <- lm(y1 ~ x123, data = data)
model2 <- lm(y2 ~ x123, data = data)
model3 <- lm(y3 ~ x123, data = data)
model4 <- lm(y4 ~ x4, data = data)

# Print summaries of the models
summary(model1)
summary(model2)
summary(model3)
summary(model4)