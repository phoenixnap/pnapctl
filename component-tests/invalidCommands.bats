#!/usr/bin/env bats

load "./support/common/load.bash"
source "./support/constants/invalid-commands-output.sh"

runCommand="run pnapctl"
@test "Run invalid create command input" {

    $runCommand $invalidCommandCreate

    assert_failure
    assert_output "$expectedOutputWrongCommandError"
}

@test "Run nonexistent command input" {
    $runCommand $invalidCommandPost 

    assert_failure
    assert_output "$expectedOutputInvalidPostError"
}
