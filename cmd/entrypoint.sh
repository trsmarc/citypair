#!/bin/sh
APP_ENV=${APP_ENV:-local}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

CONFIG_FILE=/app/config/${APP_ENV}.yaml

echo "[`date`] Starting server..."
/app/bin/server --config ${CONFIG_FILE}

