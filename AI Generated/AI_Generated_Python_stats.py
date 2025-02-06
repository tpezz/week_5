import pandas as pd
import scipy.stats as stats

df = pd.read_csv("Anscombe_quartet_data.csv")

for i in range(1, 5):
    x = df["x123"] if i < 4 else df["x4"]
    y = df[f"y{i}"]
    slope, intercept, r_value, _, _ = stats.linregress(x, y)
    print(f"y{i}: Slope={slope:.2f}, Intercept={intercept:.2f}, R-squared={r_value**2:.2f}")
