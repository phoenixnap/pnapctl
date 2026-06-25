package com.pnapctl.tests.datasources.executor;

import com.pnapctl.tests.utils.CommandLineUtils;

public abstract class PnapctlResource {

    protected PnapctlCommand command(String... args) {
        return new PnapctlCommand(CommandLineUtils.baseCommand(args));
    }
}
