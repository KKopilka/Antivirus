import sys, os, subprocess
from mainWindow import *
from secondWindow import *
from thirdWindow import *
from scanWindow import *
from schedWindow import *
from PyQt5 import QtCore, QtGui, QtWidgets
from PyQt5.QtWidgets import QFileDialog

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
        self.ui.toolButton.clicked.connect(self.add_file)

    def disconnect(self):
        if self.ui.lineEdit.text() == '':
            MyApp3.show()
        else:
            MyApp2.show()
            Scan.show()
            
            cmd = "../main.exe"
            args = [self.ui.lineEdit.text()]
            Scan.funkStartScan(cmd, args)

    def schedule(self):
        if self.ui.lineEdit_2.text() == '':
            MyApp3.show()
        else:
    	    Schedule.show()

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
            print(fname)

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
        self.scan.txtScan.setPlainText("")
        Scan.close()
        
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

class SchedWindow(QtWidgets.QMainWindow): # текстбокс с логами называется textEdit
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.sched = Ui_SchedWindow()
        self.sched.setupUi(self)
        self.sched.CancelScan.clicked.connect(self.exitSched)
        # self.sched.ScanTime.setText('Время вместо 13:00')
    
    def exitSched(self):
        Schedule.close()
        # отмена сканирования

if __name__ == "__main__":
    app = QtWidgets.QApplication(sys.argv)
    MyApp = MainWin()
    MyApp2 = SecondWin()
    MyApp3 = ThirdWin()
    Scan = ScanWin()
    Schedule = SchedWindow()
    MyApp.show()
    sys.exit(app.exec_())
