import networkx as nx
from handlers import read_csv


def add_nodes():
    G = nx.Graph()


def get_headers():
    headers = list()
    for item in read_csv[0].items():
        headers.append(item)
    print(headers)
