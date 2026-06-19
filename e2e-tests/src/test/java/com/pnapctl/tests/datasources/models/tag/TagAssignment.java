package com.pnapctl.tests.datasources.models.tag;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

@JsonIgnoreProperties(ignoreUnknown = true)
public record TagAssignment(String id, String name, String value, boolean isBillingTag, String createdBy) {

}
