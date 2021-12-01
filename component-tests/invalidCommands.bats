#!/usr/bin/env bats

load "./support/common/load.bash"
source "./support/constants/create-error-output.sh"

@test "Run invalid create command input" {

    run pnapctl $invalidCommandCreate

    assert_failure
    assert_output "$expectedOutputWrongCommandError"
}

@test "Run nonexistent command input" {
    run pnapctl $invalidCommandPost 

    assert_failure
    assert_output "$expectedOutputInvalidPostError"
}