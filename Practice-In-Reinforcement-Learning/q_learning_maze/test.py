import numpy as np
import tkinter as tk
import time

#
h = 4
w = 4
unit = 80

tk1 = tk.Tk()
tk1.title("test")
tk1.geometry('320x320')

canvas = tk.Canvas(tk1, bg='green', height=320, width=320)
canvas.pack()

for c in range(0, 320, 80):
    x0, y0, x1, y1 = c, 0, c, 320
    canvas.create_line(x0, y0, x1, y1)
for r in range(0, 320, 80):
    x0, y0, x1, y1 = 0, r, 320, r
    canvas.create_line(x0, y0, x1, y1)

origin = np.array([40, 40])

hell_center1 = origin + np.array([160, 80])
rec = canvas.create_rectangle(hell_center1[0] - 35, hell_center1[1] - 35,
                              hell_center1[0] + 35, hell_center1[1] + 35, fill="black")
print(canvas.coords(rec))

oval_center = origin + np.array([160, 160])
oval = canvas.create_oval(oval_center[0] - 35, oval_center[1] - 35,
                          oval_center[0] + 35, oval_center[1] + 35, fill="yellow")

rec1 = canvas.create_rectangle(
    origin[0] - 35, origin[1] - 35,
    origin[0] + 35, origin[1] + 35,
    fill='red'
)

tk1.update()
canvas.delete(rec1)
time.sleep(1)
tk1.update()
time.sleep(1)
canvas.move(rec, 80, 0)
tk1.update()
time.sleep(5)
tk1.update()

