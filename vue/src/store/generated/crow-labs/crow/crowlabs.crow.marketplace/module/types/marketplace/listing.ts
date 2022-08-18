/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "crowlabs.crow.marketplace";

export interface Order {
  id: number;
  userId: number;
  listingId: number;
  maxPrice: string;
  status: string;
}

export interface Listing {
  id: number;
  producerId: number;
  title: string;
  description: string;
  minPrice: string;
  status: string;
  pendingOrders: Order[];
  purchaseOrder: Order | undefined;
}

const baseOrder: object = {
  id: 0,
  userId: 0,
  listingId: 0,
  maxPrice: "",
  status: "",
};

export const Order = {
  encode(message: Order, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.userId !== 0) {
      writer.uint32(16).uint64(message.userId);
    }
    if (message.listingId !== 0) {
      writer.uint32(24).uint64(message.listingId);
    }
    if (message.maxPrice !== "") {
      writer.uint32(34).string(message.maxPrice);
    }
    if (message.status !== "") {
      writer.uint32(42).string(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrder } as Order;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.userId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.listingId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.maxPrice = reader.string();
          break;
        case 5:
          message.status = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order {
    const message = { ...baseOrder } as Order;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = Number(object.userId);
    } else {
      message.userId = 0;
    }
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = Number(object.listingId);
    } else {
      message.listingId = 0;
    }
    if (object.maxPrice !== undefined && object.maxPrice !== null) {
      message.maxPrice = String(object.maxPrice);
    } else {
      message.maxPrice = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    return message;
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.userId !== undefined && (obj.userId = message.userId);
    message.listingId !== undefined && (obj.listingId = message.listingId);
    message.maxPrice !== undefined && (obj.maxPrice = message.maxPrice);
    message.status !== undefined && (obj.status = message.status);
    return obj;
  },

  fromPartial(object: DeepPartial<Order>): Order {
    const message = { ...baseOrder } as Order;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.userId !== undefined && object.userId !== null) {
      message.userId = object.userId;
    } else {
      message.userId = 0;
    }
    if (object.listingId !== undefined && object.listingId !== null) {
      message.listingId = object.listingId;
    } else {
      message.listingId = 0;
    }
    if (object.maxPrice !== undefined && object.maxPrice !== null) {
      message.maxPrice = object.maxPrice;
    } else {
      message.maxPrice = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    return message;
  },
};

const baseListing: object = {
  id: 0,
  producerId: 0,
  title: "",
  description: "",
  minPrice: "",
  status: "",
};

export const Listing = {
  encode(message: Listing, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.producerId !== 0) {
      writer.uint32(16).uint64(message.producerId);
    }
    if (message.title !== "") {
      writer.uint32(26).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.minPrice !== "") {
      writer.uint32(42).string(message.minPrice);
    }
    if (message.status !== "") {
      writer.uint32(50).string(message.status);
    }
    for (const v of message.pendingOrders) {
      Order.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    if (message.purchaseOrder !== undefined) {
      Order.encode(message.purchaseOrder, writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Listing {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseListing } as Listing;
    message.pendingOrders = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.producerId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.title = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.minPrice = reader.string();
          break;
        case 6:
          message.status = reader.string();
          break;
        case 7:
          message.pendingOrders.push(Order.decode(reader, reader.uint32()));
          break;
        case 8:
          message.purchaseOrder = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Listing {
    const message = { ...baseListing } as Listing;
    message.pendingOrders = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = Number(object.producerId);
    } else {
      message.producerId = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.minPrice !== undefined && object.minPrice !== null) {
      message.minPrice = String(object.minPrice);
    } else {
      message.minPrice = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.pendingOrders !== undefined && object.pendingOrders !== null) {
      for (const e of object.pendingOrders) {
        message.pendingOrders.push(Order.fromJSON(e));
      }
    }
    if (object.purchaseOrder !== undefined && object.purchaseOrder !== null) {
      message.purchaseOrder = Order.fromJSON(object.purchaseOrder);
    } else {
      message.purchaseOrder = undefined;
    }
    return message;
  },

  toJSON(message: Listing): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.producerId !== undefined && (obj.producerId = message.producerId);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined &&
      (obj.description = message.description);
    message.minPrice !== undefined && (obj.minPrice = message.minPrice);
    message.status !== undefined && (obj.status = message.status);
    if (message.pendingOrders) {
      obj.pendingOrders = message.pendingOrders.map((e) =>
        e ? Order.toJSON(e) : undefined
      );
    } else {
      obj.pendingOrders = [];
    }
    message.purchaseOrder !== undefined &&
      (obj.purchaseOrder = message.purchaseOrder
        ? Order.toJSON(message.purchaseOrder)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Listing>): Listing {
    const message = { ...baseListing } as Listing;
    message.pendingOrders = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.producerId !== undefined && object.producerId !== null) {
      message.producerId = object.producerId;
    } else {
      message.producerId = 0;
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.minPrice !== undefined && object.minPrice !== null) {
      message.minPrice = object.minPrice;
    } else {
      message.minPrice = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.pendingOrders !== undefined && object.pendingOrders !== null) {
      for (const e of object.pendingOrders) {
        message.pendingOrders.push(Order.fromPartial(e));
      }
    }
    if (object.purchaseOrder !== undefined && object.purchaseOrder !== null) {
      message.purchaseOrder = Order.fromPartial(object.purchaseOrder);
    } else {
      message.purchaseOrder = undefined;
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
