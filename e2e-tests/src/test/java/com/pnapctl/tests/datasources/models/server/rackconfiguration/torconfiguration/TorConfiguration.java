package com.pnapctl.tests.datasources.models.server.rackconfiguration.torconfiguration;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record TorConfiguration(
        String torPairId,
        List<TorDetails> torDetails) {

}
