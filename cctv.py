import cv2
from datetime import datetime
from time import sleep
import os

cap = cv2.VideoCapture(0) # video capture source camera (Here webcam of laptop) 
cap.set(3,320) # set Width
cap.set(4,240) # set Height
counter = 0
start = "{}-{}-{}__{}-{}-{}".format(datetime.now().day, datetime.now().month, datetime.now().year, datetime.now().hour, datetime.now().minute, datetime.now().second)
watermark1 = "OpenCV: KualaSempang/Apotek"

os.mkdir(start)
os.chdir(start)

 
try:
    while True:
        ret,frame = cap.read()
        #now_time = datetime.now()
        filename = "img{0:05d}.jpg".format(counter)
        now = datetime.now()
        timestamp = "Timestamp: {}-{}-{} {}:{}:{}".format(now.day, now.month, now.year, now.hour, now.minute, now.second)
        cv2.putText(frame, timestamp, (0,10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 255, 0), 1, cv2.LINE_AA)
        cv2.putText(frame, watermark1, (0,238), cv2.FONT_HERSHEY_COMPLEX, 0.3, (0, 255, 0), 1, cv2.LINE_AA)
        #cv2.putText(frame, watermark2, (200,238), cv2.FONT_HERSHEY_COMPLEX, 0.3, (0, 255, 0), 1, cv2.LINE_AA)
        cv2.imwrite(filename, frame)
        sleep(1.5)
        counter += 1
except KeyboardInterrupt:
    print("monitor selesai")

cap.release()