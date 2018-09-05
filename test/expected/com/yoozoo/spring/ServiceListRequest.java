// Code generated by protoapi; DO NOT EDIT.

package com.yoozoo.spring;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public class ServiceListRequest {
    private final List<Integer> tag_ids;
    private final int env_id;
    private final int offset;
    private final int limit;

    @JsonCreator
    public ServiceListRequest(@JsonProperty("tag_ids") List<Integer> tag_ids, @JsonProperty("env_id") int env_id, @JsonProperty("offset") int offset, @JsonProperty("limit") int limit) {
        this.tag_ids = tag_ids;
        this.env_id = env_id;
        this.offset = offset;
        this.limit = limit;
    }

    public List<Integer> getTag_ids() {
        return tag_ids;
    }
    public int getEnv_id() {
        return env_id;
    }
    public int getOffset() {
        return offset;
    }
    public int getLimit() {
        return limit;
    }
    
}
