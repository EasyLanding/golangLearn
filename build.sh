echo "Compiling started..."
go build -o main
echo "Compiling complete."

echo "Trying to launch program."
./main
echo "Program exited."

echo "Checking Golang version..."
go version