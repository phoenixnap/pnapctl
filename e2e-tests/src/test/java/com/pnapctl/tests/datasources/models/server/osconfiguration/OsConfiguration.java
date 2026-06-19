package com.pnapctl.tests.datasources.models.server.osconfiguration;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.pnapctl.tests.datasources.models.server.osconfiguration.ipxeconfiguration.IPXE;
import com.pnapctl.tests.datasources.models.server.osconfiguration.netris.NetrisController;
import com.pnapctl.tests.datasources.models.server.osconfiguration.netris.NetrisSoftgate;
import com.pnapctl.tests.datasources.models.server.osconfiguration.windows.Windows;

@JsonIgnoreProperties(ignoreUnknown = true)
public record OsConfiguration(
        NetrisController netrisController,
        NetrisSoftgate netrisSoftgate,
        String rootPassword,
        String managementUiUrl,
        List<String> managementAccessAllowedIps,
        Windows windows,
        IPXE iPXE
) {

}
