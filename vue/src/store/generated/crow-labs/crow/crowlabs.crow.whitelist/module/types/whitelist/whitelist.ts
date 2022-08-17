/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crowlabs.crow.whitelist";

export interface User {
  id: number;
  address: string;
  status: string;
  pendingOrderIds: number[];
  previousOrderIds: number[];
}

export interface Producer {
  id: number;
  address: string;
  status: string;
  listingIds: number[];
  pendingOrderIds: number[];
  previousOrderIds: number[];
}

export interface Whitelist {
  users: User[];
  producers: Producer[];
}

const baseUser: object = {
  id: 0,
  address: "",
  status: "",
  pendingOrderIds: 0,
  previousOrderIds: 0,
};

export const User = {
  encode(message: User, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    writer.uint32(34).fork();
    for (const v of message.pendingOrderIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(42).fork();
    for (const v of message.previousOrderIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): User {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUser } as User;
    message.pendingOrderIds = [];
    message.previousOrderIds = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.pendingOrderIds.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.pendingOrderIds.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 5:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.previousOrderIds.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.previousOrderIds.push(
              longToNumber(reader.uint64() as Long)
            );
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    const message = { ...baseUser } as User;
    message.pendingOrderIds = [];
    message.previousOrderIds = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (
      object.pendingOrderIds !== undefined &&
      object.pendingOrderIds !== null
    ) {
      for (const e of object.pendingOrderIds) {
        message.pendingOrderIds.push(Number(e));
      }
    }
    if (
      object.previousOrderIds !== undefined &&
      object.previousOrderIds !== null
    ) {
      for (const e of object.previousOrderIds) {
        message.previousOrderIds.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    message.status !== undefined && (obj.status = message.status);
    if (message.pendingOrderIds) {
      obj.pendingOrderIds = message.pendingOrderIds.map((e) => e);
    } else {
      obj.pendingOrderIds = [];
    }
    if (message.previousOrderIds) {
      obj.previousOrderIds = message.previousOrderIds.map((e) => e);
    } else {
      obj.previousOrderIds = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<User>): User {
    const message = { ...baseUser } as User;
    message.pendingOrderIds = [];
    message.previousOrderIds = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (
      object.pendingOrderIds !== undefined &&
      object.pendingOrderIds !== null
    ) {
      for (const e of object.pendingOrderIds) {
        message.pendingOrderIds.push(e);
      }
    }
    if (
      object.previousOrderIds !== undefined &&
      object.previousOrderIds !== null
    ) {
      for (const e of object.previousOrderIds) {
        message.previousOrderIds.push(e);
      }
    }
    return message;
  },
};

const baseProducer: object = {
  id: 0,
  address: "",
  status: "",
  listingIds: 0,
  pendingOrderIds: 0,
  previousOrderIds: 0,
};

export const Producer = {
  encode(message: Producer, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    writer.uint32(34).fork();
    for (const v of message.listingIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(42).fork();
    for (const v of message.pendingOrderIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(50).fork();
    for (const v of message.previousOrderIds) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Producer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProducer } as Producer;
    message.listingIds = [];
    message.pendingOrderIds = [];
    message.previousOrderIds = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.listingIds.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.listingIds.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 5:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.pendingOrderIds.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.pendingOrderIds.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 6:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.previousOrderIds.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.previousOrderIds.push(
              longToNumber(reader.uint64() as Long)
            );
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Producer {
    const message = { ...baseProducer } as Producer;
    message.listingIds = [];
    message.pendingOrderIds = [];
    message.previousOrderIds = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.listingIds !== undefined && object.listingIds !== null) {
      for (const e of object.listingIds) {
        message.listingIds.push(Number(e));
      }
    }
    if (
      object.pendingOrderIds !== undefined &&
      object.pendingOrderIds !== null
    ) {
      for (const e of object.pendingOrderIds) {
        message.pendingOrderIds.push(Number(e));
      }
    }
    if (
      object.previousOrderIds !== undefined &&
      object.previousOrderIds !== null
    ) {
      for (const e of object.previousOrderIds) {
        message.previousOrderIds.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: Producer): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    message.status !== undefined && (obj.status = message.status);
    if (message.listingIds) {
      obj.listingIds = message.listingIds.map((e) => e);
    } else {
      obj.listingIds = [];
    }
    if (message.pendingOrderIds) {
      obj.pendingOrderIds = message.pendingOrderIds.map((e) => e);
    } else {
      obj.pendingOrderIds = [];
    }
    if (message.previousOrderIds) {
      obj.previousOrderIds = message.previousOrderIds.map((e) => e);
    } else {
      obj.previousOrderIds = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Producer>): Producer {
    const message = { ...baseProducer } as Producer;
    message.listingIds = [];
    message.pendingOrderIds = [];
    message.previousOrderIds = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.listingIds !== undefined && object.listingIds !== null) {
      for (const e of object.listingIds) {
        message.listingIds.push(e);
      }
    }
    if (
      object.pendingOrderIds !== undefined &&
      object.pendingOrderIds !== null
    ) {
      for (const e of object.pendingOrderIds) {
        message.pendingOrderIds.push(e);
      }
    }
    if (
      object.previousOrderIds !== undefined &&
      object.previousOrderIds !== null
    ) {
      for (const e of object.previousOrderIds) {
        message.previousOrderIds.push(e);
      }
    }
    return message;
  },
};

const baseWhitelist: object = {};

export const Whitelist = {
  encode(message: Whitelist, writer: Writer = Writer.create()): Writer {
    for (const v of message.users) {
      User.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.producers) {
      Producer.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Whitelist {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWhitelist } as Whitelist;
    message.users = [];
    message.producers = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.users.push(User.decode(reader, reader.uint32()));
          break;
        case 2:
          message.producers.push(Producer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Whitelist {
    const message = { ...baseWhitelist } as Whitelist;
    message.users = [];
    message.producers = [];
    if (object.users !== undefined && object.users !== null) {
      for (const e of object.users) {
        message.users.push(User.fromJSON(e));
      }
    }
    if (object.producers !== undefined && object.producers !== null) {
      for (const e of object.producers) {
        message.producers.push(Producer.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Whitelist): unknown {
    const obj: any = {};
    if (message.users) {
      obj.users = message.users.map((e) => (e ? User.toJSON(e) : undefined));
    } else {
      obj.users = [];
    }
    if (message.producers) {
      obj.producers = message.producers.map((e) =>
        e ? Producer.toJSON(e) : undefined
      );
    } else {
      obj.producers = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Whitelist>): Whitelist {
    const message = { ...baseWhitelist } as Whitelist;
    message.users = [];
    message.producers = [];
    if (object.users !== undefined && object.users !== null) {
      for (const e of object.users) {
        message.users.push(User.fromPartial(e));
      }
    }
    if (object.producers !== undefined && object.producers !== null) {
      for (const e of object.producers) {
        message.producers.push(Producer.fromPartial(e));
      }
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
