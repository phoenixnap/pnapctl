#!/usr/bin/env bats

load "./support/common/load.bash"

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

    assert_output "Error: Command 'create server' has been performed, but something went wrong. Error code: 0303"

}

@test "Run server provision command with empty json payload" {
    echo { } >> test.json
    expectedOutputError="Error: Post \"https://api.phoenixnap.com/bmc/v0/servers\": oauth2: cannot fetch token: 400 Bad Request
Response: {\"error\":\"invalid_client\",\"error_description\":\"Invalid client credentials\"}"
    
    $runCommandWithFilename

    assert_failure

    assert_output "$expectedOutputError"
    
}

@test "Run server provision command with invalid command input" {
    invalidCommand="createe"
    expectedOutputError="Error: unknown command \"$invalidCommand\" for \"pnapctl\"

Did you mean this?
	create

Run 'pnapctl --help' for usage."
    
    run pnapctl $invalidCommand server 

    assert_failure

    assert_output "$expectedOutputError"
    
}

@test "Run server provision command with nonexitant command" {
    invalidCommand="post"
    expectedOutputError="Error: unknown command \"$invalidCommand\" for \"pnapctl\"
Run 'pnapctl --help' for usage."
    
    run pnapctl $invalidCommand server 

    assert_failure

    assert_output "$expectedOutputError"
    
}

@test "Run server provision command with without specifying the resource" {
    expectedOutputError="Error: required flag(s) \"filename\" not set"
    
    $runCommand

    assert_failure

    assert_output "$expectedOutputError"
    
}

teardown() {
    rm -f test.json
}