package com.pnapctl.tests.datasources.models.server.osconfiguration.windows;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record Windows(
        List<String> rdpAllowedIps,
        boolean bringYourOwnLicense) {

}
