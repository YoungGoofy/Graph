import csv


def read_file():
    data = list()
    with open("test.csv", 'r') as f:
        reader =csv.DictReader(f, delimiter=',')
        for row in reader:
            data.append(row)
        return data

