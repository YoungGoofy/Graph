import pandas as pd

import add_node
import create
import delete
import upd

n = 1


while(n != 0):
    m = ["Создать граф(нажмите 1)","добавить ноды(нажмите 2)","удалить ноды(нажмите 3)","обновить ноды(нажмите 4)"]
    for i in m:
        print(i)

    s = str(input())
    if(s == "1"):
        create.create_node()
    elif(s == "2"):
        add_node.add_node()
    elif (s == "3"):
        delete.del_node()
    elif (s == "4"):
        upd.update_node()
    else:
        print("???")
    ans = ["да", "нет"]
    ans_n = "/".join(ans)
    print(f"Хотите ли вы продолжить работу с данным ПО?{ans_n}")
    s = str(input("Введите ответ:"))
    if (s == ans[0]):
        n += 1
    elif (s == ans[1]):
        break
    else:
        print("Ничего непонятно")