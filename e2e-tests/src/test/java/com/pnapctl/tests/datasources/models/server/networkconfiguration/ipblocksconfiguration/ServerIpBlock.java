package com.pnapctl.tests.datasources.models.server.networkconfiguration.ipblocksconfiguration;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record ServerIpBlock(
        String id,
        int vlanId) {

}
