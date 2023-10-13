/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "mercury.api.v2";

export interface SystemInfo {
  version: string;
  mode: string;
  allowRegistration: boolean;
  disablePasswordLogin: boolean;
  additionalScript: string;
  additionalStyle: string;
  dbSize: number;
}

export interface GetSystemInfoRequest {
}

export interface GetSystemInfoResponse {
  systemInfo?: SystemInfo | undefined;
}

export interface UpdateSystemInfoRequest {
  systemInfo?: SystemInfo | undefined;
  updateMask: string[];
}

export interface UpdateSystemInfoResponse {
  systemInfo?: SystemInfo | undefined;
}

function createBaseSystemInfo(): SystemInfo {
  return {
    version: "",
    mode: "",
    allowRegistration: false,
    disablePasswordLogin: false,
    additionalScript: "",
    additionalStyle: "",
    dbSize: 0,
  };
}

export const SystemInfo = {
  encode(message: SystemInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.version !== "") {
      writer.uint32(10).string(message.version);
    }
    if (message.mode !== "") {
      writer.uint32(18).string(message.mode);
    }
    if (message.allowRegistration === true) {
      writer.uint32(24).bool(message.allowRegistration);
    }
    if (message.disablePasswordLogin === true) {
      writer.uint32(32).bool(message.disablePasswordLogin);
    }
    if (message.additionalScript !== "") {
      writer.uint32(42).string(message.additionalScript);
    }
    if (message.additionalStyle !== "") {
      writer.uint32(50).string(message.additionalStyle);
    }
    if (message.dbSize !== 0) {
      writer.uint32(56).int64(message.dbSize);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SystemInfo {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSystemInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.version = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.mode = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.allowRegistration = reader.bool();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.disablePasswordLogin = reader.bool();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.additionalScript = reader.string();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.additionalStyle = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.dbSize = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SystemInfo {
    return {
      version: isSet(object.version) ? String(object.version) : "",
      mode: isSet(object.mode) ? String(object.mode) : "",
      allowRegistration: isSet(object.allowRegistration) ? Boolean(object.allowRegistration) : false,
      disablePasswordLogin: isSet(object.disablePasswordLogin) ? Boolean(object.disablePasswordLogin) : false,
      additionalScript: isSet(object.additionalScript) ? String(object.additionalScript) : "",
      additionalStyle: isSet(object.additionalStyle) ? String(object.additionalStyle) : "",
      dbSize: isSet(object.dbSize) ? Number(object.dbSize) : 0,
    };
  },

  toJSON(message: SystemInfo): unknown {
    const obj: any = {};
    message.version !== undefined && (obj.version = message.version);
    message.mode !== undefined && (obj.mode = message.mode);
    message.allowRegistration !== undefined && (obj.allowRegistration = message.allowRegistration);
    message.disablePasswordLogin !== undefined && (obj.disablePasswordLogin = message.disablePasswordLogin);
    message.additionalScript !== undefined && (obj.additionalScript = message.additionalScript);
    message.additionalStyle !== undefined && (obj.additionalStyle = message.additionalStyle);
    message.dbSize !== undefined && (obj.dbSize = Math.round(message.dbSize));
    return obj;
  },

  create(base?: DeepPartial<SystemInfo>): SystemInfo {
    return SystemInfo.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<SystemInfo>): SystemInfo {
    const message = createBaseSystemInfo();
    message.version = object.version ?? "";
    message.mode = object.mode ?? "";
    message.allowRegistration = object.allowRegistration ?? false;
    message.disablePasswordLogin = object.disablePasswordLogin ?? false;
    message.additionalScript = object.additionalScript ?? "";
    message.additionalStyle = object.additionalStyle ?? "";
    message.dbSize = object.dbSize ?? 0;
    return message;
  },
};

function createBaseGetSystemInfoRequest(): GetSystemInfoRequest {
  return {};
}

export const GetSystemInfoRequest = {
  encode(_: GetSystemInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSystemInfoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetSystemInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): GetSystemInfoRequest {
    return {};
  },

  toJSON(_: GetSystemInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<GetSystemInfoRequest>): GetSystemInfoRequest {
    return GetSystemInfoRequest.fromPartial(base ?? {});
  },

  fromPartial(_: DeepPartial<GetSystemInfoRequest>): GetSystemInfoRequest {
    const message = createBaseGetSystemInfoRequest();
    return message;
  },
};

function createBaseGetSystemInfoResponse(): GetSystemInfoResponse {
  return { systemInfo: undefined };
}

export const GetSystemInfoResponse = {
  encode(message: GetSystemInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.systemInfo !== undefined) {
      SystemInfo.encode(message.systemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetSystemInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetSystemInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.systemInfo = SystemInfo.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetSystemInfoResponse {
    return { systemInfo: isSet(object.systemInfo) ? SystemInfo.fromJSON(object.systemInfo) : undefined };
  },

  toJSON(message: GetSystemInfoResponse): unknown {
    const obj: any = {};
    message.systemInfo !== undefined &&
      (obj.systemInfo = message.systemInfo ? SystemInfo.toJSON(message.systemInfo) : undefined);
    return obj;
  },

  create(base?: DeepPartial<GetSystemInfoResponse>): GetSystemInfoResponse {
    return GetSystemInfoResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<GetSystemInfoResponse>): GetSystemInfoResponse {
    const message = createBaseGetSystemInfoResponse();
    message.systemInfo = (object.systemInfo !== undefined && object.systemInfo !== null)
      ? SystemInfo.fromPartial(object.systemInfo)
      : undefined;
    return message;
  },
};

function createBaseUpdateSystemInfoRequest(): UpdateSystemInfoRequest {
  return { systemInfo: undefined, updateMask: [] };
}

export const UpdateSystemInfoRequest = {
  encode(message: UpdateSystemInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.systemInfo !== undefined) {
      SystemInfo.encode(message.systemInfo, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.updateMask) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateSystemInfoRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateSystemInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.systemInfo = SystemInfo.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.updateMask.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateSystemInfoRequest {
    return {
      systemInfo: isSet(object.systemInfo) ? SystemInfo.fromJSON(object.systemInfo) : undefined,
      updateMask: Array.isArray(object?.updateMask) ? object.updateMask.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: UpdateSystemInfoRequest): unknown {
    const obj: any = {};
    message.systemInfo !== undefined &&
      (obj.systemInfo = message.systemInfo ? SystemInfo.toJSON(message.systemInfo) : undefined);
    if (message.updateMask) {
      obj.updateMask = message.updateMask.map((e) => e);
    } else {
      obj.updateMask = [];
    }
    return obj;
  },

  create(base?: DeepPartial<UpdateSystemInfoRequest>): UpdateSystemInfoRequest {
    return UpdateSystemInfoRequest.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpdateSystemInfoRequest>): UpdateSystemInfoRequest {
    const message = createBaseUpdateSystemInfoRequest();
    message.systemInfo = (object.systemInfo !== undefined && object.systemInfo !== null)
      ? SystemInfo.fromPartial(object.systemInfo)
      : undefined;
    message.updateMask = object.updateMask?.map((e) => e) || [];
    return message;
  },
};

function createBaseUpdateSystemInfoResponse(): UpdateSystemInfoResponse {
  return { systemInfo: undefined };
}

export const UpdateSystemInfoResponse = {
  encode(message: UpdateSystemInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.systemInfo !== undefined) {
      SystemInfo.encode(message.systemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateSystemInfoResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateSystemInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.systemInfo = SystemInfo.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateSystemInfoResponse {
    return { systemInfo: isSet(object.systemInfo) ? SystemInfo.fromJSON(object.systemInfo) : undefined };
  },

  toJSON(message: UpdateSystemInfoResponse): unknown {
    const obj: any = {};
    message.systemInfo !== undefined &&
      (obj.systemInfo = message.systemInfo ? SystemInfo.toJSON(message.systemInfo) : undefined);
    return obj;
  },

  create(base?: DeepPartial<UpdateSystemInfoResponse>): UpdateSystemInfoResponse {
    return UpdateSystemInfoResponse.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<UpdateSystemInfoResponse>): UpdateSystemInfoResponse {
    const message = createBaseUpdateSystemInfoResponse();
    message.systemInfo = (object.systemInfo !== undefined && object.systemInfo !== null)
      ? SystemInfo.fromPartial(object.systemInfo)
      : undefined;
    return message;
  },
};

export type SystemServiceDefinition = typeof SystemServiceDefinition;
export const SystemServiceDefinition = {
  name: "SystemService",
  fullName: "mercury.api.v2.SystemService",
  methods: {
    getSystemInfo: {
      name: "GetSystemInfo",
      requestType: GetSystemInfoRequest,
      requestStream: false,
      responseType: GetSystemInfoResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              18,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              115,
              121,
              115,
              116,
              101,
              109,
              47,
              105,
              110,
              102,
              111,
            ]),
          ],
        },
      },
    },
    updateSystemInfo: {
      name: "UpdateSystemInfo",
      requestType: UpdateSystemInfoRequest,
      requestStream: false,
      responseType: UpdateSystemInfoResponse,
      responseStream: false,
      options: {
        _unknownFields: {
          578365826: [
            new Uint8Array([
              21,
              34,
              19,
              47,
              97,
              112,
              105,
              47,
              118,
              50,
              47,
              115,
              121,
              115,
              116,
              101,
              109,
              47,
              105,
              110,
              102,
              111,
            ]),
          ],
        },
      },
    },
  },
} as const;

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
