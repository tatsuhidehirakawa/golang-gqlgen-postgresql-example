CREATE TABLE IF NOT EXISTS account_master (
    account_id VARCHAR(30),
    account_name varchar(30) NOT NULL,
    primary key(account_id)
);

CREATE TABLE IF NOT EXISTS offer_list (
    offer_id VARCHAR(30),
    account_id VARCHAR(30) NOT NULL,
    primary key(offer_id),
    foreign key (account_id) references account_master(account_id)
);

