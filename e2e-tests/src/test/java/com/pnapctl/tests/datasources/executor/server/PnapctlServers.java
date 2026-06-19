package com.pnapctl.tests.datasources.executor.server;

import static com.pnapctl.tests.datasources.enums.MethodEnum.CREATE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.DELETE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.GET;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.SERVER;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.SERVERS;

import com.pnapctl.tests.datasources.executor.PnapctlCommand;
import com.pnapctl.tests.datasources.executor.PnapctlResource;

public class PnapctlServers extends PnapctlResource {

    public PnapctlCommand get() {
        return command(GET.value(), SERVERS.value());
    }

    public PnapctlCommand getById(String id) {
        return command(GET.value(), SERVER.value(), id);
    }

    public PnapctlCommand deleteById(String id) {
        return command(DELETE.value(), SERVER.value(), id);
    }

    public PnapctlCommand create() {
        return command(CREATE.value(), SERVER.value());
    }
}
