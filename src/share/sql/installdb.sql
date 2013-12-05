--Make Tables
-- global_defs
create table if not exists global_defs(
        id int(4) not null default 0 primary key,
        notifi_mail_from char(50) not null default 'glvsadmin@localhost', 
        smtp_server char(15) not null default '127.0.0.1', 
        router_id char(50) not null default 'LVS_DEVEL'
);
insert into global_defs();

-- global_defs_mails
create table if not exists global_defs_mails(
        id int not null auto_increment primary key, 
        is_enabled bool not null default false,
        mail_add char(50) not null
);

-- vrrp_instance
create table if not exists vrrp_instance(
        id int not null auto_increment primary key,
        is_enabled bool not null default false,
        name char(30) not null default 'LVS_Instance',
        interface char(10) not null default 'eth0',
        virtual_rt_id int not null  default 1,
        auth_type int not null default 0, -- 0 means PASS
        auth_pass char not null default '1111'
);

-- vrrp_instance_rt
create table if not exists vrrp_instance_rts(
        id int not null auto_increment primary key,
        inst_id int not null,
        hostname char not null,
        state bool not null default false, -- false means MASTER, true means BACKUP
        priority int not null default 100,
        is_enabled int not null default false
);

-- vrrp_instance_vip
create table if not exists vrrp_instance_vip(
        id int not null auto_increment primary key,
        inst_id int not null,
        vip char(15) not null,
        is_enabled int not null default false
);

-- virtual_server
create table if not exists virtual_server(
        id int not null auto_increment primary key,
        vip_id int not null,
        port int not null,
        delay_loop int default 6,
        lb_algo int not null default 0, -- 0 means rr, 1 means ...
        lb_kind int not null default 0, -- 0 means NAT, 1 means DR, 2 means TUN
        persist_timeout int not null default 10,
        protocol int not null default 0, -- 0 means TCP, 1 means UDP ...
        nat_mask char(16), -- NAT mode only
        sorry_server_ip char(16),
        sorry_server_port int
);

-- real_server
create table if not exists real_server(
        id int not null auto_increment primary key,
        vs_id int not null,
        hostip char not null,
        port int not null,
        weight int default 3,
        check_method int not null default 0 -- 0 means TCP_CHECK, 1 means UDP_CHECK, 
                                            -- 2 means HTTP_GET, 3 means SSL_GET
);

-- real_server_check
create table if not exists real_server_check(
        id int not null auto_increment primary key,
        rs_id int not null,
        connect_timeout int not null default 10,
        nb_get_retry int not null default 3,
        delay_before_retry not null default 3,
        connect_port int -- TCP_CHECK/UDP_CHECK only
);

-- real_server_url
create table if not exists real_server_url(
        -- for HTTP_GET/SSL_GET
        id int not null auto_increment primary key,
        rs_id int not null,
        path char not null default '/',
        digest char not null
);
