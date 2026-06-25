package com.pnapctl.tests.datasources.command;

public record CommandResult(
        int exitCode,
        String stdout,
        String stderr
) {
    public boolean isSuccessful() {
        return exitCode == 0;
    }
}
