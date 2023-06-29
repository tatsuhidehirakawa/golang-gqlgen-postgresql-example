CREATE TABLE IF NOT EXISTS account_master (
    account_id INT,
    account_name varchar(30) NOT NULL,
    primary key(account_id)
);

CREATE TABLE IF NOT EXISTS offer_list (
    offer_id INT,
    account_id INT,
    primary key(offer_id),
    foreign key (account_id) references account_master(account_id)
);

