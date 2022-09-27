-- data definition language
create database tugaspostgres;

create table customers (id serial, name text, primary key (id));
create table products (id serial, name text, primary key (id));
create table orders (id serial, order_date date, total int, primary key (id));

-- data manipulation language
insert into public.customers(id,name) values (1, 'Koje');
insert into public.products(id,name) values (1, 'Feeder');
insert into public.orders(id,order_date,total) values (1, '2022-01-01', 3);

alter table orders add constraint FK_customers foreign key (id) references customers(id);
alter table orders add constraint FK_products foreign key (id) references products(id);

drop table table_first;