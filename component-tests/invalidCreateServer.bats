#!/usr/bin/env bats

load "./support/common/load.bash"
source "./support/constants/create-error-output.sh"

runCommand="run pnapctl create server"
runCommandWithJsonFile="$runCommand --filename $jsonFile"
runCommandWithYamlFile="$runCommand --filename $yamlFile"

@test "Run server provision command without specifying filename" {
    $runCommand
    
    assert_failure
    assert_output "$expectedOutputNotSpecifiedFilename"
}

@test "Run server provision command with nonexistent json file" {
    $runCommandWithJsonFile
    
    assert_failure
    assert_output "$expectedOutputNonexistentJsonFile"
}

@test "Run server provision command with nonexistent yaml file" {
    $runCommandWithYamlFile
    
    assert_failure
    assert_output "$expectedOutputNonexistentYamlFile"
}

@test "Run server provision command with invalid json payload" {
    echo { "unknownField" : "anc" } >> $jsonFile
    $runCommandWithJsonFile
    
    assert_failure
    assert_output "$expectedOutputWrongFile"
}

@test "Run server provision command with invalid filename command" {
    echo { "unknownField" : "anc" } >> $jsonFile
    $runCommand -- filename $jsonFile
    
    assert_failure
    assert_output "$expectedOutputInvalidFilenameCommand"
}

@test "Run server provision command with invalid yaml payload" {
    echo  "test:test"  >> $yamlFile
    $runCommandWithYamlFile
    
    assert_failure
    assert_output "$expectedOutputWrongFile"
}

@test "Run server provision command with invalid credentials" {
    echo { } >> $jsonFile
    echo "clientId: <CLIENT_ID>
clientSecret: <CLIENT_SECRET>" >> $configFile
    
    $runCommandWithJsonFile --config=$configFile

    assert_failure
    assert_output "$expectedOutputCredentialError"
}

@test "Run server provision command with without specifying the resource" {    
    $runCommand

    assert_failure
    assert_output "$expectedOutputFileNotSetError"
}

teardown() {
    rm -f $jsonFile
    rm -f $yamlFile
    rm -f $configFile
}