To run this on a Windows terminal use these commands in three different terminals:

``go run . --type=server --port=:5050 --name=server1 > logs/server1.log 2>&1``

``go run . --type=server --port=:6000 --name=server2 > logs/server2.log 2>&1``

``go run . --type=client --name=halw > logs/node_halw.log 2>&1``

``go run . --type=client --name=frap > logs/node_frap.log 2>&1``

``go run . --type=client --name=nivl > logs/node_nivl.log 2>&1``

If you have a UNIX terminal, Linux, Mac or Git, you can run the script run_nodes.sh in the scripts folder with the commands:

``chmod +x scripts/run_all.sh``

``./scripts/run_all.sh``

All logs generated for the program can be seen in the logs folder
