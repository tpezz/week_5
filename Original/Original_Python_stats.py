import numpy as np
import pandas as pd
import time
from scipy.stats import linregress

#Pull in Kaggle Anscombe quartetddata
file_path = "Anscombe_quartet_data.csv"
df = pd.read_csv(file_path)

#Print out results from the file
print(df.columns)
print(df.head())

#Create X and Y arrays
anscombe_x = np.array([
    df["x123"].tolist(), 
    df["x123"].tolist(),
    df["x123"].tolist(),
    df["x4"].tolist()  
])

anscombe_y = np.array([
    df["y1"].tolist(),  
    df["y2"].tolist(),  
    df["y3"].tolist(),  
    df["y4"].tolist()   
])

#create function for Scipy regression
def regression(x, y):
    slope, intercept, r_value, p_value, std_err = linregress(x, y)
    return slope, intercept

#Create for loop over each of the 4 datasets
#Ensure timing is set up
results = []
timings = []
for i in range(4):
    start_time = time.time()
    slope, intercept = regression(anscombe_x[i], anscombe_y[i])
    end_time = time.time()
    
    exec_time = end_time - start_time
    results.append((slope, intercept))
    timings.append(exec_time)

#print results
for i, (slope, intercept) in enumerate(results):
    print(f'''Dataset {i+1}: Slope = {slope:.5f}, Intercept = {intercept:.5f}''')
    print(f'''Time = {timings[i]:.5f}''')
