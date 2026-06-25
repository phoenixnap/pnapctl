package com.pnapctl.tests.tests;

import static com.pnapctl.tests.datasources.constants.TestGroups.PNAPCTL_VERSION_GROUP;
import static com.pnapctl.tests.utils.CommandLineUtils.PNAPCTL_VERSION;
import static org.assertj.core.api.AssertionsForClassTypes.assertThat;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.testng.annotations.Test;

import com.pnapctl.tests.datasources.command.CommandResult;
import com.pnapctl.tests.datasources.executor.Pnapctl;

public class PnapctlVersionIT {

    private static final Logger LOGGER = LoggerFactory.getLogger(PnapctlVersionIT.class);

    /**
     * <b>Retrieve `pnapctl` version.</b>
     *
     * <p>
     * Test flow:
     * <ul>
     *     <li>Use the CLI command to get the pnapctl latest version.</li>
     *     <li>Assert that the command response was successful.</li>
     *     <li>Assert the command output version matches the latest one released.</li>
     * </ul>
     * </p>
     */
    @Test(groups = PNAPCTL_VERSION_GROUP)
    void pnapctl_version() {
        final CommandResult commandResult = Pnapctl.version().execute();

        LOGGER.info("Command details: \n{}", commandResult.stdout());
        assertThat(commandResult.stdout()).as("`pnapctl` version is not matching the latest").contains(PNAPCTL_VERSION);
    }

}
