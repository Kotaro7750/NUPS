import os
import csv
import shutil

koujipath = "../kouji/"

files = os.listdir(koujipath)
dirs = [f for f in files if os.path.isdir(os.path.join(koujipath, f))]
print(dirs)

#dirの分だけcsvにする
with open('ki.csv', 'w') as file:
    writer = csv.writer(file, lineterminator='\n')
    writer.writerow(dirs)

#dirの分だけテーブル追加
with open('ki_table.sql','w') as f:
    f.write("USE NUPS\n")
    for ki in dirs:
        f.write("CREATE TABLE `"+ ki +"` (id INT NOT NULL,title TEXT NOT NULL);\n")

shutil.copy2("./ki.csv","../db/")
shutil.copy2("./ki_table.sql","../db/")