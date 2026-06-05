// This file is auto-generated. DO NOT EDIT.

import { Cart, CartProjection } from './cart-model';
import { OwnerUser, OwnerUserProjection } from './owner-user-model';
import { Registry, RegistryProjection } from './registry-model';
import { RegistryPaymentMethod, RegistryPaymentMethodProjection } from './registry-payment-method-model';
import { Reservation, ReservationProjection } from './reservation-model';
import { ActorTraceSearchQuery } from './actor-trace-api';
import { CartStatus } from './cart-status-enum';
import { PaymentMethodType } from './payment-method-type-enum';

export type CartWithRefs = {
    cart: Cart;
    reservations?: Reservation[];
    owner?: OwnerUser;
    paymentMethod?: RegistryPaymentMethod;
    registry?: Registry;
}

export type CartWithRefsProjection = CartProjection & {
    Reservations?: ReservationProjection;
    Owner?: OwnerUserProjection;
    PaymentMethod?: RegistryPaymentMethodProjection;
    Registry?: RegistryProjection;
}

export type SelectCartByIdQuery = {
    id: string;
}

export type SelectCartByReferenceUniqueQuery = {
    referenceCode: string;
}

export type CartSearchQuery = {
    // id (Ref<Cart>) search options
    idEq?: string;
    idIn?: string[];
    idNin?: string[];
    idExists?: boolean;
    // amountCents (int) search options
    amountCentsEq?: number;
    amountCentsNe?: number;
    amountCentsGt?: number;
    amountCentsGte?: number;
    amountCentsLt?: number;
    amountCentsLte?: number;
    amountCentsIn?: number[];
    amountCentsNin?: number[];
    amountCentsExists?: boolean;
    // claimedAt (timestamp) search options
    claimedAtEq?: string;
    claimedAtNe?: string;
    claimedAtGt?: string;
    claimedAtGte?: string;
    claimedAtLt?: string;
    claimedAtLte?: string;
    claimedAtIn?: string[];
    claimedAtNin?: string[];
    claimedAtExists?: boolean;
    // contributorEmail (string) search options
    contributorEmailEq?: string;
    contributorEmailNe?: string;
    contributorEmailGt?: string;
    contributorEmailGte?: string;
    contributorEmailLt?: string;
    contributorEmailLte?: string;
    contributorEmailIn?: string[];
    contributorEmailNin?: string[];
    contributorEmailExists?: boolean;
    contributorEmailLike?: string;
    contributorEmailNlike?: string;
    // contributorName (string) search options
    contributorNameEq?: string;
    contributorNameNe?: string;
    contributorNameGt?: string;
    contributorNameGte?: string;
    contributorNameLt?: string;
    contributorNameLte?: string;
    contributorNameIn?: string[];
    contributorNameNin?: string[];
    contributorNameExists?: boolean;
    contributorNameLike?: string;
    contributorNameNlike?: string;
    // created (ActorTrace) search options
    created?: ActorTraceSearchQuery;
    // currency (string) search options
    currencyEq?: string;
    currencyNe?: string;
    currencyGt?: string;
    currencyGte?: string;
    currencyLt?: string;
    currencyLte?: string;
    currencyIn?: string[];
    currencyNin?: string[];
    currencyExists?: boolean;
    currencyLike?: string;
    currencyNlike?: string;
    // decidedAt (timestamp) search options
    decidedAtEq?: string;
    decidedAtNe?: string;
    decidedAtGt?: string;
    decidedAtGte?: string;
    decidedAtLt?: string;
    decidedAtLte?: string;
    decidedAtIn?: string[];
    decidedAtNin?: string[];
    decidedAtExists?: boolean;
    // decisionReason (string) search options
    decisionReasonEq?: string;
    decisionReasonNe?: string;
    decisionReasonGt?: string;
    decisionReasonGte?: string;
    decisionReasonLt?: string;
    decisionReasonLte?: string;
    decisionReasonIn?: string[];
    decisionReasonNin?: string[];
    decisionReasonExists?: boolean;
    decisionReasonLike?: string;
    decisionReasonNlike?: string;
    // message (string) search options
    messageEq?: string;
    messageNe?: string;
    messageGt?: string;
    messageGte?: string;
    messageLt?: string;
    messageLte?: string;
    messageIn?: string[];
    messageNin?: string[];
    messageExists?: boolean;
    messageLike?: string;
    messageNlike?: string;
    // methodDisplayName (string) search options
    methodDisplayNameEq?: string;
    methodDisplayNameNe?: string;
    methodDisplayNameGt?: string;
    methodDisplayNameGte?: string;
    methodDisplayNameLt?: string;
    methodDisplayNameLte?: string;
    methodDisplayNameIn?: string[];
    methodDisplayNameNin?: string[];
    methodDisplayNameExists?: boolean;
    methodDisplayNameLike?: string;
    methodDisplayNameNlike?: string;
    // methodType (PaymentMethodType) search options
    methodTypeEq?: PaymentMethodType;
    methodTypeNe?: PaymentMethodType;
    methodTypeGt?: PaymentMethodType;
    methodTypeGte?: PaymentMethodType;
    methodTypeLt?: PaymentMethodType;
    methodTypeLte?: PaymentMethodType;
    methodTypeIn?: PaymentMethodType[];
    methodTypeNin?: PaymentMethodType[];
    methodTypeExists?: boolean;
    // ownerId (Ref<OwnerUser>) search options
    ownerIdEq?: string;
    ownerIdIn?: string[];
    ownerIdNin?: string[];
    ownerIdExists?: boolean;
    // paymentMethodId (Ref<RegistryPaymentMethod>) search options
    paymentMethodIdEq?: string;
    paymentMethodIdIn?: string[];
    paymentMethodIdNin?: string[];
    paymentMethodIdExists?: boolean;
    // referenceCode (string) search options
    referenceCodeEq?: string;
    referenceCodeNe?: string;
    referenceCodeGt?: string;
    referenceCodeGte?: string;
    referenceCodeLt?: string;
    referenceCodeLte?: string;
    referenceCodeIn?: string[];
    referenceCodeNin?: string[];
    referenceCodeExists?: boolean;
    referenceCodeLike?: string;
    referenceCodeNlike?: string;
    // registryId (Ref<Registry>) search options
    registryIdEq?: string;
    registryIdIn?: string[];
    registryIdNin?: string[];
    registryIdExists?: boolean;
    // status (CartStatus) search options
    statusEq?: CartStatus;
    statusNe?: CartStatus;
    statusGt?: CartStatus;
    statusGte?: CartStatus;
    statusLt?: CartStatus;
    statusLte?: CartStatus;
    statusIn?: CartStatus[];
    statusNin?: CartStatus[];
    statusExists?: boolean;
    // updated (ActorTrace) search options
    updated?: ActorTraceSearchQuery;
    // updatedByOwnerUser (ActorTrace) search options
    updatedByOwnerUser?: ActorTraceSearchQuery;
}
