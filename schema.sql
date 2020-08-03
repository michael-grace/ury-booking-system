create schema bookings;

comment on schema bookings is 'For the URY booking system.';

create table bookings.request_levels(
    request_level int unique,
    description varchar
);

create table bookings.resources(
    resource int,
    name varchar,
    PRIMARY KEY (resource)
);

create table bookings.unique_resources(
    unique_id int,
    resource int,
    name varchar,
    primary key (unique_id),
    foreign key (resource) references bookings.resources(resource)
);

create table bookings.public_contacts(
    public_id int,
    contact varchar,
    primary key (public_id)
);

create table bookings.bookings(
    booking_id int,
    member_id int,
    request_level int,
    resource int,
    preference int,
    given_resource int,
    timeslot_id int,
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    public_id int,
    application_datetime timestamp with time zone,
    PRIMARY KEY (booking_id),
    foreign key (member_id) references public.member(memberid),
    foreign key (request_level) references bookings.request_levels(request_level),
    foreign key (resource) references bookings.resources(resource),
    foreign key (preference) references bookings.unique_resources(unique_id),
    foreign key (given_resource) references bookings.unique_resources(unique_id),
    foreign key (timeslot_id) references schedule.show_season_timeslot(show_season_timeslot_id),
    foreign key (public_id) references bookings.public_contacts(public_id)
);
