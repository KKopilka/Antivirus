from PyQt5 import QtCore, QtGui, QtWidgets


class Ui_ScanWindow(object):
    def setupUi(self, ScanWindow):
        ScanWindow.setObjectName("ScanWindow")
        ScanWindow.resize(757, 423)
        sizePolicy = QtWidgets.QSizePolicy(QtWidgets.QSizePolicy.Fixed, QtWidgets.QSizePolicy.Fixed)
        sizePolicy.setHorizontalStretch(0)
        sizePolicy.setVerticalStretch(0)
        sizePolicy.setHeightForWidth(ScanWindow.sizePolicy().hasHeightForWidth())
        ScanWindow.setSizePolicy(sizePolicy)
        ScanWindow.setMinimumSize(QtCore.QSize(757, 423))
        ScanWindow.setMaximumSize(QtCore.QSize(757, 423))
        icon = QtGui.QIcon()
        icon.addPixmap(QtGui.QPixmap(":/clamguard/Gerald-G-Clam-Security-Guard-ico32x32.png"), QtGui.QIcon.Normal, QtGui.QIcon.Off)
        ScanWindow.setWindowIcon(icon)
        self.centralwidget = QtWidgets.QWidget(ScanWindow)
        self.centralwidget.setObjectName("centralwidget")
        #self.txtScan = QtWidgets.QTextEdit(self.centralwidget)
        self.txtScan = QtWidgets.QPlainTextEdit(self.centralwidget)
        self.txtScan.setReadOnly(True)
        self.txtScan.setGeometry(QtCore.QRect(10, 30, 741, 341))
        self.txtScan.setObjectName("txtScan")
        self.lblOutput = QtWidgets.QLabel(self.centralwidget)
        self.lblOutput.setGeometry(QtCore.QRect(10, 10, 54, 17))
        self.lblOutput.setObjectName("lblOutput")
        self.btnCancelScan = QtWidgets.QPushButton(self.centralwidget)
        self.btnCancelScan.setGeometry(QtCore.QRect(280, 380, 121, 29))
        self.btnCancelScan.setObjectName("btnCancelScan")
        self.btnClose = QtWidgets.QPushButton(self.centralwidget)
        self.btnClose.setGeometry(QtCore.QRect(410, 380, 87, 29))
        self.btnClose.setObjectName("btnClose")
        ScanWindow.setCentralWidget(self.centralwidget)

        self.retranslateUi(ScanWindow)
        QtCore.QMetaObject.connectSlotsByName(ScanWindow)

    def retranslateUi(self, ScanWindow):
        _translate = QtCore.QCoreApplication.translate
        ScanWindow.setWindowTitle(_translate("ScanWindow", "Scan for viruses"))
        self.lblOutput.setText(_translate("ScanWindow", "Output"))
        self.btnCancelScan.setText(_translate("ScanWindow", "Остановить"))
        self.btnClose.setText(_translate("ScanWindow", "Закрыть"))
