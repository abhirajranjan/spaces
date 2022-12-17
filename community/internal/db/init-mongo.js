db.createUser(
    {
        user  : "communitydb",
        pwd   : "communitydbpassword",
        roles : [
            {
                role : "readWrite",
                db   : "community"
            }
        ]
    }
)