package com.pnapctl.tests.datasources.enums;

public enum MethodEnum {

    GET("get"),
    CREATE("create"),
    UPDATE("update"),
    PATCH("patch"),
    DELETE("delete");

    private final String value;

    MethodEnum(String value) {
        this.value = value;
    }

    public String value() {
        return value;
    }
}
