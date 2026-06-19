package com.pnapctl.tests.tests;

import static com.pnapctl.tests.datasources.constants.TestGroups.CRUD_IP_BLOCK_GROUP;
import static com.pnapctl.tests.datasources.enums.OutputFormatEnum.JSON;
import static com.pnapctl.tests.datasources.enums.ipblock.CidrBlockSizeEnum.CIDR_BLOCK_SIZE_31;
import static com.pnapctl.tests.datasources.enums.ipblock.IpVersionEnum.V4;
import static com.pnapctl.tests.datasources.payload.IpamPayload.generateCreateIpBlockPayload;

import org.assertj.core.api.SoftAssertions;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.testng.annotations.AfterClass;
import org.testng.annotations.BeforeClass;
import org.testng.annotations.Test;

import com.pnapctl.tests.datasources.command.CommandResult;
import com.pnapctl.tests.datasources.executor.Pnapctl;
import com.pnapctl.tests.datasources.models.ipblock.IpBlock;
import com.pnapctl.tests.datasources.payload.IpamPayload;

public class CrudIpBlocksIT {

    private static final Logger LOGGER = LoggerFactory.getLogger(CrudIpBlocksIT.class);
    private static IpBlock ipBlock;
    private static IpamPayload payload;

    /**
     * <p>Setup method for this test class.</p>
     * <ul>
     *     <li>Initialize IP Block payload needed for the creation.</li>
     * </ul>
     */
    @BeforeClass(groups = CRUD_IP_BLOCK_GROUP)
    void setup() {
        payload = generateCreateIpBlockPayload(CIDR_BLOCK_SIZE_31.getValue(), V4.toString());
    }

    /**
     * <b>Create IP Block IPv4</b>
     *
     * <p>
     * Test flow:
     * <ul>
     *     <li>Use the CLI command to create an IP Block resource using the predefined payload for it.</li>
     *     <li>Assert that the command response was successful.</li>
     *     <li>Assert the command output matches the generated IP Block payload.</li>
     * </ul>
     * </p>
     */
    @Test(testName = "Create IP Block Test", groups = CRUD_IP_BLOCK_GROUP)
    void create_ip_block() {
        LOGGER.info("Starting Test: [Create IP Block Test]");
        ipBlock = Pnapctl.ipBlocks().create().withBody(payload).output(JSON).executeAndParse(IpBlock.class);

        LOGGER.info("IP Block with ID [{}] has been created.", ipBlock.id());
        assertIpBlock(ipBlock, payload);
    }

    /**
     * <b>Retrieve Created IP Block IPv4</b>
     *
     * <p>
     * Test flow:
     * <ul>
     *     <li>Use the CLI command to retrieve an IP Block resource created in the previous test.</li>
     *     <li>Assert that the command response was successful.</li>
     *     <li>Assert the command output matches the generated IP Block payload.</li>
     * </ul>
     * </p>
     */
    @Test(testName = "Retrieve IP Block Test", dependsOnMethods = "create_ip_block", groups = CRUD_IP_BLOCK_GROUP)
    void retrieve_ip_block_by_id() {
        LOGGER.info("Starting Test: [Retrieve IP Block Test]");
        IpBlock retrievedIpBlock = Pnapctl.ipBlocks().getById(ipBlock.id()).output(JSON).executeAndParse(IpBlock.class);
        assertIpBlock(retrievedIpBlock, payload);
    }

    /**
     * <b>Delete Created IP Block IPv4</b>
     *
     * <p>
     * Test flow:
     * <ul>
     *     <li>Use the CLI command to delete an IP Block resource created in the previous test.</li>
     *     <li>Use the CLI command to retrieve an IP Block resource which was just deleted.</li>
     *     <li>Assert that the command response was unsuccessful.</li>
     *     <li>Assert the command output error message matches expectations.</li>
     * </ul>
     * </p>
     */
    @Test(testName = "Delete IP Block Test", dependsOnMethods = "retrieve_ip_block_by_id", groups = CRUD_IP_BLOCK_GROUP)
    void delete_ip_block_by_id() {
        LOGGER.info("Starting Test: [Delete IP Block Test]");
        Pnapctl.ipBlocks().deleteById(ipBlock.id()).execute();

        LOGGER.info("IP Block with ID [{}] has been deleted.", ipBlock.id());
        CommandResult commandResult = Pnapctl.ipBlocks().getById(ipBlock.id()).execute();
        ipBlock = null;
        assertIpBlockIsNotPresent(commandResult);
    }

    /**
     * <b>Cleanup method for this test class.</b>
     * <ul>
     *     <li>If the IP Block was not deleted, use CLI command to request IP Block deletion.</li>
     * </ul>
     */
    @AfterClass(groups = CRUD_IP_BLOCK_GROUP)
    void cleanup() {
        if (ipBlock != null) {

            LOGGER.info("IP Block with ID [{}] has not been deleted. Deleting now...", ipBlock.id());
            Pnapctl.ipBlocks().deleteById(ipBlock.id()).execute();
        }
    }

    private void assertIpBlock(IpBlock ipBlock, IpamPayload ipamPayload) {
        final SoftAssertions softly = new SoftAssertions();
        softly.assertThat(ipBlock.id()).as("IP Block ID is null.").isNotNull();
        softly.assertThat(ipBlock.cidrBlockSize()).as("IP Block CIDR Size is matching.").isEqualTo(ipamPayload.cidrBlockSize());
        softly.assertThat(ipBlock.ipVersion()).as("IP Block IP Version is matching.").isEqualTo(ipamPayload.ipVersion());
        softly.assertThat(ipBlock.location()).as("IP Block Location is matching.").isEqualTo(ipamPayload.location());
        softly.assertThat(ipBlock.description()).as("IP Block Description is matching.").isEqualTo(ipamPayload.description());
        softly.assertAll();
    }

    private void assertIpBlockIsNotPresent(CommandResult commandResult) {
        final SoftAssertions softly = new SoftAssertions();
        softly.assertThat(commandResult.isSuccessful()).as("IP Block Get by ID state is not matching.").isFalse();
        softly.assertThat(commandResult.stderr()).as("IP Block Get by ID error message is not matching.")
                .contains("The request failed since this resource cannot be accessed by the provided credentials.");
        softly.assertAll();
    }
}
