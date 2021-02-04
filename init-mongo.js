db.createUser(
    {
        user    : "admin", //insert username here
        pwd     : "admin", //insert password here
        roles   : [
            {
                role    : "readWrite",
                db      : "Magic:The-Gathering-Archive" //insert database name here
            }
        ]
    }
)