import time
import threading
from queue import Queue
import os
from urllib import request


def target(q):
    while True:
        msg = q.get()
        for i in range(1):
            print('running thread-{}:{}'.format(threading.get_ident(), i))
            time.sleep(1)
            q.task_done()


def pool(workers, queue):
    for n in range(workers):
        t = threading.Thread(target=target, args=(queue,))
        t.daemon = True
        t.start()


if __name__ == '__main__':
    queue = Queue()
    # 创建一个线程池：并设置线程数为5
    pool(5, queue)

    for i in range(12):
        queue.put("start")

    # 消息都被消费才能结束
    queue.join()
