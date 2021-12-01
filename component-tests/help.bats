#!/usr/bin/env bats

load "./support/common/load.bash"
source "./support/constants/help-outputs.sh"

runCommand="run pnapctl"

@test "No Input displays help" {
  $runCommand

  assert_success
  assert_output "$outputHelpLong"
}

@test "Run help command" {
  $runCommand help

  assert_success
  assert_output "$outputHelpLong"
}

@test "Run completion command" {
  $runCommand completion

  assert_success
  assert_output "$outputHelpCompletion"
}

@test "Get help command" {
  $runCommand get
  
  assert_success
  assert_output "$outputHelpGet"
}

@test "Create help command" {
  $runCommand create
  
  assert_success
  assert_output "$outputHelpCreate"
}

@test "Delete help command" {
  $runCommand delete
  
  assert_success
  assert_output "$outputHelpDelete"
}

@test "Patch help command" {
  $runCommand patch
  
  assert_success
  assert_output "$outputHelpPatch"
}

@test "Reset help command" {
  $runCommand reset
  
  assert_success
  assert_output "$outputHelpReset"
}

@test "Power Off help command" {
  $runCommand power-off
  
  assert_success
  assert_output "$outputHelpPowerOff"
}

@test "Power On help command" {
  $runCommand power-on
  
  assert_success
  assert_output "$outputHelpPowerOn"
}

@test "Reboot help command" {
  $runCommand reboot
  
  assert_success
  assert_output "$outputHelpReboot"
}

@test "Tag help command" {
  $runCommand tag
  
  assert_success
  assert_output "$outputHelpTag"
}

@test "Request Edit help command" {
  $runCommand request-edit
  
  assert_success
  assert_output "$outputHelpRequestEdit"
}

@test "Reserve help command" {
  $runCommand reserve
  
  assert_success
  assert_output "$outputHelpReserve"
}

@test "Shutdown help command" {
  $runCommand shutdown
  
  assert_success
  assert_output "$outputHelpShutdown"
}

@test "Update help command" {
  $runCommand update
  
  assert_success
  assert_output "$outputHelpUpdate"
}