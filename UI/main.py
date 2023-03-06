import sys
from mainWindow import *
from secondWindow import *
from thirdWindow import *
from scanWindow import *
from PyQt5 import QtCore, QtGui, QtWidgets
from PyQt5.QtWidgets import QFileDialog


class MainWin(QtWidgets.QMainWindow):
    def __init__(self, parent=None):
        QtWidgets.QWidget.__init__(self, parent)
        self.ui = Ui_MainWindow()
        self.ui.setupUi(self)
        self.ui.ScanButton.clicked.connect(self.disconnect)
        self.ui.toolButton.clicked.connect(self.add_file)

    def disconnect(self):
        if self.ui.lineEdit.text() == '':
            MyApp3.show()
        else:
            MyApp2.show()
            Scan.show()


    @QtCore.pyqtSlot()
    def add_file(self):
        fname, filetype = QFileDialog.getOpenFileName(
            self, 
            "Open file",
            ".", 
            "Zip Files (*.zip)"
        )
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

    def funkClose(self):
        Scan.close()

if __name__ == "__main__":
    app = QtWidgets.QApplication(sys.argv)
    MyApp = MainWin()
    MyApp2 = SecondWin()
    MyApp3 = ThirdWin()
    Scan = ScanWin()
    MyApp.show()
    sys.exit(app.exec_())