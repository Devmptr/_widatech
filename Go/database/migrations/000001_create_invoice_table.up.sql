create table invoices
(
    invoice_no    text                    not null,
    date          date                    not null,
    customer_name text                    not null,
    sales_person_name text                    not null,
    payment_type  enum ('CASH', 'CREDIT') not null,
    notes         text                    null
);