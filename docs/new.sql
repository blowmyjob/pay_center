create table pay_account(
   id bigint comment '自增主键' primary key auto_increment,
   app_id varchar(200) comment '',
   amount decimal comment '账户金额',
   status varchar(2) comment '账户状态',
   create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
   update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间'
)comment '支付账户表';

create table pay_record(
    id bigint comment '' primary key auto_increment,
    seq_id varchar(200) not null comment '支付记录流水号',
    app_id varchar(200) not null comment 'app_id',
    third_id varchar(200) not null commit '第三方支付流水号',
    product_id bigint not null comment '支付产品id',
    user_id bigint not null comment '支付用户id',
    pay_amount decimal not null comment '支付金额',
    pay_status varchar(10) not null comment '支付状态',
    pay_method varchar(10) not null comment '支付方式',
    create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间'
)comment '支付流水表';

create table pay_channel(
    id bigint comment '' primary key auto_increment,
    pay_name varchar(200) comment '支付名称' not null,
    pay_code varchar(200) comment '支付代码' not null,
    app_id   varchar(50) comment 'app_id' not null,
    app_secret varchar(100) comment 'app密钥' not null,
    merchant_id varchar(50) comment '商户号' not null,
    status varchar(50) comment '渠道状态' not null
)comment '支付商户渠道表'



