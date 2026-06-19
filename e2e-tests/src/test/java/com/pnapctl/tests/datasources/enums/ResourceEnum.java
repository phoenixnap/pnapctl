package com.pnapctl.tests.datasources.enums;

public enum ResourceEnum {
    SERVERS("servers"),
    SERVER("server"),
    TAGS("tags"),
    TAG("tag"),
    PRIVATE_NETWORKS("private-networks"),
    PRIVATE_NETWORK("private-network"),
    IP_BLOCKS("ip-blocks"),
    IP_BLOCK("ip-block");

    private final String value;

    ResourceEnum(String value) {
        this.value = value;
    }

    public String value() {
        return value;
    }
}
