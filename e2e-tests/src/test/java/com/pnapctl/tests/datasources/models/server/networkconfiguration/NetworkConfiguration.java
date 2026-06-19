package com.pnapctl.tests.datasources.models.server.networkconfiguration;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.pnapctl.tests.datasources.models.server.networkconfiguration.ipblocksconfiguration.IpBlocksConfiguration;
import com.pnapctl.tests.datasources.models.server.networkconfiguration.privatenetworkconfiguration.PrivateNetworkConfiguration;
import com.pnapctl.tests.datasources.models.server.networkconfiguration.publicnetworkconfiguration.PublicNetworkConfiguration;

@JsonIgnoreProperties(ignoreUnknown = true)
public record NetworkConfiguration(
        IpBlocksConfiguration ipBlocksConfiguration,
        String gatewayAddress,
        PrivateNetworkConfiguration privateNetworkConfiguration,
        PublicNetworkConfiguration publicNetworkConfiguration) {

}
