import sys, os, subprocess, time, threading
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
from datetime import datetime, timedelta
from mainWindow import *
from secondWindow import *
from thirdWindow import *
from scanWindow import *
from schedWindow import *
from dirMonitoring import *
from PyQt5 import QtCore, QtGui, QtWidgets
from PyQt5.QtWidgets import QFileDialog
from subprocess import Popen, PIPE

os.environ['PYTHONIOENCODING'] = 'utf-8'

def buffer_to_str(buf):
    codec = QtCore.QTextCodec.codecForName("UTF-8")
    return str(codec.toUnicode(buf))

# This process is created only to read stdout of main.exe
class Process(QtCore.QObject):
    stdout = QtCore.pyqtSignal(str)
    stderr = QtCore.pyqtSignal(str)
    finished = QtCore.pyqtSignal(int)
    def start(self, program, args):
        process = QtCore.QProcess()
        process.setProgram(program)
        process.setArguments(args)
        process.readyReadStandardError.connect(lambda: self.stderr.emit(buffer_to_str(process.readAllStandardError())))
        process.readyReadStandardOutput.connect(lambda: self.stderr.emit(buffer_to_str(process.readAllStandardOutput())))
        process.finished.connect(self.finished)
        process.start()
        
        self._process = process

class MainWin(QtWidgets.QMainWindow):
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.ui = Ui_MainWindow()
        self.ui.setupUi(self)
        self.ui.ScanButton.clicked.connect(self.disconnect)
        self.ui.SchedButton.clicked.connect(self.schedule)
        self.ui.MonitorButton.clicked.connect(self.monitor)
        self.ui.toolButton.clicked.connect(self.add_file)

    def disconnect(self):
        if self.ui.lineEdit.text() == '':
            MyApp3.show()
        else:
            self.ui.ServiceStatus.setText("запущен")
            MyApp2.show()
            Scan.show()
            #Scan.scan.textEdit.setPlainText("")
            
            cmd = "../main.exe"
            args = [self.ui.lineEdit.text()]
            Scan.funkStartScan(cmd, args)

    def schedule(self):
        if self.ui.lineEdit_2.text() == '':
            MyApp3.show()
        else:
            self.ui.ServiceStatus.setText("запущен")
            scan_period = self.ui.lineEdit_2.text()
            Schedule.sched.textEdit.setPlainText("")
            Schedule.show()
            Schedule.setScanWindow(scan_period)
            Schedule.makeSchedule(self.ui.lineEdit.text())

    def monitor(self):
        if self.ui.lineEdit.text() == '' or not os.path.isdir(self.ui.lineEdit.text()):
            MyApp3.show()
        else:
            self.ui.ServiceStatus.setText("запущен")
            Monitor.show()
            Monitor.startMonitor(self.ui.lineEdit.text())

    @QtCore.pyqtSlot()
    def add_file(self):
        def getOpenFilesAndDirs(parent=None, caption='', directory='', 
                        filter='', initialFilter='', options=None):
            def updateText():
                # update the contents of the line edit widget with the selected files
                selected = []
                for index in view.selectionModel().selectedRows():
                    selected.append('"{}"'.format(index.data()))
                lineEdit.setText(' '.join(selected))

            dialog = QtWidgets.QFileDialog(parent, windowTitle=caption)
            dialog.setFileMode(dialog.ExistingFiles)
            if options:
                dialog.setOptions(options)
            dialog.setOption(dialog.DontUseNativeDialog, True)
            if directory:
                dialog.setDirectory(directory)
            if filter: 
                dialog.setNameFilter(filter)
                if initialFilter:
                    dialog.selectNameFilter(initialFilter)

            # by default, if a directory is opened in file listing mode, 
            # QFileDialog.accept() shows the contents of that directory, but we 
            # need to be able to "open" directories as we can do with files, so we 
            # just override accept() with the default QDialog implementation which 
            # will just return exec_()
            dialog.accept = lambda: QtWidgets.QDialog.accept(dialog)

            # there are many item views in a non-native dialog, but the ones displaying 
            # the actual contents are created inside a QStackedWidget; they are a 
            # QTreeView and a QListView, and the tree is only used when the 
            # viewMode is set to QFileDialog.Details, which is not this case
            stackedWidget = dialog.findChild(QtWidgets.QStackedWidget)
            view = stackedWidget.findChild(QtWidgets.QListView)
            view.selectionModel().selectionChanged.connect(updateText)
            lineEdit = dialog.findChild(QtWidgets.QLineEdit)
            # clear the line edit contents whenever the current directory changes
            dialog.directoryEntered.connect(lambda: lineEdit.setText(''))

            dialog.exec_()
            return dialog.selectedFiles()
        fname = getOpenFilesAndDirs(self, "Open file or directory", "", "")[0]
        
        if fname:
            self.ui.lineEdit.setText(fname)

class SecondWin(QtWidgets.QMainWindow):
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.second = Ui_SecondWindow()
        self.second.setupUi(self)
        self.second.pushButton.clicked.connect(self.funk1)
    
    def funk1(self):
        MyApp2.close()
        MyApp.show()


class ThirdWin(QtWidgets.QMainWindow):
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.third = Ui_TrirdWindow()
        self.third.setupUi(self)
        self.third.pushButton_2.clicked.connect(self.funk2)

    def funk2(self):
        MyApp3.close()
        MyApp.show()

