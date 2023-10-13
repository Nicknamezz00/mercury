/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { MethodPolicy } from "./policy";

export const protobufPackage = "google.api";

/**
 * Selects and configures the service controller used by the service.
 *
 * Example:
 *
 *     control:
 *       environment: servicecontrol.googleapis.com
 */
export interface Control {
  /**
   * The service controller environment to use. If empty, no control plane
   * feature (like quota and billing) will be enabled. The recommended value for
   * most services is servicecontrol.googleapis.com
   */
  environment: string;
  /** Defines policies applying to the API methods of the service. */
  methodPolicies: MethodPolicy[];
}

function createBaseControl(): Control {
  return { environment: "", methodPolicies: [] };
}

export const Control = {
  encode(message: Control, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.environment !== "") {
      writer.uint32(10).string(message.environment);
    }
    for (const v of message.methodPolicies) {
      MethodPolicy.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Control {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseControl();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.environment = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.methodPolicies.push(MethodPolicy.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Control {
    return {
      environment: isSet(object.environment) ? String(object.environment) : "",
      methodPolicies: Array.isArray(object?.methodPolicies)
        ? object.methodPolicies.map((e: any) => MethodPolicy.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Control): unknown {
    const obj: any = {};
    message.environment !== undefined && (obj.environment = message.environment);
    if (message.methodPolicies) {
      obj.methodPolicies = message.methodPolicies.map((e) => e ? MethodPolicy.toJSON(e) : undefined);
    } else {
      obj.methodPolicies = [];
    }
    return obj;
  },

  create(base?: DeepPartial<Control>): Control {
    return Control.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<Control>): Control {
    const message = createBaseControl();
    message.environment = object.environment ?? "";
    message.methodPolicies = object.methodPolicies?.map((e) => MethodPolicy.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
