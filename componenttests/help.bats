#!/usr/bin/env bats

load '/usr/lib/node_modules/bats-support/load.bash'
load '/usr/lib/node_modules/bats-assert/load.bash'

@test "No Input displays help" {
  run bin/pnapctl
  assert_success

  assert_line --index 0 'pnapctl creates new and manages existing bare metal servers provided by the phoenixNAP Bare Metal Cloud service.'
  # We can add more assertions here
}

@test "Help" {
  run bin/pnapctl help
  assert_success
}

@test "Get command Help" {
  run bin/pnapctl get
  
  assert_line --index 0 'Display one or many resources.'
}

@test "Create Command Help" {
  run bin/pnapctl create
  
  assert_line --index 0 'Create a resource.'
}

