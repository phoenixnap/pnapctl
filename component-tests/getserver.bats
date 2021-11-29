#!/usr/bin/env bats

load './support/common/load.bash'

@test "Get Servers success" {
    run pnapctl get servers

    assert_success
}