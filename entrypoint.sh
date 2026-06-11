#!/bin/bash
set -e

mkdir -p /root/.pnap

# Create config from env vars EVERY TIME container starts
if [ "$PNAP_ENV" = "dev" ]; then
  # DEV config (with hostnames)
  cat <<EOF > /root/.pnap/config.yaml
clientId: ${PNAP_CLIENT_ID}
clientSecret: ${PNAP_CLIENT_SECRET}

# URLs
bmcApiHostname: https://api-dev.phoenixnap.com/bmc/v1
rancherHostname: https://api-dev.phoenixnap.com/solutions/rancher/v1beta
tagsHostname: https://api-dev.phoenixnap.com/tag-manager/v1
auditHostname: https://api-dev.phoenixnap.com/audit/v1
billingHostname: https://api-dev.phoenixnap.com/billing/v0
networksHostname: https://api-dev.phoenixnap.com/networks/v1
ipHostname: https://api-dev.phoenixnap.com/ips/v1
networkStorageHostname: https://api-dev.phoenixnap.com/network-storage/v1
locationHostname: https://api-dev.phoenixnap.com/location-api/v1

# TokenURL represents the URL of the OpenID Connect provider from where we can retrieve a token
tokenURL: https://auth-dev.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token
EOF
else
  # PROD / default config (no hostnames)
  cat <<EOF > /root/.pnap/config.yaml
clientId: ${PNAP_CLIENT_ID}
clientSecret: ${PNAP_CLIENT_SECRET}
EOF
fi

export PNAPCTL_HOME=/root/.pnap

# run real CLI command
exec pnapctl "$@"
