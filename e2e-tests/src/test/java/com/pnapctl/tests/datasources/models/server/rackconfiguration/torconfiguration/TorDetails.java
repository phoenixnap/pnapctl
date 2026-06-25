package com.pnapctl.tests.datasources.models.server.rackconfiguration.torconfiguration;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record TorDetails(
        String address,
        int port) {
}
