package com.pnapctl.tests.datasources.executor;

import static com.pnapctl.tests.datasources.enums.MethodEnum.CREATE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.DELETE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.GET;
import static com.pnapctl.tests.datasources.enums.MethodEnum.UPDATE;

import com.pnapctl.tests.datasources.enums.ResourceEnum;
import com.pnapctl.tests.datasources.executor.ipblock.PnapctlIpBlocks;
import com.pnapctl.tests.datasources.executor.network.PnapctlPrivateNetworks;
import com.pnapctl.tests.datasources.executor.server.PnapctlServers;
import com.pnapctl.tests.datasources.executor.tag.PnapctlTags;
import com.pnapctl.tests.utils.CommandLineUtils;

/**
 * <a href="https://github.com/phoenixnap/pnapctl/blob/master/docs/">PNAPCTL Documentation</a>
 */
public class Pnapctl {

    private static final PnapctlServers SERVERS = new PnapctlServers();
    private static final PnapctlIpBlocks IP_BLOCKS = new PnapctlIpBlocks();
    private static final PnapctlTags TAGS = new PnapctlTags();
    private static final PnapctlPrivateNetworks PRIVATE_NETWORKS = new PnapctlPrivateNetworks();

    private Pnapctl() {
    }

    public static PnapctlCommand version() {
        return new PnapctlCommand(CommandLineUtils.baseCommand("version"));
    }

    // Basic HTTP methods
    public static PnapctlCommand get(ResourceEnum resource) {
        return new PnapctlCommand(CommandLineUtils.baseCommand(GET.value(), resource.value()));
    }

    public static PnapctlCommand getById(ResourceEnum resource, String id) {
        return new PnapctlCommand(CommandLineUtils.baseCommand(GET.value(), resource.value(), id));
    }

    public static PnapctlCommand create(ResourceEnum resource) {
        return new PnapctlCommand(CommandLineUtils.baseCommand(CREATE.value(), resource.value()));
    }

    public static PnapctlCommand update(ResourceEnum resource) {
        return new PnapctlCommand(CommandLineUtils.baseCommand(UPDATE.value(), resource.value()));
    }

    public static PnapctlCommand delete(ResourceEnum resource) {
        return new PnapctlCommand(CommandLineUtils.baseCommand(DELETE.value(), resource.value()));
    }

    // Predefined methods
    public static PnapctlServers servers() {
        return SERVERS;
    }

    public static PnapctlIpBlocks ipBlocks() {
        return IP_BLOCKS;
    }

    public static PnapctlTags tags() {
        return TAGS;
    }

    public static PnapctlPrivateNetworks privateNetworks() {
        return PRIVATE_NETWORKS;
    }
}
