package com.pnapctl.tests.datasources.enums;

import static com.pnapctl.tests.utils.ObjectMapperUtils.OBJECT_MAPPER;
import static com.pnapctl.tests.utils.ObjectMapperUtils.YAML_MAPPER;

import com.fasterxml.jackson.databind.ObjectMapper;

public enum OutputFormatEnum {

    // Default
    TABLE("table", null),
    JSON("json", OBJECT_MAPPER),
    YAML("yaml", YAML_MAPPER);

    private final String value;
    private final ObjectMapper mapper;

    OutputFormatEnum(String value, ObjectMapper mapper) {
        this.value = value;
        this.mapper = mapper;
    }

    public String value() {
        return value;
    }

    public ObjectMapper mapper() {
        return mapper;
    }
}