class ScanWin(QtWidgets.QMainWindow):
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.scan = Ui_ScanWindow()
        self.scan.setupUi(self)
        self.scan.btnClose.clicked.connect(self.funkClose)
        self.scan.btnCancelScan.clicked.connect(self.funkCancelScan)
        self.process = None
        
    def funkClose(self):
        Scan.close()
        self.scan.txtScan.setPlainText("")
        MyApp.ui.ServiceStatus.setText("не запущен")
        
    def funkCancelScan(self):
        assert self.process != None
        # Признаю, это костыль
        try:
            self.process.stderr.disconnect(self.scan.txtScan.appendPlainText)
        except:
            pass
            
    def funkStartScan(self, cmd, args):
        self.process = Process()
        self.process.stderr.connect(self.scan.txtScan.appendPlainText)
        self.process.start(cmd, args)
        
    def closeEvent(self, event):
        assert self.process != None
        # Признаю, это костыль
        try:
            self.process.stderr.disconnect(self.scan.txtScan.appendPlainText)
        except:
            pass
        self.scan.txtScan.setPlainText("")
        MyApp.ui.ServiceStatus.setText("не запущен")

class SchedWindow(QtWidgets.QMainWindow): # текстбокс с логами называется textEdit
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.sched = Ui_SchedWindow()
        self.sched.setupUi(self)
        self.sched.CancelScan.clicked.connect(self.exitSched)
        self.path = None
        self.timePeriod = None

    def setScanWindow(self, timePeriod):
        self.timePeriod = int(timePeriod)
        self.sched.ScanTime.setText(str((datetime.now()+timedelta(seconds=self.timePeriod)).hour)+":"+str((datetime.now()+timedelta(seconds=self.timePeriod)).minute))
    
    def makeSchedule(self, path):
        self.path = path
        if self.path == '':
            MyApp3.show()
        else:
            self.schedScanner = SchedScanner(self.path, self.timePeriod)
            self.schedScanner.stdout.connect(self.sched.textEdit.appendPlainText)
            self.schedScanner.nextScan.connect(self.sched.ScanTime.setText)
            self.schedScanner.start()

    def exitSched(self):
        self.schedScanner.terminate()
        Schedule.close()
        self.sched.textEdit.setPlainText("")
        MyApp.ui.ServiceStatus.setText("не запущен")
        
    def closeEvent(self, event):
        self.schedScanner.terminate()
        self.sched.textEdit.setPlainText("")
        MyApp.ui.ServiceStatus.setText("не запущен")

class SchedScanner(QtCore.QThread):
    stdout = QtCore.pyqtSignal(str)
    nextScan = QtCore.pyqtSignal(str)
    def __init__(self, path, period):
        super().__init__()
        self.path = path
        self.period = period
    def run(self):
        while True:
            time.sleep(self.period)
            with Popen(["../main.exe", self.path], stdout=PIPE) as p:
                while True:
                    text = p.stdout.read1().decode("utf-8")
                    time.sleep(0.1)
                    
                    # Если сервис за 0.1 секунду ничего не вывел - значит прекращаем считывание
                    if text == "":
                        break

            
                    self.stdout.emit(text)
            scanTime = datetime.now()+timedelta(seconds=self.period)
            self.nextScan.emit(str(scanTime.hour)+":"+str(scanTime.minute))

class DirMonitoring(QtWidgets.QMainWindow): # текстбокс с логами называется textEdit
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.mon = Ui_DirMonitoring()
        self.mon.setupUi(self)
        self.mon.StopMonitor.clicked.connect(self.exitMon)
        self.path = None
        self.last_trigger = time.time()

    def startMonitor(self, path):
        self.path = path
        if self.path == '':
            MyApp3.show()
        else:
            self.mon.DirPath.setText(self.path)
            self.event_handler = MonHandler()
            self.observer = Observer()
            self.observer.schedule(self.event_handler, path=self.path, recursive=True)
            self.observer.start()

    def exitMon(self):
        Monitor.close()
        self.observer.stop()
        self.mon.textEdit.setPlainText("")
        MyApp.ui.ServiceStatus.setText("не запущен")
        
    def closeEvent(self, event):
        self.observer.stop()
        self.mon.textEdit.setPlainText("")
        MyApp.ui.ServiceStatus.setText("не запущен")

class MonHandler(FileSystemEventHandler):
    def on_created(self, event):
        if not os.path.isdir(event.src_path) and (time.time() - Monitor.last_trigger) > 1:
            Monitor.last_trigger = time.time()
            Monitor.mon.textEdit.appendPlainText("В директории создан новый файл: "+event.src_path+"\nЗапущено сканирование всей директории...")
            cmd = "../main.exe "+event.src_path
            p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            stdout, stderr = p.communicate()
            Monitor.mon.textEdit.appendPlainText(str(stdout.decode()+stderr.decode()))

    def on_modified(self, event):
        if not os.path.isdir(event.src_path) and (time.time() - Monitor.last_trigger) > 1:
            Monitor.last_trigger = time.time()
            Monitor.mon.textEdit.appendPlainText("В директории изменён файл: "+event.src_path+"\nЗапущено сканирование всей директории...")
            cmd = "../main.exe "+event.src_path
            p = subprocess.Popen(cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
            stdout, stderr = p.communicate()
            Monitor.mon.textEdit.appendPlainText(str(stdout.decode()+stderr.decode()))

if __name__ == "__main__":
    app = QtWidgets.QApplication(sys.argv)
    MyApp = MainWin()
    MyApp2 = SecondWin()
    MyApp3 = ThirdWin()
    Scan = ScanWin()
    Schedule = SchedWindow()
    Monitor = DirMonitoring()
    MyApp.show()
    sys.exit(app.exec_())
