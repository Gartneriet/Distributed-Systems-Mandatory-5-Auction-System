go run . --type=server > logs/server1.log 2>&1 &

go run . --type=client > logs/node_halw.log 2>&1 &
go run . --type=client > logs/node_frap.log 2>&1 &
go run . --type=client > logs/node_nivl.log 2>&1 &




wait