package com.pnapctl.tests.datasources.models.ipblock;

import java.util.List;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.pnapctl.tests.datasources.models.tag.TagAssignment;

@JsonIgnoreProperties(ignoreUnknown = true)
public record IpBlock(
        String id,
        String location,
        String cidrBlockSize,
        String cidr,
        String ipVersion,
        String status,
        String parentIpBlockAllocationId,
        String assignedResourceId,
        String assignedResourceType,
        String description,
        List<TagAssignment> tags,
        Boolean isSystemManaged,
        boolean isBringYourOwn,
        String createdOn) {

}
