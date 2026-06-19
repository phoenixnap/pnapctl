package com.pnapctl.tests.datasources.payload;

import com.pnapctl.tests.datasources.enums.LocationEnum;
import com.pnapctl.tests.utils.ObjectMapperUtils;

public record IpamPayload(
        String cidrBlockSize,
        String location,
        String ipVersion,
        String description
) {

    public static IpamPayload generateCreateIpBlockPayload(String cidrBlockSize, String ipVersion) {
        return IpamPayload.builder()
                .cidrBlockSize(cidrBlockSize)
                .ipVersion(ipVersion)
                .location(LocationEnum.PHX.getLocation())
                .description("PNAPCTL Creation")
                .build();
    }

    public static Builder builder() {
        return new Builder();
    }

    public String convertToJsonString() {
        return ObjectMapperUtils.convertToJsonString(this);
    }

    public static final class Builder {

        private String cidrBlockSize;
        private String location;
        private String ipVersion;
        private String description;

        public Builder cidrBlockSize(String cidrBlockSize) {
            this.cidrBlockSize = cidrBlockSize;
            return this;
        }

        public Builder location(String location) {
            this.location = location;
            return this;
        }

        public Builder ipVersion(String ipVersion) {
            this.ipVersion = ipVersion;
            return this;
        }

        public Builder description(String description) {
            this.description = description;
            return this;
        }

        public IpamPayload build() {
            return new IpamPayload(cidrBlockSize, location, ipVersion, description);
        }
    }
}
