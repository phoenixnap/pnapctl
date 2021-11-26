#!/usr/bin/env bats

load '/usr/lib/node_modules/bats-support/load.bash'
load '/usr/lib/node_modules/bats-assert/load.bash'

@test "Get Servers success" {
    run bin/pnapctl get servers

    assert_success
}