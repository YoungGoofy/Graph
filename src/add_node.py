import pandas as pd
from graph import Graph


def add_node():
    s = open("1.txt", "r").readlines()
    v = "".join(s)
    u = v.split()
    for i in u:
        print("Файл: ", i)
    print('Введите команду "ADD NODE <filename>"')
    k2 = [str(i) for i in input().split()]
    if "ADD" in k2 and "NODE" in k2 and len(k2) == 3 and k2[2] in u:
        # конвертация графа в таблицу
        df = pd.read_csv(f"{k2[2]}.csv")  # пишу считывание данных для показа
        if len(df) == 0:
            col = []
            val = [str(i) for i in input("Столбцы:").split()]
            col.extend(val)

            for i in col:
                df[i] = [str(i) for i in input("Значения:").split()]

            df = df.loc[:, ~df.columns.str.contains('^Unnamed')]
            df.to_csv(f"{k2[2]}.csv")
            # конвертор таблицы в граф с экранированием и сохранением
            graph = Graph(k2[2])
            graph.draw_graph()

        else:
            print("----")


    else:
        print("WRONG!")
