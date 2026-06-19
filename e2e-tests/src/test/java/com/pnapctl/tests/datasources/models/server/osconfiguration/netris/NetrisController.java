package com.pnapctl.tests.datasources.models.server.osconfiguration.netris;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record NetrisController(

        // All fields are read only whereas `netrisWebConsoleUrl` and `netrisUserPassword` will only be returned in response to provisioning a server.
        String hostOs,
        String netrisWebConsoleUrl,
        String netrisUserPassword) {

}
