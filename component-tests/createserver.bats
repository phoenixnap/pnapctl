#!/usr/bin/env bats

load '/usr/lib/node_modules/bats-support/load.bash'
load '/usr/lib/node_modules/bats-assert/load.bash'

@test "Create Server no input Fails" {
    run bin/pnapctl create server
    
    assert_failure

    assert_output 'Error: required flag(s) "filename" not set'
}


@test "Create Server non existent input file Fails" {
    run bin/pnapctl create server --filename test.json
    
    assert_failure

    assert_output "Error: The file 'test.json' does not exist."
}

@test "Create Server invalid JSON input file content Fails" {
    echo { , } >> test.json

    run bin/pnapctl create server --filename test.json
    
    assert_failure

    assert_output "Error: Command 'create server' has been performed, but something went wrong. Error code: 0303"
}

@test "Create Server unknown JSON field input file content Fails" {
    echo { "unknownField" : "anc" } >> test.json

    run bin/pnapctl create server --filename test.json
    
    assert_failure

    assert_output "Error: Command 'create server' has been performed, but something went wrong. Error code: 0303"
}

@test "Create Server Invalid Request Fails" {
    echo { } >> test.json

    run bin/pnapctl create server --filename test.json
    
    assert_failure

    assert_output "Error: 400 Bad Request"
}

teardown() {
    rm -f test.json
}