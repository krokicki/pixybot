import web
import os
import time
import threading
import multiprocessing
from Queue import Queue

urls = (
    '/say', 'say'
)
app = web.application(urls, globals())

def sayText(text):
    print "Saying \"%s\"" % text
    os.system("say -v Victoria \"%s\"" % text)

def run():
    while True:
        text = web.q.get()
        # Drain the queue
        while not(web.q.empty()):
            text = web.q.get()
        if text:
            sayText(text)

class say:
    def GET(self):
        user_data = web.input(text="")
        text = user_data.text.strip()
        if not(text): return 'Specify what to say by passing ?text=Your Saying'
        print "Request to say '%s'" % text
        web.q.put(text)
        return "OK"

if __name__ == "__main__":
    web.q = Queue()
    p = threading.Thread(target=run)
    p.daemon = True
    p.start()
    app.run()

