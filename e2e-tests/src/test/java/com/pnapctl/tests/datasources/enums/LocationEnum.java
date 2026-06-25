package com.pnapctl.tests.datasources.enums;

import java.util.Arrays;
import java.util.List;

public enum LocationEnum {

    PHX("PHX", "Phoenix, USA"),
    ASH("ASH", "Ashburn, USA"),
    SGP("SGP", "Singapore, SGP"),
    NLD("NLD", "Amsterdam, NLD"),
    CHI("CHI", "Chicago, USA"),
    SEA("SEA", "Seattle, USA");

    private final String location;
    private final String description;

    LocationEnum(final String location, final String description) {
        this.location = location;
        this.description = description;
    }

    public String getLocation() {
        return location;
    }

    public String getDescription() {
        return description;
    }

    public static List<String> getAllLocations() {
        return Arrays.stream(LocationEnum.values())
                .map(LocationEnum::getLocation)
                .toList();
    }
}
