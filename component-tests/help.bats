#!/usr/bin/env bats

load "./support/common/load.bash"

runCommand="run pnapctl"

@test "No Input displays help" {
  $runCommand
  assert_success
  assert_line --index 0 "$outputHelp"
  # We can add more assertions here
}

@test "Help" {
  
  $runCommand help
  assert_success
  assert_output "$outputHelpLong"
}

@test "Get command Help" {
  $runCommand get
  
  assert_success
  assert_output "$outputHelpGet"
}

@test "Create Command Help" {

  $runCommand create
  
  assert_success
  assert_output "$outputHelpCreate"
}

