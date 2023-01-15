import networkx as nx
import matplotlib.pyplot as plt

from handlers import read_csv


class Graph:
    def __init__(self):
        self.G = nx.Graph()
        self.add_nodes()
        self.add_edges()

    def get_headers(self):
        headers = list()
        filename, data = read_csv.read_file()
        for item in data[0]:
            headers.append(item)
        for header in headers:
            self.G.add_edge(filename, header)
        return headers

    @staticmethod
    def get_values():
        values = list()
        _, lst = read_csv.read_file()
        for data in lst:
            for value in data.values():
                values.append(value)
        return values

    def add_nodes(self):
        self.G.add_nodes_from(self.get_headers())
        self.G.add_nodes_from(self.get_values())

    def add_edges(self):
        _, lst = read_csv.read_file()
        for data in lst:
            for item, value in data.items():
                self.G.add_edge(item, value)

    def draw_graph(self):
        nx.draw(self.G, with_labels=True)
        plt.savefig("test.png")

