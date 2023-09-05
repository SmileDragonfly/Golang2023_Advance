CREATE OR REPLACE PROCEDURE ewallet.get_transactions_by_user(in_user_id character varying, OUT rs refcursor)
    LANGUAGE plpgsql
AS $$
BEGIN
    OPEN rs FOR
        SELECT *
        FROM ewallet.tb_transaction
        WHERE s_user_id = in_user_id;
END;
$$;