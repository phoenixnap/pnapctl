#!/usr/bin/env bats

load "./support/common/load.bash"

@test "No Input displays help" {
  run pnapctl
  assert_failure

  assert_line --index 0 'pnapctl creates new and manages existing bare metal servers provided by the phoenixNAP Bare Metal Cloud service.'
  # We can add more assertions here
}

@test "Help" {
  run pnapctl help
  assert_success
}

@test "Get command Help" {
  run pnapctl get
  
  assert_line --index 0 'Display one or many resources.'
}

@test "Create Command Help" {
  run pnapctl create
  
  assert_line --index 0 'Create a resource.'
}

