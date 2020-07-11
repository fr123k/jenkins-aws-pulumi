#!/bin/bash

MAX_ATTEMPS=10
JOB_URL=http://admin:$3@$1/job/$2
JOB_STATUS_URL=${JOB_URL}/lastBuild/api/json
JOB_CONSOLE_URL=${JOB_URL}/lastBuild/consoleText

GREP_RETURN_CODE=0

# Start the build
echo "curl -X POST $JOB_URL/build?delay=0sec -v"
curl -X POST $JOB_URL/build?delay=0sec -vvv

# Poll every 7seconds until the build is finished
while [ $GREP_RETURN_CODE -eq 0 ]
do
    ((ATTEMPS++))
    echo "Wait for jenkins job $1 to finish tries=$ATTEMPS."
    sleep $4
    # Grep will return 0 while the build is running:
    curl -s $JOB_STATUS_URL | grep result\":null > /dev/null 
    #|| if [ "$?" == "1" ];then
    GREP_RETURN_CODE=$?
    if [[ $ATTEMPS = $MAX_ATTEMPS ]]; then
        echo "Reached MAX_ATTEMPS $MAX_ATTEMPS==$ATTEMPS."
        break
    fi
done

curl -s $JOB_CONSOLE_URL
