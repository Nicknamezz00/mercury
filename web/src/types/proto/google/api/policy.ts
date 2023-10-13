/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "google.api";

/**
 * Google API Policy Annotation
 *
 * This message defines a simple API policy annotation that can be used to
 * annotate API request and response message fields with applicable policies.
 * One field may have multiple applicable policies that must all be satisfied
 * before a request can be processed. This policy annotation is used to
 * generate the overall policy that will be used for automatic runtime
 * policy enforcement and documentation generation.
 */
export interface FieldPolicy {
  /**
   * Selects one or more request or response message fields to apply this
   * `FieldPolicy`.
   *
   * When a `FieldPolicy` is used in proto annotation, the selector must
   * be left as empty. The service config generator will automatically fill
   * the correct value.
   *
   * When a `FieldPolicy` is used in service config, the selector must be a
   * comma-separated string with valid request or response field paths,
   * such as "foo.bar" or "foo.bar,foo.baz".
   */
  selector: string;
  /**
   * Specifies the required permission(s) for the resource referred to by the
   * field. It requires the field contains a valid resource reference, and
   * the request must pass the permission checks to proceed. For example,
   * "resourcemanager.projects.get".
   */
  resourcePermission: string;
  /** Specifies the resource type for the resource referred to by the field. */
  resourceType: string;
}

/** Defines policies applying to an RPC method. */
export interface MethodPolicy {
  /**
   * Selects a method to which these policies should be enforced, for example,
   * "google.pubsub.v1.Subscriber.CreateSubscription".
   *
   * Refer to [selector][google.api.DocumentationRule.selector] for syntax
   * details.
   *
   * NOTE: This field must not be set in the proto annotation. It will be
   * automatically filled by the service config compiler .
   */
  selector: string;
  /** Policies that are applicable to the request message. */
  requestPolicies: FieldPolicy[];
}

function createBaseFieldPolicy(): FieldPolicy {
  return { selector: "", resourcePermission: "", resourceType: "" };
}

export const FieldPolicy = {
  encode(message: FieldPolicy, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.selector !== "") {
      writer.uint32(10).string(message.selector);
    }
    if (message.resourcePermission !== "") {
      writer.uint32(18).string(message.resourcePermission);
    }
    if (message.resourceType !== "") {
      writer.uint32(26).string(message.resourceType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FieldPolicy {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFieldPolicy();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.selector = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.resourcePermission = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.resourceType = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FieldPolicy {
    return {
      selector: isSet(object.selector) ? String(object.selector) : "",
      resourcePermission: isSet(object.resourcePermission) ? String(object.resourcePermission) : "",
      resourceType: isSet(object.resourceType) ? String(object.resourceType) : "",
    };
  },

  toJSON(message: FieldPolicy): unknown {
    const obj: any = {};
    message.selector !== undefined && (obj.selector = message.selector);
    message.resourcePermission !== undefined && (obj.resourcePermission = message.resourcePermission);
    message.resourceType !== undefined && (obj.resourceType = message.resourceType);
    return obj;
  },

  create(base?: DeepPartial<FieldPolicy>): FieldPolicy {
    return FieldPolicy.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<FieldPolicy>): FieldPolicy {
    const message = createBaseFieldPolicy();
    message.selector = object.selector ?? "";
    message.resourcePermission = object.resourcePermission ?? "";
    message.resourceType = object.resourceType ?? "";
    return message;
  },
};

function createBaseMethodPolicy(): MethodPolicy {
  return { selector: "", requestPolicies: [] };
}

export const MethodPolicy = {
  encode(message: MethodPolicy, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.selector !== "") {
      writer.uint32(74).string(message.selector);
    }
    for (const v of message.requestPolicies) {
      FieldPolicy.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MethodPolicy {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMethodPolicy();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 9:
          if (tag !== 74) {
            break;
          }

          message.selector = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.requestPolicies.push(FieldPolicy.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MethodPolicy {
    return {
      selector: isSet(object.selector) ? String(object.selector) : "",
      requestPolicies: Array.isArray(object?.requestPolicies)
        ? object.requestPolicies.map((e: any) => FieldPolicy.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MethodPolicy): unknown {
    const obj: any = {};
    message.selector !== undefined && (obj.selector = message.selector);
    if (message.requestPolicies) {
      obj.requestPolicies = message.requestPolicies.map((e) => e ? FieldPolicy.toJSON(e) : undefined);
    } else {
      obj.requestPolicies = [];
    }
    return obj;
  },

  create(base?: DeepPartial<MethodPolicy>): MethodPolicy {
    return MethodPolicy.fromPartial(base ?? {});
  },

  fromPartial(object: DeepPartial<MethodPolicy>): MethodPolicy {
    const message = createBaseMethodPolicy();
    message.selector = object.selector ?? "";
    message.requestPolicies = object.requestPolicies?.map((e) => FieldPolicy.fromPartial(e)) || [];
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
