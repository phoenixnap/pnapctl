package com.pnapctl.tests.utils;

import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;

import ch.qos.logback.classic.pattern.ClassicConverter;
import ch.qos.logback.classic.spi.ILoggingEvent;

public class ThreadIdConverter extends ClassicConverter {

    private static final ConcurrentHashMap<String, String> map = new ConcurrentHashMap<>();
    private static final AtomicInteger counter = new AtomicInteger(1);

    @Override
    public String convert(ILoggingEvent event) {
        String threadName = Thread.currentThread().getName();
        return map.computeIfAbsent(threadName, _ -> String.format("%03d", counter.getAndIncrement()));
    }
}