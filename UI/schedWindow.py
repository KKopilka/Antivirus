# -*- coding: utf-8 -*-

# Form implementation generated from reading ui file 'SchedWindow.ui'
#
# Created by: PyQt5 UI code generator 5.15.9
#
# WARNING: Any manual changes made to this file will be lost when pyuic5 is
# run again.  Do not edit this file unless you know what you are doing.


from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_SchedWindow(object):
    def setupUi(self, SchedWindow):
        SchedWindow.setObjectName("SchedWindow")
        SchedWindow.resize(539, 561)
        self.centralwidget = QtWidgets.QWidget(SchedWindow)
        self.centralwidget.setObjectName("centralwidget")
        self.CancelScan = QtWidgets.QPushButton(self.centralwidget)
        self.CancelScan.setGeometry(QtCore.QRect(180, 510, 181, 41))
        self.CancelScan.setObjectName("CancelScan")
        self.label = QtWidgets.QLabel(self.centralwidget)
        self.label.setGeometry(QtCore.QRect(90, 0, 331, 51))
        font = QtGui.QFont()
        font.setPointSize(12)
        self.label.setFont(font)
        self.label.setObjectName("label")
        self.label_2 = QtWidgets.QLabel(self.centralwidget)
        self.label_2.setGeometry(QtCore.QRect(90, 40, 331, 51))
        font = QtGui.QFont()
        font.setPointSize(12)
        self.label_2.setFont(font)
        self.label_2.setObjectName("label_2")
        self.ScanTime = QtWidgets.QLabel(self.centralwidget)
        self.ScanTime.setGeometry(QtCore.QRect(370, 40, 111, 51))
        font = QtGui.QFont()
        font.setPointSize(12)
        self.ScanTime.setFont(font)
        self.ScanTime.setObjectName("ScanTime")
        self.textEdit = QtWidgets.QTextEdit(self.centralwidget)
        self.textEdit.setGeometry(QtCore.QRect(10, 110, 521, 391))
        self.textEdit.setReadOnly(True)
        self.textEdit.setObjectName("textEdit")
        self.label_3 = QtWidgets.QLabel(self.centralwidget)
        self.label_3.setGeometry(QtCore.QRect(10, 70, 331, 51))
        font = QtGui.QFont()
        font.setPointSize(12)
        self.label_3.setFont(font)
        self.label_3.setObjectName("label_3")
        SchedWindow.setCentralWidget(self.centralwidget)

        self.retranslateUi(SchedWindow)
        QtCore.QMetaObject.connectSlotsByName(SchedWindow)

    def retranslateUi(self, SchedWindow):
        _translate = QtCore.QCoreApplication.translate
        SchedWindow.setWindowTitle(_translate("SchedWindow", "Antivirus"))
        self.CancelScan.setText(_translate("SchedWindow", "Отменить сканирование"))
        self.label.setText(_translate("SchedWindow", "Запущено сканирование по расписанию"))
        self.label_2.setText(_translate("SchedWindow", "Время следующего сканирования:"))
        self.ScanTime.setText(_translate("SchedWindow", "13:00"))
        self.label_3.setText(_translate("SchedWindow", "Output"))
