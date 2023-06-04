import matplotlib.pyplot as plt
import pandas as pd
import numpy as np
import glob

# Get a list of all CSV files
csv_files = sorted(glob.glob('csv/frame_*.csv'))

for i, csv_file in enumerate(csv_files):
    # Read the CSV file into a DataFrame
    df = pd.read_csv(csv_file, header=None)

    # Convert the DataFrame values to float
    data = df.values.astype(float)

    # Plot the data
    plt.imshow(data, cmap='hot', interpolation='nearest')

    # Save the plot to a PNG file
    plt.savefig('images/frame_{:04d}.png'.format(i))
