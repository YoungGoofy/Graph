import csv
import os.path


def read_file(filename: str):
    data = list()
    with open(f'{filename}.csv', 'r') as f:
        reader = csv.DictReader(f, delimiter=',')
        for row in reader:
            data.append(row)
        return filename, data
