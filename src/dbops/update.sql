db.test.update(
	{
        "title":"GoodTitle"
    },
    {
        "$push":
            {
                "lines":
                    {
                        sx:1,
                        sy:1,
                        tx:30,
                        ty:30
                    }
            }
    }
)



db.test.update(
    {
        title:"GoodTitle"
    },
    {
        "$push":
            {
                components:
                    {
                        id: 4,
                        name: "BadGate",
                        shape: [
                            { sx: 30, sy: 20, tx: 100, ty: 20 },
                            { sx: 100, sy: 0, tx: 100, ty: 100 },
                            { sx: 0, sy: 80, tx: 100, ty: 80 },
                            { sx: 100, sy: 0, tx: 186, ty: 50 },
                            { sx: 100, sy: 100, tx: 186, ty: 50 },
                             { sx: 186, sy: 50, tx: 250, ty: 50 }
                        ],
                        pin: [
                            { x: 0, y: 20 },
                            { x: 0, y: 80 },
                            { x: 250, y: 50 }
                        ]
                    }
            }
    }
)


    db.test.update(
	{
        "title":"GoodTitle"
    },
    {
        "$push":
            {
                "lines":
                    {
                        sx:1,
                        sy:1,
                        tx:30,
                        ty:30
                    }
            }
    }
)
