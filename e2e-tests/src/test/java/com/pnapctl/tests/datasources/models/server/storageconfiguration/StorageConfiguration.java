package com.pnapctl.tests.datasources.models.server.storageconfiguration;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record StorageConfiguration(
        RootPartition rootPartition) {

}
