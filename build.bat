start /wait "" /d "." cmd /C go build -o main.exe .\cmd\antivirus\antivirus.go
start .\main.exe tests
