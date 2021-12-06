export PNAPCTL_HOME=$BATS_FILE_TMPDIR/.pnap

mkdir -p $PNAPCTL_HOME
cp $(dirname "${BASH_SOURCE[0]}")/../../../sample-config.yaml $PNAPCTL_HOME/config.yaml 