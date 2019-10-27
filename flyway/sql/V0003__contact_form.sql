create schema nixie;

create table if not exists nixie.contact_form
(
    contact_form_id             bigint  not null generated by default as identity,
    email                       text not null,
    description                 text not null,
    domain                      text not null default 'nixiepixel.com',
    primary key (contact_form_id)
);