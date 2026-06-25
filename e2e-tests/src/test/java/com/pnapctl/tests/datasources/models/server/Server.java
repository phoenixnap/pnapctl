package com.pnapctl.tests.datasources.models.server;

import java.math.BigDecimal;
import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.pnapctl.tests.datasources.models.server.networkconfiguration.NetworkConfiguration;
import com.pnapctl.tests.datasources.models.server.osconfiguration.OsConfiguration;
import com.pnapctl.tests.datasources.models.server.storageconfiguration.StorageConfiguration;
import com.pnapctl.tests.datasources.models.tag.TagAssignment;

@JsonIgnoreProperties(ignoreUnknown = true)
public record Server(
        String id,
        String status,
        String hostname,
        String description,
        String os,
        String type,
        String location,
        String cpu,
        int cpuCount,
        int coresPerCpu,
        BigDecimal cpuFrequency,
        String ram,
        String storage,
        List<String> ipAddresses,
        List<String> privateIpAddresses,
        List<String> publicIpAddresses,
        String reservationId,
        String pricingModel,
        String password,
        String networkType,
        String clusterId,
        List<TagAssignment> tags,
        String provisionedOn,
        OsConfiguration osConfiguration,
        NetworkConfiguration networkConfiguration,
        StorageConfiguration storageConfiguration,
        GpuConfiguration gpuConfiguration,
        String supersededBy,
        String supersedes) {

}
