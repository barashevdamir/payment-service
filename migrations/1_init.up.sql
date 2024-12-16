create table transactions
(
    id            bigserial not null,
    amount        numeric   not null,
    description   varchar   not null,
    recipient     varchar   not null,
    payment_token bigserial not null,
    payment_method varchar  not null,
    metadata      varchar   not null,
    refundable    boolean   not null,
    status        varchar   not null,
    created_at    timestamp not null,

    primary key (id),
    CONSTRAINT tx_payment_token_unique UNIQUE (payment_token)
);