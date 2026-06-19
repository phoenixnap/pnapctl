package com.pnapctl.tests.datasources.executor.tag;

import static com.pnapctl.tests.datasources.enums.MethodEnum.CREATE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.DELETE;
import static com.pnapctl.tests.datasources.enums.MethodEnum.GET;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.TAG;
import static com.pnapctl.tests.datasources.enums.ResourceEnum.TAGS;

import com.pnapctl.tests.datasources.executor.PnapctlCommand;
import com.pnapctl.tests.datasources.executor.PnapctlResource;

public class PnapctlTags extends PnapctlResource {

    public PnapctlCommand get() {
        return command(GET.value(), TAGS.value());
    }

    public PnapctlCommand getById(String id) {
        return command(GET.value(), TAG.value(), id);
    }

    public PnapctlCommand deleteById(String id) {
        return command(DELETE.value(), TAG.value(), id);
    }

    public PnapctlCommand create() {
        return command(CREATE.value(), TAG.value());
    }
}
