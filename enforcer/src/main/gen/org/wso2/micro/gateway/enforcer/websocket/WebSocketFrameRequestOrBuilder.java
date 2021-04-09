// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/service/websocket/frame_service.proto

package org.wso2.micro.gateway.enforcer.websocket;

public interface WebSocketFrameRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:envoy.extensions.filters.http.mgw_wasm_websocket.v3.WebSocketFrameRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string node_id = 1;</code>
   * @return The nodeId.
   */
  java.lang.String getNodeId();
  /**
   * <code>string node_id = 1;</code>
   * @return The bytes for nodeId.
   */
  com.google.protobuf.ByteString
      getNodeIdBytes();

  /**
   * <code>.envoy.extensions.filters.http.mgw_wasm_websocket.v3.Metadata filter_metadata = 2;</code>
   * @return Whether the filterMetadata field is set.
   */
  boolean hasFilterMetadata();
  /**
   * <code>.envoy.extensions.filters.http.mgw_wasm_websocket.v3.Metadata filter_metadata = 2;</code>
   * @return The filterMetadata.
   */
  org.wso2.micro.gateway.enforcer.websocket.Metadata getFilterMetadata();
  /**
   * <code>.envoy.extensions.filters.http.mgw_wasm_websocket.v3.Metadata filter_metadata = 2;</code>
   */
  org.wso2.micro.gateway.enforcer.websocket.MetadataOrBuilder getFilterMetadataOrBuilder();

  /**
   * <code>int32 frame_length = 3;</code>
   * @return The frameLength.
   */
  int getFrameLength();

  /**
   * <code>string remote_ip = 4;</code>
   * @return The remoteIp.
   */
  java.lang.String getRemoteIp();
  /**
   * <code>string remote_ip = 4;</code>
   * @return The bytes for remoteIp.
   */
  com.google.protobuf.ByteString
      getRemoteIpBytes();
}
