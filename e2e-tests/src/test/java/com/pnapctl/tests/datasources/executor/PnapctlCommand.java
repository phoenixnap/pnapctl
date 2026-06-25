package com.pnapctl.tests.datasources.executor;

import static com.pnapctl.tests.datasources.enums.FlagEnum.FILENAME;
import static com.pnapctl.tests.datasources.enums.FlagEnum.OUTPUT;
import static com.pnapctl.tests.datasources.enums.OutputFormatEnum.JSON;
import static com.pnapctl.tests.utils.ObjectMapperUtils.convertToJsonString;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

import org.apache.commons.exec.CommandLine;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.JavaType;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.pnapctl.tests.datasources.command.CommandResult;
import com.pnapctl.tests.datasources.enums.FlagEnum;
import com.pnapctl.tests.datasources.enums.OutputFormatEnum;

public class PnapctlCommand {

    private final CommandLine commandLine;
    private OutputFormatEnum outputFormat = JSON; // Defaulting always to expect JSON as returned format

    public PnapctlCommand(CommandLine commandLine) {
        this.commandLine = commandLine;
    }

    private static <T> T readResultValue(ObjectMapper objectMapper, String stringObject, JavaType type) {
        try {
            return objectMapper.readValue(stringObject, type);
        } catch (JsonProcessingException e) {
            throw new IllegalStateException("Failed to parse CLI response", e);
        }
    }

    public PnapctlCommand withBody(Object body) {
        try {
            Path file = Files.createTempFile("pnapctl-", ".json");
            Files.writeString(file, convertToJsonString(body));
            commandLine.addArgument(FILENAME.shortValue());
            commandLine.addArgument(file.toString());
            return this;

        } catch (IOException e) {
            throw new IllegalStateException(e);
        }
    }

    public PnapctlCommand flag(FlagEnum... flagEnums) {
        for (FlagEnum flagEnum : flagEnums) {
            commandLine.addArgument(flagEnum.longValue());
        }
        return this;
    }

    public PnapctlCommand output(OutputFormatEnum format) {
        this.outputFormat = format;
        commandLine.addArgument(OUTPUT.shortValue());
        commandLine.addArgument(format.value());
        return this;
    }

    public CommandResult execute() {
        return CommandExecutor.execute(commandLine);
    }

    public CommandResult executeSuccessfully() {
        CommandResult result = execute();
        if (!result.isSuccessful()) {
            throw new AssertionError("""
                            Command failed
                            Exit code: [%s]
                            STDERR:
                            [%s]
                            """.formatted(result.exitCode(), result.stderr()));
        }
        return result;
    }

    public <T> T executeAndParse(Class<T> clazz) {
        final ObjectMapper mapper = getMapper();
        final CommandResult result = executeSuccessfully();
        final JavaType type = mapper.getTypeFactory().constructType(clazz);
        return readResultValue(mapper, result.stdout(), type);
    }

    public <T> T executeAndParse(TypeReference<T> typeReference) {
        final ObjectMapper mapper = getMapper();
        final CommandResult result = executeSuccessfully();
        final JavaType type = mapper.getTypeFactory().constructType(typeReference);
        return readResultValue(mapper, result.stdout(), type);
    }

    private ObjectMapper getMapper() {
        final ObjectMapper mapper = outputFormat.mapper();
        if (mapper == null) {
            throw new IllegalStateException("Output format '%s' cannot be parsed. Use JSON or YAML.".formatted(outputFormat));
        }
        return mapper;
    }
}
