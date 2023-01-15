import pandas as pd
def update_node():
    s = open("1.txt", "r").readlines()
    v = "".join(s)
    u = v.split()
    for i in u:
        print(i)
    k4 = [str(i) for i in input().split()]
    if ("UPDATE" in k4 and "NODE" in k4 and len(k4) == 3 and k4[2] in u):
        # конвертация графа в таблицу
        df = pd.read_csv(f"{k4[2]}.csv")  # пишу считывание данных для показа
        s = str(input("Индекс:"))
        if(s.isdigit() == True):
            s = int(s)
            s1 = len(df.columns[1])
            if(0 <= s < s1):
                c = str(input("Введите имя столбца:"))
                k1 = list(df.keys())
                if(c in k1):
                    k = str(input("Введите значение:"))
                    df.at[s, c] = k
                    df = df.loc[:, ~df.columns.str.contains('^Unnamed')]
                    df.to_csv(f"{k4[2]}.csv")  # делаешь конвертор таблицы в граф с экранированием и сохранением
                else:
                    print("WRONG NAME!")
            else:
                print("WRONG INDEX")

        else:
            print("----")

    else:
        print("WRONG!")