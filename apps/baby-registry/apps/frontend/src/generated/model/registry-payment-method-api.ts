// This file is auto-generated. DO NOT EDIT.

import { RegistryPaymentMethod, RegistryPaymentMethodProjection } from './registry-payment-method-model';
import { Cart, CartProjection } from './cart-model';
import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { Registry, RegistryProjection } from './registry-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { PaymentMethodType } from './payment-method-type-enum';

export type RegistryPaymentMethodWithRefs = {
    registryPaymentMethod: RegistryPaymentMethod;
    carts?: Cart[];
    owner?: OwnerUser;
    registry?: Registry;
}

export type RegistryPaymentMethodWithRefsProjection = RegistryPaymentMethodProjection & {
    Carts?: CartProjection;
    Owner?: OwnerUserProjection;
    Registry?: RegistryProjection;
}

export type SelectRegistryPaymentMethodByIdQuery = {
    id: string;
}

export type RegistryPaymentMethodSearchQuery = {
    // id (Ref<RegistryPaymentMethod>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // bankAccountName (string) search options
    bankAccountNameEq?: string;
    bankAccountNameNe?: string;
    bankAccountNameGt?: string;
    bankAccountNameGte?: string;
    bankAccountNameLt?: string;
    bankAccountNameLte?: string;
    bankAccountNameIn?: string[];
    bankAccountNameNin?: string[];
    bankAccountNameExists?: boolean;
    bankAccountNameLike?: string;
    bankAccountNameNlike?: string;
    // bankAccountNumber (string) search options
    bankAccountNumberEq?: string;
    bankAccountNumberNe?: string;
    bankAccountNumberGt?: string;
    bankAccountNumberGte?: string;
    bankAccountNumberLt?: string;
    bankAccountNumberLte?: string;
    bankAccountNumberIn?: string[];
    bankAccountNumberNin?: string[];
    bankAccountNumberExists?: boolean;
    bankAccountNumberLike?: string;
    bankAccountNumberNlike?: string;
    // bankIban (string) search options
    bankIbanEq?: string;
    bankIbanNe?: string;
    bankIbanGt?: string;
    bankIbanGte?: string;
    bankIbanLt?: string;
    bankIbanLte?: string;
    bankIbanIn?: string[];
    bankIbanNin?: string[];
    bankIbanExists?: boolean;
    bankIbanLike?: string;
    bankIbanNlike?: string;
    // bankName (string) search options
    bankNameEq?: string;
    bankNameNe?: string;
    bankNameGt?: string;
    bankNameGte?: string;
    bankNameLt?: string;
    bankNameLte?: string;
    bankNameIn?: string[];
    bankNameNin?: string[];
    bankNameExists?: boolean;
    bankNameLike?: string;
    bankNameNlike?: string;
    // bankRoutingNumber (string) search options
    bankRoutingNumberEq?: string;
    bankRoutingNumberNe?: string;
    bankRoutingNumberGt?: string;
    bankRoutingNumberGte?: string;
    bankRoutingNumberLt?: string;
    bankRoutingNumberLte?: string;
    bankRoutingNumberIn?: string[];
    bankRoutingNumberNin?: string[];
    bankRoutingNumberExists?: boolean;
    bankRoutingNumberLike?: string;
    bankRoutingNumberNlike?: string;
    // bankSwift (string) search options
    bankSwiftEq?: string;
    bankSwiftNe?: string;
    bankSwiftGt?: string;
    bankSwiftGte?: string;
    bankSwiftLt?: string;
    bankSwiftLte?: string;
    bankSwiftIn?: string[];
    bankSwiftNin?: string[];
    bankSwiftExists?: boolean;
    bankSwiftLike?: string;
    bankSwiftNlike?: string;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // displayName (string) search options
    displayNameEq?: string;
    displayNameNe?: string;
    displayNameGt?: string;
    displayNameGte?: string;
    displayNameLt?: string;
    displayNameLte?: string;
    displayNameIn?: string[];
    displayNameNin?: string[];
    displayNameExists?: boolean;
    displayNameLike?: string;
    displayNameNlike?: string;
    // enabled (bool) search options
    enabledEq?: boolean;
    enabledNe?: boolean;
    enabledGt?: boolean;
    enabledGte?: boolean;
    enabledLt?: boolean;
    enabledLte?: boolean;
    enabledIn?: boolean[];
    enabledNin?: boolean[];
    enabledExists?: boolean;
    // instructions (string) search options
    instructionsEq?: string;
    instructionsNe?: string;
    instructionsGt?: string;
    instructionsGte?: string;
    instructionsLt?: string;
    instructionsLte?: string;
    instructionsIn?: string[];
    instructionsNin?: string[];
    instructionsExists?: boolean;
    instructionsLike?: string;
    instructionsNlike?: string;
    // ownerId (Ref<OwnerUser>) search options
    ownerIdEq?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    // paymentUrl (string) search options
    paymentUrlEq?: string;
    paymentUrlNe?: string;
    paymentUrlGt?: string;
    paymentUrlGte?: string;
    paymentUrlLt?: string;
    paymentUrlLte?: string;
    paymentUrlIn?: string[];
    paymentUrlNin?: string[];
    paymentUrlExists?: boolean;
    paymentUrlLike?: string;
    paymentUrlNlike?: string;
    // position (int) search options
    positionEq?: number;
    positionNe?: number;
    positionGt?: number;
    positionGte?: number;
    positionLt?: number;
    positionLte?: number;
    positionIn?: number[];
    positionNin?: number[];
    positionExists?: boolean;
    // recipientEmail (string) search options
    recipientEmailEq?: string;
    recipientEmailNe?: string;
    recipientEmailGt?: string;
    recipientEmailGte?: string;
    recipientEmailLt?: string;
    recipientEmailLte?: string;
    recipientEmailIn?: string[];
    recipientEmailNin?: string[];
    recipientEmailExists?: boolean;
    recipientEmailLike?: string;
    recipientEmailNlike?: string;
    // recipientPhone (string) search options
    recipientPhoneEq?: string;
    recipientPhoneNe?: string;
    recipientPhoneGt?: string;
    recipientPhoneGte?: string;
    recipientPhoneLt?: string;
    recipientPhoneLte?: string;
    recipientPhoneIn?: string[];
    recipientPhoneNin?: string[];
    recipientPhoneExists?: boolean;
    recipientPhoneLike?: string;
    recipientPhoneNlike?: string;
    // registryId (Ref<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
    // type (PaymentMethodType) search options
    typeEq?: PaymentMethodType;
    typeNe?: PaymentMethodType;
    typeGt?: PaymentMethodType;
    typeGte?: PaymentMethodType;
    typeLt?: PaymentMethodType;
    typeLte?: PaymentMethodType;
    typeIn?: PaymentMethodType[];
    typeNin?: PaymentMethodType[];
    typeExists?: boolean;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
