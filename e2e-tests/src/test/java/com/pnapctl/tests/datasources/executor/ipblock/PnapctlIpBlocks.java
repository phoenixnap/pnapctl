package com.pnapctl.tests.datasources.executor.ipblock;

import static com.pnapctl.tests.datasources.enums.MethodEnum.CREATE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.DELETE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.GET;
import static com.pnapctl.tests.datasources.enums.MethodEnum.UPDATE;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.IP_BLOCK;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.IP_BLOCKS;

import com.pnapctl.tests.datasources.executor.PnapctlCommand;
import com.pnapctl.tests.datasources.executor.PnapctlResource;

public class PnapctlIpBlocks extends PnapctlResource {

    public PnapctlCommand get() {
        return command(GET.value(), IP_BLOCKS.value());
    }

    public PnapctlCommand getById(String id) {
        return command(GET.value(), IP_BLOCK.value(), id);
    }

    public PnapctlCommand deleteById(String id) {
        return command(DELETE.value(), IP_BLOCK.value(), id);
    }

    public PnapctlCommand create() {
        return command(CREATE.value(), IP_BLOCK.value());
    }

    public PnapctlCommand updateById(String id) {
        return command(UPDATE.value(), IP_BLOCK.value(), id);
    }
}
