from PyQt5 import QtCore, QtGui, QtWidgets

class Ui_TrirdWindow(object):
    def setupUi(self, TrirdWindow):
        TrirdWindow.setObjectName("TrirdWindow")
        TrirdWindow.resize(444, 167)
        self.centralwidget = QtWidgets.QWidget(TrirdWindow)
        self.centralwidget.setObjectName("centralwidget")
        self.label = QtWidgets.QLabel(self.centralwidget)
        self.label.setGeometry(QtCore.QRect(90, 10, 271, 51))
        font = QtGui.QFont()
        font.setPointSize(12)
        self.label.setFont(font)
        self.label.setObjectName("label")
        self.label_2 = QtWidgets.QLabel(self.centralwidget)
        self.label_2.setGeometry(QtCore.QRect(130, 50, 201, 51))
        font = QtGui.QFont()
        font.setPointSize(12)
        self.label_2.setFont(font)
        self.label_2.setObjectName("label_2")

        self.pushButton_2 = QtWidgets.QPushButton(self.centralwidget)
        self.pushButton_2.setGeometry(QtCore.QRect(320, 110, 111, 41))
        self.pushButton_2.setObjectName("pushButton_2")
        

        TrirdWindow.setCentralWidget(self.centralwidget)

        self.retranslateUi(TrirdWindow)
        QtCore.QMetaObject.connectSlotsByName(TrirdWindow)

    def retranslateUi(self, TrirdWindow):
        _translate = QtCore.QCoreApplication.translate
        TrirdWindow.setWindowTitle(_translate("TrirdWindow", "Antivirus"))
        self.label.setText(_translate("TrirdWindow", "Ошибка запуска антивируса."))
        self.label_2.setText(_translate("TrirdWindow", "Попробуйте еще раз."))
        self.pushButton_2.setText(_translate("TrirdWindow", "Закрыть"))

    # def on_click(self):
    #     TrirdWindow.close()
    #     object_mainWindow = Ui_MainWindow()
    #     object_mainWindow.show()
