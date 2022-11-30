// package: tracking
// file: tracking.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

export class Request extends jspb.Message { 
    getNumber(): string;
    setNumber(value: string): Request;
    getScac(): string;
    setScac(value: string): Request;
    getCountry(): Country;
    setCountry(value: Country): Request;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
    export type AsObject = {
        number: string,
        scac: string,
        country: Country,
    }
}

export class InfoAboutMoving extends jspb.Message { 
    getTime(): number;
    setTime(value: number): InfoAboutMoving;
    getOperationName(): string;
    setOperationName(value: string): InfoAboutMoving;
    getLocation(): string;
    setLocation(value: string): InfoAboutMoving;
    getVessel(): string;
    setVessel(value: string): InfoAboutMoving;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InfoAboutMoving.AsObject;
    static toObject(includeInstance: boolean, msg: InfoAboutMoving): InfoAboutMoving.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InfoAboutMoving, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InfoAboutMoving;
    static deserializeBinaryFromReader(message: InfoAboutMoving, reader: jspb.BinaryReader): InfoAboutMoving;
}

export namespace InfoAboutMoving {
    export type AsObject = {
        time: number,
        operationName: string,
        location: string,
        vessel: string,
    }
}

export class TrackingByContainerNumberResponse extends jspb.Message { 
    getContainer(): string;
    setContainer(value: string): TrackingByContainerNumberResponse;
    getContainerSize(): string;
    setContainerSize(value: string): TrackingByContainerNumberResponse;
    getScac(): string;
    setScac(value: string): TrackingByContainerNumberResponse;
    clearInfoAboutMovingList(): void;
    getInfoAboutMovingList(): Array<InfoAboutMoving>;
    setInfoAboutMovingList(value: Array<InfoAboutMoving>): TrackingByContainerNumberResponse;
    addInfoAboutMoving(value?: InfoAboutMoving, index?: number): InfoAboutMoving;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TrackingByContainerNumberResponse.AsObject;
    static toObject(includeInstance: boolean, msg: TrackingByContainerNumberResponse): TrackingByContainerNumberResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TrackingByContainerNumberResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TrackingByContainerNumberResponse;
    static deserializeBinaryFromReader(message: TrackingByContainerNumberResponse, reader: jspb.BinaryReader): TrackingByContainerNumberResponse;
}

export namespace TrackingByContainerNumberResponse {
    export type AsObject = {
        container: string,
        containerSize: string,
        scac: string,
        infoAboutMovingList: Array<InfoAboutMoving.AsObject>,
    }
}

export class TrackingByBillNumberResponse extends jspb.Message { 
    getBillno(): string;
    setBillno(value: string): TrackingByBillNumberResponse;
    getScac(): string;
    setScac(value: string): TrackingByBillNumberResponse;
    clearInfoAboutMovingList(): void;
    getInfoAboutMovingList(): Array<InfoAboutMoving>;
    setInfoAboutMovingList(value: Array<InfoAboutMoving>): TrackingByBillNumberResponse;
    addInfoAboutMoving(value?: InfoAboutMoving, index?: number): InfoAboutMoving;
    getEtaFinalDelivery(): number;
    setEtaFinalDelivery(value: number): TrackingByBillNumberResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TrackingByBillNumberResponse.AsObject;
    static toObject(includeInstance: boolean, msg: TrackingByBillNumberResponse): TrackingByBillNumberResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TrackingByBillNumberResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TrackingByBillNumberResponse;
    static deserializeBinaryFromReader(message: TrackingByBillNumberResponse, reader: jspb.BinaryReader): TrackingByBillNumberResponse;
}

export namespace TrackingByBillNumberResponse {
    export type AsObject = {
        billno: string,
        scac: string,
        infoAboutMovingList: Array<InfoAboutMoving.AsObject>,
        etaFinalDelivery: number,
    }
}

export class Scac extends jspb.Message { 
    getScac(): string;
    setScac(value: string): Scac;
    getFullname(): string;
    setFullname(value: string): Scac;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Scac.AsObject;
    static toObject(includeInstance: boolean, msg: Scac): Scac.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Scac, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Scac;
    static deserializeBinaryFromReader(message: Scac, reader: jspb.BinaryReader): Scac;
}

export namespace Scac {
    export type AsObject = {
        scac: string,
        fullname: string,
    }
}

export class GetAllScacResponse extends jspb.Message { 
    clearAllScacList(): void;
    getAllScacList(): Array<Scac>;
    setAllScacList(value: Array<Scac>): GetAllScacResponse;
    addAllScac(value?: Scac, index?: number): Scac;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetAllScacResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetAllScacResponse): GetAllScacResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetAllScacResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetAllScacResponse;
    static deserializeBinaryFromReader(message: GetAllScacResponse, reader: jspb.BinaryReader): GetAllScacResponse;
}

export namespace GetAllScacResponse {
    export type AsObject = {
        allScacList: Array<Scac.AsObject>,
    }
}

export enum Country {
    RU = 0,
    OTHER = 1,
}
