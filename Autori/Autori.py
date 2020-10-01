us = list(input(""))
toReturn = "";
for item in us:
    if item.isupper():
        toReturn += item
        pass
print(toReturn)