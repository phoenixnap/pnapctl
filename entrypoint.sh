#!/bin/bash
set -e

mkdir -p /root/.pnap

cat <<EOF > /root/.pnap/config.yaml
clientId: ${PNAP_CLIENT_ID}
clientSecret: ${PNAP_CLIENT_SECRET}
EOF

export PNAPCTL_HOME=/root/.pnap

# run real CLI command
exec pnapctl "$@"
