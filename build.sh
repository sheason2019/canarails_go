git_hash=$(git rev-parse HEAD)
build_time=$(date +"%Y.%m.%d_%H.%M.%S")

go build \
 -o bin/canarails \
 -ldflags=" \
 -X 'canarails.dev/services/aboutsvc.GitHash=$git_hash' \
 -X 'canarails.dev/services/aboutsvc.BuildTime=$build_time'" \
 .
