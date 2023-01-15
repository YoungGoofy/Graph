import csv
import os.path


def read_file(filename: str):
    data = list()
    with open(filename, 'r') as f:
        reader = csv.DictReader(f, delimiter=',')
        for row in reader:
            data.append(row)
        return os.path.basename("test.csv").split('.')[0], data
