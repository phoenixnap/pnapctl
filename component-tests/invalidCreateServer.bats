#!/usr/bin/env bats

load "./support/common/load.bash"
source "./support/constants/create-error-output.sh"

runCommand="run pnapctl create server"
runCommandWithFilename="$runCommand --filename test.json"

@test "Run server provision command without specifying json file" {
    $runCommand
    
    assert_failure
    assert_output 'Error: required flag(s) "filename" not set'
}

@test "Run server provision command with nonexistent json file" {
    $runCommandWithFilename
    
    assert_failure
    assert_output "Error: The file 'test.json' does not exist."
}

@test "Run server provision command with invalid json payload" {
    echo { "unknownField" : "anc" } >> test.json
    $runCommandWithFilename
    
    assert_failure
    assert_output "$expectedOutputWrongFile"
}

@test "Run server provision command with invalid json payload command" {
    echo { "unknownField" : "anc" } >> test.json
    $runCommand -- filename test.json
    
    assert_failure
    assert_output "Error: accepts 0 arg(s), received 2"
}

@test "Run server provision command with invalid yaml payload" {
    echo { a } >> test.yaml
    $runCommand --filename test.yaml
    
    assert_failure
    assert_output "$expectedOutputWrongFile"
}

@test "Run server provision command with invalid credentials" {
    echo { } >> test.json
    
    $runCommandWithFilename

    assert_failure
    assert_output "$expectedOutputCredentialError"
}

@test "Run server provision command with invalid command input" {

    run pnapctl $invalidCommandCreate server 

    assert_failure
    assert_output "$expectedOutputWrongCommandError"
}

@test "Run server provision command with nonexitant command" {
    run pnapctl $invalidCommandPost server 

    assert_failure
    assert_output "$expectedOutputInvalidPostError"
}

@test "Run server provision command with without specifying the resource" {    
    $runCommand

    assert_failure
    assert_output "$expectedOutputFileNotSetError"
}

teardown() {
    rm -f test.json
    rm -f test.yaml
}