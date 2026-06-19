package com.pnapctl.tests.utils;

import org.apache.commons.exec.CommandLine;

public class CommandLineUtils {

    public static final String PNAPCTL_IMAGE = requireEnvironmentVars("PNAPCTL_IMAGE");
    public static final String PNAPCTL_VERSION = requireEnvironmentVars("PNAPCTL_VERSION");
    private static final String PNAP_CLIENT_ID = requireEnvironmentVars("PNAP_CLIENT_ID");
    private static final String PNAP_CLIENT_SECRET = requireEnvironmentVars("PNAP_CLIENT_SECRET");

    // Private constructor to avoid instantiation of utility class
    private CommandLineUtils() {
        throw new AssertionError("Utility class cannot be instantiated.");
    }

    public static CommandLine baseCommand(String... commandArgs) {
        CommandLine cmd = new CommandLine("docker");
        cmd.addArguments(new String[] {
                "run", "--rm",
                "-v", "/tmp:/tmp", // Mount the temp directory into container so it visible when reading temp files as payloads
                "-e", "PNAP_CLIENT_ID=%s".formatted(PNAP_CLIENT_ID),
                "-e", "PNAP_CLIENT_SECRET=%s".formatted(PNAP_CLIENT_SECRET),
                PNAPCTL_IMAGE
        });

        cmd.addArguments(commandArgs);
        return cmd;
    }

    private static String requireEnvironmentVars(String name) {
        final String value = System.getenv(name);

        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable '%s' is not set".formatted(name));
        }

        return value;
    }

}
