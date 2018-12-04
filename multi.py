from threading import Thread,Lock,Event
import time,random;

mutex = Lock();
value = 0
loop = True
first= True
answ=''

def thread(data,sec,evst,evsp):
    global value
    print('Поток '+str(data) + ' начался')
    evst.wait()
    evst.clear()
    evsp.set()
    time.sleep(sec)
    mutex.acquire()
    print('Поток '+str(data) + ' проспал: '+str(sec)+' секунд')
    if random.randint(1,3) == 1:
        value+=1
    else:
        value-=1
    print('Поток '+str(data) + ' завершился с результатом:' + str(value))
    mutex.release()

while loop:
    if first:
        n = int(input("Введите количество потоков: "))
        first=False
    else:
        answ = input("Вы хотите продолжить работу с потоками(y/n)? ")
        if answ=='n':
            loop=False
        else:
            n = int(input("Введите количество потоков: "))
    start = time.time()
    events = []
    threads = []
    for i in range (0,n):
        e=Event()
        events.append(e)
    for i in range (0,n-1):
        t = Thread(target = thread, args = (i+1,random.randint(1,5),events[i],events[i+1]))
        threads.append(t)
        threads[i].start()
    t = Thread(target = thread, args = (n,random.randint(1,5),events[n-1],events[0]))
    threads.append(t)
    threads[n-1].start()
    events[0].set()
    for i in range (0,n):
        threads[i].join()
    stop = time.time()
    print ("Время работы потоков: "+str(stop-start)+ " секунд")
