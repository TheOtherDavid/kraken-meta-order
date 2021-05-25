CREATE SCHEMA kraken_meta_order;

CREATE TABLE meta_order(
    meta_order_id 		INT         PRIMARY KEY,
    meta_order_type 	TEXT        NOT NULL,
    status			    TEXT        NOT NULL,
    crt_dtm			    TIMESTAMP   NOT NULL,
    crt_usr_nm		    TEXT        NOT NULL,
    last_udt_dtm		TIMESTAMP   NOT NULL,
    last_udt_usr_nm	    TEXT        NOT NULL
);