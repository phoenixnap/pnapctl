package com.pnapctl.tests.datasources.models.server.networkconfiguration.ipblocksconfiguration;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record IpBlocksConfiguration(
        String configurationType,
        List<ServerIpBlock> ipBlocks) {

}
