go run . --type=server --port=:5050 --name=server1 > logs/server1.log 2>&1 &
go run . --type=server --port=:6000 --name=server2 > logs/server2.log 2>&1 &
go run . --type=client --name=halw > logs/node_halw.log 2>&1 &
go run . --type=client --name=frap > logs/node_frap.log 2>&1 &
go run . --type=client --name=nivl > logs/node_nivl.log 2>&1 &




wait