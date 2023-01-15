import os

import networkx as nx
import matplotlib.pyplot as plt

from handlers import read_csv


class Graph:
    def __init__(self, filename):
        self.filename = filename
        self.G = nx.Graph()

    def get_headers(self):
        headers = list()
        filename, data = read_csv.read_file(self.filename)
        for item in data[0]:
            headers.append(item)
        for header in headers:
            self.G.add_edge(filename, header)
        return headers

    def get_values(self):
        values = list()
        _, lst = read_csv.read_file(self.filename)
        for data in lst:
            for value in data.values():
                values.append(value)
        return values

    def add_nodes_from_table(self):
        self.G.add_nodes_from(self.get_headers())
        self.G.add_nodes_from(self.get_values())

    def add_edges_from_table(self):
        _, lst = read_csv.read_file(self.filename)
        for data in lst:
            for item, value in data.items():
                self.G.add_edge(item, value)

    def draw_graph(self):
        self.add_nodes_from_table()
        self.add_edges_from_table()
        nx.draw(self.G, with_labels=True)
        plt.savefig(f"{os.path.basename('test.csv').split('.')[0]}.png")
