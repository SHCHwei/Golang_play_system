module myapp

go 1.16

replace myapp/dbcon => /dbcon

replace myapp/handlers => /handlers

require myapp/handlers v0.0.0-00010101000000-000000000000
