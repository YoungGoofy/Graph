import pandas as pd
from graph import Graph


def create_node():
    k1 = [str(i) for i in input().split()]
    if "CREATE" in k1 and "NODE" in k1 and len(k1) == 3:
        df = pd.DataFrame({})
        with open("1.txt", "a+") as f:
            f.write(f'{k1[2]}\n')
        df.to_csv(f"{k1[2]}.csv")
        # перевести наименование таблицы в вершину графа с показом и сохранением
        graph = Graph(k1[2])
        graph.add_node()

    else:
        print("WRONG!")
