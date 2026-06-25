package com.pnapctl.tests.utils;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.fasterxml.jackson.datatype.jdk8.Jdk8Module;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;

public class ObjectMapperUtils {

    public static final ObjectMapper OBJECT_MAPPER = configure(new ObjectMapper());
    public static final ObjectMapper YAML_MAPPER = configure(new ObjectMapper(new YAMLFactory()));

    // Private constructor to avoid instantiation of utility class
    private ObjectMapperUtils() {
        throw new AssertionError("Utility class cannot be instantiated.");
    }

    private static ObjectMapper configure(ObjectMapper mapper) {
        return mapper
                .configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)
                .registerModule(new JavaTimeModule())
                .registerModule(new Jdk8Module())
                .disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS);
    }

    /**
     * Converts an object to a JSON string using the Jackson library.
     *
     * @param object The object to be converted.
     * @return The JSON string representing the object.
     */
    public static String convertToJsonString(Object object) {
        try {
            return OBJECT_MAPPER.writeValueAsString(object);
        } catch (JsonProcessingException e) {
            throw new RuntimeException("Error converting object to JSON string", e);
        }
    }
}
