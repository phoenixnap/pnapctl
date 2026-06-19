package com.pnapctl.tests.datasources.models.server;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.pnapctl.tests.datasources.models.server.osconfiguration.ipxeconfiguration.IPXE;

@JsonIgnoreProperties(ignoreUnknown = true)
public record OperatingSystemConfiguration(
        String osName,
        Boolean provisionOsInRam,
        Boolean cloudInit,
        IPXE iPXE) {
}
