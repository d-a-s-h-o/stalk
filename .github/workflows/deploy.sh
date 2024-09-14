echo Started
scp -o StrictHostKeyChecking=no -r -P 4242 *.yml *.go plugin talkapi go.sum go.mod ubuntu@150.136.142.44:~/talk
echo Copied files
ssh -o StrictHostKeyChecking=no -p 4242 ubuntu@150.136.142.44 <<EOL # Unquote so lines are expanded
	cd ~/talk
	go build -ldflags "-X 'main.unameCommit=$(git rev-parse HEAD)' -X 'main.unameTime=$(date)'" && echo Built
	echo $SERVER_PASS | sudo -S pkill talk && echo Killed
	sleep 2
	echo $SERVER_PASS | sudo -S pkill -9 talk && echo Killed with SIGKILL
	echo $SERVER_PASS | nohup sudo -S GOMAXPROCS=2 talk_CONFIG=mainserver.yml ./talk > /dev/null 2>stderr </dev/null &
	echo Started server
	disown
	exit
EOL
echo Finished
