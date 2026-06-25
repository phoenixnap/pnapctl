package com.pnapctl.tests.datasources.enums;

public enum FlagEnum {
    OUTPUT("-o", "--output"),       // Define the output format. Possible values: table, json, yaml (default "table").
    FILENAME("-f", "--filename"),   // File containing required information for creation of the resource.
    FULL("", "--full"),             // Shows all response details.
    VERBOSE("-v", "--verbose");     // Change log level from Warn (default) to Debug.

    private final String shortValue;
    private final String longValue;

    FlagEnum(String shortValue, String longValue) {
        this.shortValue = shortValue;
        this.longValue = longValue;
    }

    public String shortValue() {
        return shortValue;
    }

    public String longValue() {
        return longValue;
    }
}
