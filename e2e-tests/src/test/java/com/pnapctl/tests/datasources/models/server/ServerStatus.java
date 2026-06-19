package com.pnapctl.tests.datasources.models.server;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record ServerStatus(
        String name,
        String details,
        String setOn,
        String setBy) {
}
