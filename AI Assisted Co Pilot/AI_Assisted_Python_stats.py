import numpy as np
import pandas as pd
import time
from scipy.stats import linregress

# Load Anscombe quartet data from CSV file
file_path = "Anscombe_quartet_data.csv"
df = pd.read_csv(file_path)

# Print column names and first few rows of the dataframe
print(df.columns)
print(df.head())

# Create X and Y arrays from the dataframe
anscombe_x = np.array([df["x123"].values, df["x123"].values, df["x123"].values, df["x4"].values])
anscombe_y = np.array([df["y1"].values, df["y2"].values, df["y3"].values, df["y4"].values])

# Define a function for performing linear regression using scipy
def regression(x, y):
    slope, intercept, _, _, _ = linregress(x, y)
    return slope, intercept

# Initialize lists to store results and timings
results = []
timings = []

# Iterate over each of the 4 datasets
for i in range(4):
    start_time = time.time()
    slope, intercept = regression(anscombe_x[i], anscombe_y[i])
    exec_time = time.time() - start_time
    
    results.append((slope, intercept))
    timings.append(exec_time)

# Print results
for i, (slope, intercept) in enumerate(results):
    print(f"Dataset {i+1}: Slope = {slope:.5f}, Intercept = {intercept:.5f}")
    print(f"Time = {timings[i]:.5f} seconds")
