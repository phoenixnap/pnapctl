package com.pnapctl.tests.datasources.models.server.networkconfiguration.privatenetworkconfiguration;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record PrivateNetworkConfiguration(
        String gatewayAddress,
        String configurationType,
        List<ServerPrivateNetwork> privateNetworks) {

}
