package com.pnapctl.tests.datasources.models.server.rackconfiguration;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.pnapctl.tests.datasources.models.server.rackconfiguration.torconfiguration.TorConfiguration;

@JsonIgnoreProperties(ignoreUnknown = true)
public record RackConfiguration(
        TorConfiguration torConfiguration,
        IpmiConfiguration ipmiConfiguration,
        List<String> macAddresses) {

    public record IpmiConfiguration(String address) {

    }
}
