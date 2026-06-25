package com.pnapctl.tests.datasources.models.server.networkconfiguration.publicnetworkconfiguration;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record ServerPublicNetwork(
        String id,
        List<String> ips,
        String statusDescription,
        int vlanId) {

}
