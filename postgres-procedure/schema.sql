-- Table: ewallet.tb_transaction

-- DROP TABLE IF EXISTS ewallet.tb_transaction;

CREATE TABLE IF NOT EXISTS ewallet.tb_transaction
(
    s_id character varying COLLATE pg_catalog."default" NOT NULL,
    s_trace_number character varying COLLATE pg_catalog."default" NOT NULL,
    s_user_id character varying COLLATE pg_catalog."default" NOT NULL,
    s_merchant_id character varying COLLATE pg_catalog."default" NOT NULL,
    n_amount bigint NOT NULL,
    n_fee bigint NOT NULL,
    n_open_balance bigint NOT NULL,
    n_close_balance bigint NOT NULL,
    s_transaction_type character varying COLLATE pg_catalog."default" NOT NULL,
    s_source_ref character varying COLLATE pg_catalog."default",
    s_destination_ref character varying COLLATE pg_catalog."default",
    s_destination_name character varying COLLATE pg_catalog."default",
    s_remark character varying COLLATE pg_catalog."default",
    s_state character varying COLLATE pg_catalog."default",
    s_provider_code character varying COLLATE pg_catalog."default",
    s_chanel_code character varying COLLATE pg_catalog."default",
    s_chanel_ref character varying COLLATE pg_catalog."default",
    s_chanel_response_code character varying COLLATE pg_catalog."default",
    s_chanel_data character varying COLLATE pg_catalog."default",
    d_create timestamp with time zone NOT NULL DEFAULT now(),
    d_update timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT tb_transaction_pkey PRIMARY KEY (s_id),
    CONSTRAINT tb_transaction_s_trace_number_key UNIQUE (s_trace_number),
    CONSTRAINT tb_transaction_s_merchant_id_fkey FOREIGN KEY (s_merchant_id)
        REFERENCES ewallet.tb_merchant (s_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT tb_transaction_s_user_id_fkey FOREIGN KEY (s_user_id)
        REFERENCES ewallet.tb_user (s_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS ewallet.tb_transaction
    OWNER to postgres;

ALTER TABLE ewallet.tb_transaction ADD COLUMN n_total_amount bigint DEFAULT 0;