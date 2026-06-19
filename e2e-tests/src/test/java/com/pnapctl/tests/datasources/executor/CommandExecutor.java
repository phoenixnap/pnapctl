package com.pnapctl.tests.datasources.executor;

import static com.pnapctl.tests.utils.CommandLineUtils.PNAPCTL_IMAGE;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.Arrays;

import org.apache.commons.exec.CommandLine;
import org.apache.commons.exec.DefaultExecutor;
import org.apache.commons.exec.ExecuteException;
import org.apache.commons.exec.PumpStreamHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.pnapctl.tests.datasources.command.CommandResult;

public final class CommandExecutor {

    private static final Logger LOGGER = LoggerFactory.getLogger(CommandExecutor.class);

    private CommandExecutor() {
    }

    public static CommandResult execute(CommandLine commandLine) {
        final ByteArrayOutputStream stdout = new ByteArrayOutputStream();
        final ByteArrayOutputStream stderr = new ByteArrayOutputStream();

        final DefaultExecutor executor = DefaultExecutor.builder().get();
        executor.setStreamHandler(new PumpStreamHandler(stdout, stderr));

        final long start = System.nanoTime();
        try {
            int exitCode = executor.execute(commandLine);
            long durationMs = (System.nanoTime() - start) / 1_000_000;

            LOGGER.info("Command completed in {} ms (exitCode={}): [{}]", durationMs, exitCode, commandDescription(commandLine));
            return new CommandResult(
                    exitCode,
                    stdout.toString().trim(),
                    stderr.toString().trim()
            );

        } catch (ExecuteException e) {
            long durationMs = (System.nanoTime() - start) / 1_000_000;

            LOGGER.warn("""
                    Command failed in {} ms (exitCode={}): [{}]
                    Output: [{}]
                    Error: [{}]
                    """, durationMs, e.getExitValue(), commandDescription(commandLine), stdout.toString().trim(), stderr.toString().trim());

            return new CommandResult(
                    e.getExitValue(),
                    stdout.toString().trim(),
                    stderr.toString().trim()
            );
        } catch (IOException e) {
            long durationMs = (System.nanoTime() - start) / 1_000_000;

            LOGGER.error("Command execution error after [{}] ms: [{}]", durationMs, commandDescription(commandLine), e);
            throw new IllegalStateException("Failed to execute command: %s".formatted(commandDescription(commandLine)), e);
        }
    }

    private static String commandDescription(CommandLine commandLine) {
        String[] args = commandLine.toStrings();

        for (int i = 0; i < args.length; i++) {
            if (PNAPCTL_IMAGE.equals(args[i])) {
                if (i == args.length - 1) {
                    return "";
                }
                return String.join(" ", Arrays.copyOfRange(args, i + 1, args.length));
            }
        }
        return "";
    }
}
