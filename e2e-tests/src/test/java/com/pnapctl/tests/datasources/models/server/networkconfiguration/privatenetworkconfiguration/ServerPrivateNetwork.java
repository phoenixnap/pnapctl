package com.pnapctl.tests.datasources.models.server.networkconfiguration.privatenetworkconfiguration;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record ServerPrivateNetwork(
        String id,
        List<String> ips,
        boolean dhcp,
        String statusDescription,
        int vlanId) {

}
