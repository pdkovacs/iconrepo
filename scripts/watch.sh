#!/bin/bash

export ICON_REPO_CONFIG_FILE=deployments/dev/app-configs/dev-oidc-proxy-gitlab.json
cmd="make app"
settle_down_secs=1

app_executable="igo-repo"
app_instance_count=2

logs_home=~/workspace/logs
mkdir -p $logs_home
webpack_log=$logs_home/igonrepo-webpack-build
app_log=$logs_home/iconrepo-app-

project_dir="$(dirname $0)/.."
. "$project_dir/scripts/functions.sh"

start_app() {
  set -x
  for i in $(seq 0 $((app_instance_count -1)));
  do
    export SERVER_PORT=$((8091 + $i))
    export LOAD_BALANCER_ADDRESS=$(get_my_ip):9999
    ./"$app_executable" -l debug >"$app_log$i" 2>&1 &
  done
  set +x
}

pkill webpack
pkill "$app_executable"

tail_command_pattern="[t]ail.*${webpack_log}"
ps -ef | awk -v tail_command_pattern="$tail_command_pattern" '$0 ~ tail_command_pattern { print $0; system("kill " $2); }'

unset stopping

cleanup() {
  echo "Cleaning up $$..."
  stopping=true
  ps -ef | awk -v tail_command_pattern="$tail_command_pattern" '$0 ~ tail_command_pattern { print $0; system("kill " $2); }'
  pkill -9 -P $$
}

trap cleanup EXIT SIGINT SIGTERM

fswatch_pid_file="$project_dir/fswatch.pid"

get_fswatch_pid() {
  while ! test -f "$fswatch_pid_file"; do
    sleep 1
    echo "Still waiting for $fswatch_pid_file..."
  done
  cat $fswatch_pid_file  
}

watch_webpack() {
  tail -F -n5000 $webpack_log | while IFS= read -r line;
  do
    if echo $line | grep 'webpack.*compiled successfully';
    then
      fswatch_pid=$(get_fswatch_pid)
      echo "Client bundle recompiled, restarting app (pid: $fswatch_pid)..."
      kill $fswatch_pid
    fi
  done
}

watch_backend() {
  eval "$cmd" || exit 1
  while true
  do
    start_app
    sleep $settle_down_secs
    set -x
    fswatch -r -1 --event Created --event Updated --event Removed -e '.*/[.]git/.*' -e 'web' -e $fswatch_pid_file'$' -e '.*/igo-repo/igo-repo$' . &
    fswatch_pid=$!
    echo $fswatch_pid > "$fswatch_pid_file"
    wait $fswatch_pid
    [[ "$stopping" == "true" ]] && exit
    rm -rf $fswatch_pid_file
    set +x
    pkill "$app_executable"
    eval "$cmd"
  done
}

cd $project_dir/web
echo "" > $webpack_log
watch_webpack &
npx webpack --watch 2>&1 | tee $webpack_log &
cd -

watch_backend

# You can watch the app instances' outputs with something like this:
# for i in $(seq 0 $((app_instance_count -1))); do tilix -a session-add-down -x "tail -f $app_log$i" & ; done
