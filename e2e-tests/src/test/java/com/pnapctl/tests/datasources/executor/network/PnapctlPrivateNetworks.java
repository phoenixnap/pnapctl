package com.pnapctl.tests.datasources.executor.network;

import static com.pnapctl.tests.datasources.enums.MethodEnum.CREATE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.DELETE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.GET;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.PRIVATE_NETWORK;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.PRIVATE_NETWORKS;

import com.pnapctl.tests.datasources.executor.PnapctlCommand;
import com.pnapctl.tests.datasources.executor.PnapctlResource;

public class PnapctlPrivateNetworks extends PnapctlResource {

    public PnapctlCommand get() {
        return command(GET.value(), PRIVATE_NETWORKS.value());
    }

    public PnapctlCommand getById(String id) {
        return command(GET.value(), PRIVATE_NETWORK.value(), id);
    }

    public PnapctlCommand deleteById(String id) {
        return command(DELETE.value(), PRIVATE_NETWORK.value(), id);
    }

    public PnapctlCommand create() {
        return command(CREATE.value(), PRIVATE_NETWORK.value());
    }
}
