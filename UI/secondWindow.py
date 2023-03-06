from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_SecondWindow(object):
    def setupUi(self, SecondWindow):
        SecondWindow.setObjectName("SecondWindow")
        SecondWindow.resize(444, 167)
        self.centralwidget = QtWidgets.QWidget(SecondWindow)
        self.centralwidget.setObjectName("centralwidget")

        self.pushButton = QtWidgets.QPushButton(self.centralwidget)
        self.pushButton.setGeometry(QtCore.QRect(320, 110, 111, 41))
        self.pushButton.setObjectName("pushButton")


        self.label = QtWidgets.QLabel(self.centralwidget)
        self.label.setGeometry(QtCore.QRect(80, 40, 281, 51))
        font = QtGui.QFont()
        font.setPointSize(12)
        self.label.setFont(font)
        self.label.setObjectName("label")
        SecondWindow.setCentralWidget(self.centralwidget)
        self.retranslateUi(SecondWindow)
        QtCore.QMetaObject.connectSlotsByName(SecondWindow)

    def retranslateUi(self, SecondWindow):
        _translate = QtCore.QCoreApplication.translate
        SecondWindow.setWindowTitle(_translate("SecondWindow", "Antivirus"))
        self.pushButton.setText(_translate("SecondWindow", "Закрыть"))
        self.label.setText(_translate("SecondWindow", "Антивирус успешно запущен!"))