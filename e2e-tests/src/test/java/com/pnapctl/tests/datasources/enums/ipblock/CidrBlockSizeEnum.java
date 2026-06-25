package com.pnapctl.tests.datasources.enums.ipblock;

import java.util.Arrays;

public enum CidrBlockSizeEnum {

    // Used only for Private Networks
    CIDR_BLOCK_SIZE_1("/1", 1),
    CIDR_BLOCK_SIZE_2("/2", 2),
    CIDR_BLOCK_SIZE_3("/3", 3),
    CIDR_BLOCK_SIZE_4("/4", 4),
    CIDR_BLOCK_SIZE_5("/5", 5),
    CIDR_BLOCK_SIZE_6("/6", 6),
    CIDR_BLOCK_SIZE_7("/7", 7),
    CIDR_BLOCK_SIZE_8("/8", 8),
    CIDR_BLOCK_SIZE_9("/9", 9),
    CIDR_BLOCK_SIZE_10("/10", 10),
    CIDR_BLOCK_SIZE_11("/11", 11),
    CIDR_BLOCK_SIZE_12("/12", 12),
    CIDR_BLOCK_SIZE_13("/13", 13),
    CIDR_BLOCK_SIZE_14("/14", 14),
    CIDR_BLOCK_SIZE_15("/15", 15),

    // Used for BYO - from /23 up to /16
    CIDR_BLOCK_SIZE_16("/16", 16),
    CIDR_BLOCK_SIZE_17("/17", 17),
    CIDR_BLOCK_SIZE_18("/18", 18),
    CIDR_BLOCK_SIZE_19("/19", 19),
    CIDR_BLOCK_SIZE_20("/20", 20),
    CIDR_BLOCK_SIZE_21("/21", 21),
    CIDR_BLOCK_SIZE_22("/22", 22),
    CIDR_BLOCK_SIZE_23("/23", 23),

    CIDR_BLOCK_SIZE_24("/24", 24),
    CIDR_BLOCK_SIZE_25("/25", 25),
    CIDR_BLOCK_SIZE_26("/26", 26),
    CIDR_BLOCK_SIZE_27("/27", 27),
    CIDR_BLOCK_SIZE_28("/28", 28),
    CIDR_BLOCK_SIZE_29("/29", 29),
    CIDR_BLOCK_SIZE_30("/30", 30),
    CIDR_BLOCK_SIZE_31("/31", 31),
    CIDR_BLOCK_SIZE_32("/32", 32),
    CIDR_BLOCK_SIZE_33("/33", 33),
    CIDR_BLOCK_SIZE_34("/34", 34),
    CIDR_BLOCK_SIZE_35("/35", 35),
    CIDR_BLOCK_SIZE_36("/36", 36),
    CIDR_BLOCK_SIZE_37("/37", 37),
    CIDR_BLOCK_SIZE_38("/38", 38),
    CIDR_BLOCK_SIZE_39("/39", 39),
    CIDR_BLOCK_SIZE_40("/40", 40),
    CIDR_BLOCK_SIZE_41("/41", 41),
    CIDR_BLOCK_SIZE_42("/42", 42),
    CIDR_BLOCK_SIZE_43("/43", 43),
    CIDR_BLOCK_SIZE_44("/44", 44),
    CIDR_BLOCK_SIZE_45("/45", 45),
    CIDR_BLOCK_SIZE_46("/46", 46),
    CIDR_BLOCK_SIZE_47("/47", 47),
    CIDR_BLOCK_SIZE_48("/48", 48),
    CIDR_BLOCK_SIZE_49("/49", 49),
    CIDR_BLOCK_SIZE_50("/50", 50),
    CIDR_BLOCK_SIZE_51("/51", 51),
    CIDR_BLOCK_SIZE_52("/52", 52),
    CIDR_BLOCK_SIZE_53("/53", 53),
    CIDR_BLOCK_SIZE_54("/54", 54),
    CIDR_BLOCK_SIZE_55("/55", 55),
    CIDR_BLOCK_SIZE_56("/56", 56),
    CIDR_BLOCK_SIZE_57("/57", 57),
    CIDR_BLOCK_SIZE_58("/58", 58),
    CIDR_BLOCK_SIZE_59("/59", 59),
    CIDR_BLOCK_SIZE_60("/60", 60),
    CIDR_BLOCK_SIZE_61("/61", 61),
    CIDR_BLOCK_SIZE_62("/62", 62),
    CIDR_BLOCK_SIZE_63("/63", 63),
    CIDR_BLOCK_SIZE_64("/64", 64);

    private final String cidrSize;
    private final int intCidrSize;

    CidrBlockSizeEnum(String cidrSize, int intCidrSize) {
        this.cidrSize = cidrSize;
        this.intCidrSize = intCidrSize;
    }

    public static CidrBlockSizeEnum getCidrBlockSizeFromInteger(Integer cidrBlockSizeInteger) {
        return Arrays.stream(CidrBlockSizeEnum.values())
                .filter(cidrBlockSizeEnum -> cidrBlockSizeEnum.getIntValue() == cidrBlockSizeInteger)
                .findFirst()
                .orElseThrow(() -> new RuntimeException(
                        String.format("The provided integer [%s] has no corresponding CIDR Block Size", cidrBlockSizeInteger))
                );
    }

    public String getValue() {
        return cidrSize;
    }

    public int getIntValue() {
        return intCidrSize;
    }
}
