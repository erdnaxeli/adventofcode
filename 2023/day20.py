from itertools import product
import json


print(
    json.dumps(
        ["".join(reversed(x)) for x in sorted(product("01", repeat=12))],
        indent="\t",
    )
)
