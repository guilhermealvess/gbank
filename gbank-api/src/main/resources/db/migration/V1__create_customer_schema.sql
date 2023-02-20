CREATE TABLE customers(
    id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    CONSTRAINT customer_pkey PRIMARY KEY(id)
);
