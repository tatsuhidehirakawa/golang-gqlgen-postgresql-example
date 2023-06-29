-- account_master `offer_list`とのリレーションがわかりやすいように101~104のaccount_idを持つレコードを作成
INSERT INTO account_master VALUES  (101,'account1');
INSERT INTO account_master VALUES  (102,'account2');
INSERT INTO account_master VALUES  (103,'account3');
INSERT INTO account_master VALUES  (104,'account4');

-- offer_list
INSERT INTO offer_list VALUES  (1,101);
INSERT INTO offer_list VALUES  (2,102);
INSERT INTO offer_list VALUES  (3,103);
INSERT INTO offer_list VALUES  (4,104);

