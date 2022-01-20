import threading
import backend.telegram_bot
import backend.app
import subprocess
from backend.app import db, add_admin
from config import HOST_PORT

class FlaskThread(threading.Thread):
    def run(self) -> None:
        backend.app.app.run(port=HOST_PORT)


class TelegramBotThread(threading.Thread):
    def run(self) -> None:
        backend.telegram_bot.main()


class ReactThread(threading.Thread):
    def run(self) -> None:
        subprocess.check_call('npm start --scripts-prepend-node-path=auto', shell=True)


if __name__ == '__main__':
    print("start...")
    # create all tables
    db.create_all()

    # create initial admin
    add_admin(username="admin", password="236369")
    flask_thread = FlaskThread()
    flask_thread.start()
    # telegram_bot_thread = TelegramBotThread()
    # telegram_bot_thread.start()
    react_thread = ReactThread()
    react_thread.start()
    backend.telegram_bot.main()




'''import os
from backend.app import db, add_admin
import subprocess
from multiprocessing import Process, Pool


# This block of code enables us to call the script from command line.
def execute(process):
    os.system(f'python {process}')


def run_backend():
    os.system('python backend/app.py')


def run_telegram_bot():
    os.system('python backend/telegram_bot.py')


def run_npm():
    subprocess.check_call('npm start --scripts-prepend-node-path=auto')

if __name__ == '__main__':
    # create all tables
    db.create_all()

    # create initial admin
    add_admin(username="admin", password="236369")

    # Creating the tuple of all the processes
    all_processes = ('backend/app.py', 'backend/telegram_bot.py')

    process_pool = Pool(processes=2)
    process_pool.map(execute, all_processes)

    p1 = Process(target=run_backend(), args=())
    p2 = Process(target=run_telegram_bot(), args=())
    p3 = Process(target=run_npm(), args=())

    p1.start()
    p2.start()
    p3.start()

    p1.join()
    p2.join()
    p3.join()'''

