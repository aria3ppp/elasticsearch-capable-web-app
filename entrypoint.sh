#!/bin/bash -e

echo "[`date`] Running entrypoint script..."

if [[ -z ${APP_DSN} ]]; then
  export APP_DSN=`sed -n 's/^DSN="\(.*\)"/\1/p' .env`
fi

echo "[`date`] Running DB migrations..."
migrate -database "${APP_DSN}" -path ./migrations up

echo "[`date`] Starting server..."
./server
