package com.pnapctl.tests.tests;

import static org.testng.Assert.assertEquals;
import static org.testng.Assert.assertFalse;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.testng.annotations.Test;

public class PnapctlE2ETest {

    private static final Logger LOGGER = LoggerFactory.getLogger(PnapctlE2ETest.class);

    @Test
    void pnapctl_help() throws Exception {

        String image = System.getenv("PNAPCTL_IMAGE");

        LOGGER.info("Logging image {}", image);
        Process process = new ProcessBuilder(
                "docker",
                "run",
                "--rm",
                image,
                "--help")
                .redirectErrorStream(true)
                .start();

        String output =
                new String(process.getInputStream().readAllBytes());

        LOGGER.info("Logging output {}", output);

        int exitCode = process.waitFor();

        assertEquals(0, exitCode);
        assertFalse(output.isBlank());
    }
}
