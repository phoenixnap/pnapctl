package com.pnapctl.tests.datasources.models.server.osconfiguration.netris;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record NetrisSoftgate(
        String hostOs, // ReadOnly
        String controllerAddress, // WriteOnly
        String controllerVersion, // WriteOnly
        String controllerAuthKey // WriteOnly
) {

}
