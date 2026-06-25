package com.pnapctl.tests.datasources.models.server;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record GpuConfiguration (
    String longName,
    Integer count){

}
