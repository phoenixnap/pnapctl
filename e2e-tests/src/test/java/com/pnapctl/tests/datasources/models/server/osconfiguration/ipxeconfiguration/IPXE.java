package com.pnapctl.tests.datasources.models.server.osconfiguration.ipxeconfiguration;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record IPXE(String url, NativeVlanConfiguration nativeVlanConfiguration) {

    @JsonIgnoreProperties(ignoreUnknown = true)
    public record NativeVlanConfiguration(Integer vlanId, String staticDhcpAddressV4, String status) {}
}
