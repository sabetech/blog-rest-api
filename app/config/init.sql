create table vehicle_categories (
  id serial not null unique,
  vehicle_type varchar(64),
  vehicle_desc text,
  primary key(id)
);

create table vehicle_taxes (
  id serial not null unique,
  vehicle_categories_id integer,
  import_duty numeric(5,4),
  vat numeric(5,4),
  nhil numeric(5,4),
  getfund_levy numeric(5,4),
  au_levy numeric(5,4),
  ecowas numeric(5,4),
  exim_levy numeric(5,4),
  exam_fee numeric(5,4),
  process_fee numeric(5,4),
  special_import_levy numeric(5,4)
);


insert into vehicle_categories(vehicle_type, vehicle_desc)
values
    ('ambulance', 'Ambulance'),
    ('hearse', 'Hearse');

insert into vehicle_taxes 
  ( vehicle_categories_id, 
    import_duty, 
    vat, 
    nhil, 
    au_levy, 
    ecowas, 
    exim_levy, 
    exam_fee, 
    process_fee, 
    special_import_levy)
    values 
    (1,
     0.20,
     0.125,
     0.025,
     0.025,
     0.002,
     0.005,
     0.0075,
     0.01,
     0.02
     ),
     (2,
     0.20,
     0.125,
     0.025,
     0.025,
     0.002,
     0.005,
     0.0075,
     0.01,
     0.02
     );